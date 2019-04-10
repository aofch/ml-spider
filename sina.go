// project ml-crawler
// ide GoLand
// author Administrator
// create time 2019-04-10 08:56
// description 新浪抓取
// Copyright (c) 2019, fucaihe@gmail.com All Rights Reserved.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"log"
	"strings"
	"time"
)

type Sina struct{}

func NewSina() *Sina {
	return new(Sina)
}

type SinaResult struct {
	Result struct {
		Status struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		} `json:"status"`
		Datas []struct {
			CategoryID      string `json:"categoryid"`
			Link            string `json:"url"`
			VideoTimeLength string `json:"video_time_length"`
			Title           string `json:"title"`
			MediaName       string `json:"media_name"`
		} `json:"data"`
	} `json:"result"`
}

// 新浪文化抓取
func (s *Sina) spiderCulColly() {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:66.0) Gecko/20100101 Firefox/66.0"

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visit", r.URL)

		r.Headers.Set("Referer", "http://cul.news.sina.com.cn/")
		r.Headers.Set("Host", "feed.mix.sina.com.cn")
	})

	c.OnError(func(r *colly.Response, e error) {
		log.Fatalln(r.StatusCode, e)
	})

	c.OnScraped(func(r *colly.Response) {
		if r.StatusCode != 200 {
			log.Println("failed statuscode =", r.StatusCode)
			return
		}
		var result SinaResult
		if err := json.Unmarshal(r.Body, &result); err != nil {
			log.Println(err)
			return
		}

		if result.Result.Status.Code != 0 {
			log.Fatal(result.Result.Status.Code, result.Result.Status.Msg)
		}

		for i, data := range result.Result.Datas {
			data.Title = strings.TrimSpace(data.Title)
			data.Link = strings.TrimSpace(data.Link)
			if data.Title == "" || data.VideoTimeLength != "0" || data.CategoryID != "1" {
				log.Println(i, "该条数据不正确", data.CategoryID)
				continue
			}

			temp := struct {
				Title string
				Link  string
			}{
				Title: data.Title,
				Link:  data.Link,
			}
			body, err := json.Marshal(&temp)
			if err != nil {
				log.Println("failed to error marshal", err)
				continue
			}
			tmd5 := Get16MD5(temp.Title)

			log.Println(i, temp.Title, temp.Link)

			// 持久化
			rdsClient.HSet("sina_culture", tmd5, body)
		}

	})

	q, _ := queue.New(
		1,                                           // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)

	for i := 0; i < 65; i++ {
		// Add URLs to the queue
		nURL := fmt.Sprintf("http://feed.mix.sina.com.cn/api/roll/get?pageid=411&lid=2595&num=20&encode=utf-8&page=%d&callback=&_=%d", i+1, time.Now().UnixNano()/1e6)
		q.AddURL(nURL)
	}
	// Consume URLs
	q.Run(c)
}
