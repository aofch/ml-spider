// project ml-crawler
// ide GoLand
// author Administrator
// create time 2019-04-08 10:32
// description 一点资讯爬虫
// Copyright (c) 2019, fucaihe@gmail.com All Rights Reserved.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/storage"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var (
	ydzxScript = `
        function good (n,e,i,t){
            for (var o = "sptoken", a = "", c = 0; c < arguments.length; c++){
                o += arguments[c];
            }

            for (var c = 0; c < o.length; c++) {
                var r = 10 ^ o.charCodeAt(c);
                a += String.fromCharCode(r)
            }
            return a
        }
`
)

func getSPT(args ...string) string {
	o := "sptoken"
	a := ""
	for c := 1; c < len(args); c++ {
		o += args[c]
	}
	for _, b := range []rune(o) {
		r := 10 ^ b
		a += string(r)
	}
	return a
}

// A Yidianzixun spider
type Yidianzixun struct{}

func NewYidianzixun() *Yidianzixun {
	return new(Yidianzixun)
}

type YidianzixunResult struct {
	Status      string `json:"status"`
	Code        int    `json:"code"`
	ChannelID   string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	ChannelType string `json:"channel_type"`
	Offset      int    `json:"offset"`
	Datas       []struct {
		Title string `json:"title"` // 标题
		DocID string `json:"docid"` // 文章ID, 用于生成文章链接
		CType string `json:"ctype"` // news, video_live
		// DType    int    `json:"dtype"`    // 1, 3, 23
		Category string `json:"category"` // 汽车,财经
	} `json:"result"`
}

func (s *Yidianzixun) spiderColly(channelID, tag, rdsTag string) {
	cURL := "http://www.yidianzixun.com/channel/" + channelID

	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:65.0) Gecko/20100101 Firefox/65.0"
	var stge storage.InMemoryStorage
	if err := c.SetStorage(&stge); err != nil {
		log.Fatal(err)
	}

	c.OnRequest(func(r *colly.Request) {
		log.Printf("-------------------- 一点资讯 -- %s -- %s --------------------\n", tag, channelID)

		log.Println(r.URL)

		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
		r.Headers.Set("Cache-Control", "no-cache")
		r.Headers.Set("Host", "www.yidianzixun.com")
		r.Headers.Set("Pragma", "no-cache")
		r.Headers.Set("Referer", cURL)

		nR, _ := url.Parse("http://www.yidianzixun.com/")
		r.Headers.Set("Cookie", stge.Cookies(nR))
	})

	c.OnError(func(r *colly.Response, e error) {
		log.Fatalln(r.StatusCode, e)
	})

	c.OnScraped(func(r *colly.Response) {
		// 登录
		if r.Request.URL.Path == "/mp_sign_in" {
			var loginRst map[string]string
			json.Unmarshal(r.Body, &loginRst)

			if loginRst["status"] != "success" {
				log.Fatal(loginRst["status"], loginRst["code"], loginRst["reason"])
			}
			log.Println("登录成功")
			return
		}

		cid := r.Request.Ctx.Get("cid")
		if cid == "" {
			log.Println("not found channel_id")
			return
		}

		if strings.Contains(r.Request.URL.Path, "/channel/") {
			n := "/home/q/news_list_for_channel?channel_id=" + cid + "&cstart=0&cend=10&infinite=true&refresh=1&__from__=pc&multi=5"
			nSPT := getSPT(n, cid, "0", "10")
			nSPT = url.QueryEscape(nSPT)
			apiURL := fmt.Sprintf("http://www.yidianzixun.com/home/q/news_list_for_channel?channel_id=%s&cstart=0&cend=10&infinite=true&refresh=1&__from__=pc&multi=5&_spt=%s&appid=web_yidian&_=%d", cid, nSPT, time.Now().UnixNano()/1e6)
			r.Ctx.Put("cstart", "10")
			r.Request.Visit(apiURL)
			return
		}

		if r.StatusCode != 200 {
			log.Println("failed statuscode =", r.StatusCode)
			return
		}

		var result YidianzixunResult
		if err := json.Unmarshal(r.Body, &result); err != nil {
			log.Println(err)
			return
		}

		if result.Code != 0 {
			log.Println("抓取数据失败", result.Code, result.Status)
			return
		}

		if len(result.Datas) == 0 {
			log.Println("数据抓取完成")
			return
		}

		for i, data := range result.Datas {
			link := "http://www.yidianzixun.com/article/" + data.DocID
			if data.CType != "news" || data.Category != tag {
				log.Println(i, "该条数据不正确", data.CType, data.Category, link)
				continue
			}

			temp := struct {
				Title string
				Link  string
			}{
				Title: strings.TrimSpace(data.Title),
				Link:  strings.TrimSpace(link),
			}
			body, err := json.Marshal(&temp)
			if err != nil {
				log.Println("failed to error marshal", err)
				continue
			}
			if temp.Title == "" {
				log.Println("标题为空")
				continue
			}
			tmd5 := Get16MD5(temp.Title)

			log.Println(i, data.Category, temp.Title, temp.Link)

			// 持久化
			rdsClient.HSet(rdsTag, tmd5, body)
		}

		cstartStr := r.Ctx.Get("cstart")
		cstart, _ := strconv.Atoi(cstartStr)
		cendStr := fmt.Sprintf("%d", cstart+10)
		r.Ctx.Put("cstart", cendStr)

		log.Println(cstartStr, cendStr)
		time.Sleep(200 * time.Millisecond)

		n := "/home/q/news_list_for_channel?channel_id=" + cid + "&cstart=" + cstartStr + "&cend=" + cendStr + "&infinite=true&refresh=1&__from__=pc&multi=5"
		nSPT := getSPT(n, cid, cstartStr, cendStr)
		nSPT = url.QueryEscape(nSPT)

		apiURL := "http://www.yidianzixun.com/home/q/news_list_for_channel?channel_id=" +
			cid + "&cstart=" +
			cstartStr + "&cend=" +
			cendStr + "&infinite=true&refresh=1&__from__=pc&multi=5&_spt=" + nSPT + "&appid=yidian&_=" + fmt.Sprintf("%d", time.Now().UnixNano()/1e6)
		r.Request.Visit(apiURL)
	})

	c.OnHTML(".channel-nav", func(e *colly.HTMLElement) {
		ne := e.DOM.Find("a.active")
		cid, _ := ne.Attr("data-channelid")
		if "" == strings.TrimSpace(ne.Text()) {
			cid = channelID
		}
		log.Println("channel_id", cid)
		e.Request.Ctx.Put("cid", cid)
	})

	loginURL := "http://www.yidianzixun.com/mp_sign_in"
	if err := c.Post(loginURL, map[string]string{"username": "fucaihe@pryun.com.cn", "password": "walterfch511"}); err != nil {
		log.Fatal("login fail", err)
	}

	if err := c.Visit(cURL); err != nil {
		log.Fatal(err)
	}
}

func (s *Yidianzixun) mediaList(txt string) []string {
	var cidArr []string
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(txt))
	doc.Find(".channel-img").Each(func(i int, selection *goquery.Selection) {
		href := strings.TrimSpace(selection.AttrOr("href", ""))
		if href == "" {
			return
		}

		cid := strings.TrimLeft(href, "/channel/")
		cidArr = append(cidArr, cid)
	})

	return cidArr
}
