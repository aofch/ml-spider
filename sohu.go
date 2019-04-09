// project ml-crawler
// ide GoLand
// author Administrator
// create time 2019-04-09 16:38
// description 搜狐爬虫
// Copyright (c) 2019, fucaihe@gmail.com All Rights Reserved.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strconv"
	"strings"
	"time"
)

// A Sohuspider
type Sohu struct{}

func NewSohu() *Sohu {
	return new(Sohu)
}

type SohuResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		WapArticleVOS []struct {
			ID     int    `json:"id"`     // 文章ID
			UserID int    `json:"userId"` // 搜狐号ID
			Type   int    `json:"type"`   // 1: 文本, 2: 图文 5: 视频
			Title  string `json:"title"`  // 标题
		} `json:"wapArticleVOS"`
	} `json:"data"`
}

// 搜狐号抓取
func (s *Sohu) spiderMPColly(mid, pageOn int, rdsTag string) {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1"

	c.OnError(func(r *colly.Response, e error) {
		log.Fatalln(r.StatusCode, e)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Printf("-------------------- 搜狐号 -- %d --------------------\n", mid)
		log.Println(r.URL)
	})

	c.OnScraped(func(r *colly.Response) {
		if r.StatusCode != 200 {
			log.Println("failed statuscode =", r.StatusCode)
			return
		}
		var result SohuResult
		if err := json.Unmarshal(r.Body, &result); err != nil {
			log.Println(err)
			return
		}

		if result.Code != 200 {
			log.Fatal(result.Code, result.Msg)
		}

		if len(result.Data.WapArticleVOS) == 0 {
			log.Println("数据抓取完成")
			return
		}
		for i, data := range result.Data.WapArticleVOS {
			data.Title = strings.TrimSpace(data.Title)
			if data.Title == "" || data.ID == 0 || data.UserID == 0 {
				log.Println("数据不规范")
				continue
			}
			if data.Type != 1 && data.Type != 2 {
				log.Println("数据类型不对", data.Type)
				continue
			}
			link := fmt.Sprintf("http://www.sohu.com/a/%d_%d", data.ID, data.UserID)

			temp := struct {
				Title string
				Link  string
			}{
				Title: data.Title,
				Link:  link,
			}
			body, err := json.Marshal(&temp)
			if err != nil {
				log.Println("failed to error marshal", err)
				continue
			}
			tmd5 := Get16MD5(temp.Title)

			log.Println(i, temp.Title, temp.Link)

			// 持久化
			rdsClient.HSet(rdsTag, tmd5, body)
		}

		// 下一页抓取
		pNoStr := r.Ctx.Get("pNo")

		pNo, _ := strconv.Atoi(pNoStr)
		if pNo == 0 {
			pNo = pageOn + 1
		}

		nextpNo := fmt.Sprintf("%d", pNo+1)
		r.Ctx.Put("pNo", nextpNo)

		time.Sleep(500 * time.Millisecond)
		nURL := fmt.Sprintf("https://v2.sohu.com/author-page-api/author-articles/wap/%d?pNo=%d", mid, pNo)
		if err := r.Request.Visit(nURL); err != nil {
			log.Fatal(err)
		}
	})

	nURL := fmt.Sprintf("https://v2.sohu.com/author-page-api/author-articles/wap/%d?pNo=%d", mid, pageOn)
	if err := c.Visit(nURL); err != nil {
		log.Fatal(err)
	}
}

// 搜狐分类新闻抓取
func (s *Sohu) spiderMTColly() {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36"

	c.OnError(func(r *colly.Response, e error) {
		log.Fatalln(r.StatusCode, e)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("-------------------- Visit", r.URL, "--------------------")
	})

	c.OnHTML(".news-list .news-box", func(e *colly.HTMLElement) {
		title := strings.TrimSpace(e.DOM.Find("h4 a").Text())
		ev := strings.TrimSpace(e.Attr("data-ev"))
		mid := strings.TrimSpace(e.Attr("data-media-id"))

		if ev == "" || mid == "" || title == "" {
			log.Println("数据不符合规矩")
			return
		}
		link := fmt.Sprintf("http://www.sohu.com/a/%s_%s", ev, mid)

		log.Println(title, link)
	})

	c.OnScraped(func(r *colly.Response) {
		pObj := r.Ctx.GetAny("p")
		var page int
		if pObj == nil {
			page = 2 // 从第二页开始
		} else {
			page = pObj.(int)
		}

		r.Ctx.Put("p", page+1)
		nextURL := fmt.Sprintf("http://mt.sohu.com/acg?p=%d", page)

		time.Sleep(200 * time.Millisecond)
		if err := r.Request.Visit(nextURL); err != nil {
			log.Fatal(err)
		}
	})

	if err := c.Visit("http://mt.sohu.com/acg?p=1"); err != nil {
		log.Fatal(err)
	}
}
