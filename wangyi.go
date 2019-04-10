// project ml-crawler
// ide GoLand
// author Administrator
// create time 2019-04-10 10:28
// description 网易新闻抓取
// Copyright (c) 2019, fucaihe@gmail.com All Rights Reserved.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"log"
	"os"
	"strings"
	"time"
)

type Wangyi struct{}

func NewWangyi() *Wangyi {
	return new(Wangyi)
}

type WangyiResult struct {
	TabList []struct {
		PostID string `json:"postid"`
		Title  string `json:"title"`
	} `json:"tab_list"`
	SubscribeInfo struct {
		EName string `json:"ename"`
		TName string `json:"tname"`
	} `json:"subscribe_info"`
}

//spiderMediaColly 网易号抓取
func (s *Wangyi) spiderMediaColly(totalPage int, mid, rdsTag string) {

	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:66.0) Gecko/20100101 Firefox/66.0"

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visit", r.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		log.Fatalln(r.StatusCode, e)
	})

	c.OnScraped(func(r *colly.Response) {
		if r.StatusCode != 200 {
			log.Println("failed statuscode =", r.StatusCode)
			return
		}
		var result WangyiResult
		if err := json.Unmarshal(r.Body, &result); err != nil {
			log.Println(err)
			return
		}

		if len(result.TabList) == 0 {
			log.Println("数据抓取完成")
			os.Exit(0)
		}

		log.Printf("---------- %s -- %s ----------", result.SubscribeInfo.TName, result.SubscribeInfo.EName)
		for i, data := range result.TabList {
			data.Title = strings.TrimSpace(data.Title)
			data.PostID = strings.TrimSpace(data.PostID)
			if data.Title == "" || data.PostID == "" {
				continue
			}

			link := fmt.Sprintf("https://c.m.163.com/news/a/%s.html", data.PostID)
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

		time.Sleep(2 * time.Second)
	})

	q, _ := queue.New(
		1,                                           // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)

	pageSize := 20
	for i := 0; i < totalPage; i++ {
		offset := i*pageSize + 1
		if i == 0 {
			offset = i * pageSize
		}
		nURL := fmt.Sprintf("https://c.m.163.com/nc/subscribe/list/%s/all/%d-%d.html", mid, offset, pageSize)
		q.AddURL(nURL)
	}
	q.Run(c)
}
