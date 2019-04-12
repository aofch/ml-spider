// project ml-crawler
// ide GoLand
// author Administrator
// create time 2019-04-12 09:17
// description 更新内容
// Copyright (c) 2019, fucaihe@gmail.com All Rights Reserved.

package main

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/json-iterator/go"
	"github.com/robertkrimen/otto"
	"html"
	"log"
	"strings"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var (
	classifyArr = []string{
		"car",           // 汽车
		"emotion",       // 情感
		"culture",       // 文化
		"home",          // 家居
		"house",         // 房产
		"job",           // 职场
		"military",      // 军事
		"science",       // 科学
		"history",       // 历史
		"sports",        // 体育
		"tech",          // 科技
		"travel",        // 旅游
		"food",          // 美食
		"comic",         // 动漫
		"baby",          // 育儿
		"education",     // 教育
		"entertainment", // 娱乐
		"fashion",       // 时尚
		"finance",       // 财经
		"fortune",       // 星座
		"game",          // 游戏
		"healthy",       // 健康
		"politics",      // 时政
		"agriculture",   // 三农
		"world",         // 国际
		"collect",       // 收藏
	}
	sourceArr = []string{"toutiao", "sina", "sohu", "ydzx"}
)

func spiderContent() {
	//for i, tag := range classifyArr {
	//	result := frdsClient.HGetAll(tag).Val()
	//	log.Printf("--------------- %d %s %d ---------------\n", i+1, tag, len(result))
	//	for k, data := range result {
	//		log.Println(k, data)
	//		break
	//	}
	//}
	spiderHandle("https://www.toutiao.com/i6630564617825305102/")
}

func spiderHandle(link string) {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:66.0) Gecko/20100101 Firefox/66.0"

	c.OnRequest(func(r *colly.Request) {
		log.Println("visit", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println(r.StatusCode)

		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(r.Body))
		handleErr(err)

		doc.Find("script").Each(func(i int, s *goquery.Selection) {
			msgTxt := strings.TrimSpace(s.Text())
			if strings.HasPrefix(msgTxt, "var BASE_DATA = {") {
				//msgTxt = strings.TrimLeft(msgTxt, "var BASE_DATA = ")
				msgTxt = strings.TrimSpace(html.UnescapeString(msgTxt))

				msgTxt = msgTxt + "\r\nfunction getData(){return BASE_DATA}"

				vvm := otto.New()
				_, err := vvm.Run(msgTxt)
				handleErr(err)

				val, err := vvm.Call("getData", nil)
				handleErr(err)
				data, err := val.Export()
				handleErr(err)
				dataMap := data.(map[string]interface{})
				aInfoi := dataMap["articleInfo"]
				ainfo := aInfoi.(map[string]interface{})
				log.Println(ainfo["content"])
				//val, err := vvm.Eval(msgTxt)
				//handleErr(err)
				//log.Println(val)

				//rstMap := make(map[string]interface{})
				//handleErr(json.Unmarshal([]byte(msgTxt), &rstMap))
				//log.Println(len(rstMap))
			}

		})
	})

	handleErr(c.Visit(link))
}

func zhenghe() {
	for i, tag := range classifyArr {
		for _, source := range sourceArr {
			result, err := rdsClient.HGetAll(source + "_" + tag).Result()
			handleErr(err)

			log.Printf("--------------- %d %s %s %d ---------------\n", i+1, tag, source, len(result))
			time.Sleep(2000 * time.Millisecond)
			for k, data := range result {
				log.Println(data)
				frdsClient.HSet(tag, k, data)
			}
		}

		log.Println("=============== ===============")
	}
}
