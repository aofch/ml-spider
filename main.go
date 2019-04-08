// project ml-crawler
// ide GoLand
// author Administrator
// create time 2019-04-02 11:00
// description 机器学习数据爬虫
// Copyright (c) 2019, fucaihe@gmail.com All Rights Reserved.

package main

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/go-redis/redis"
	"log"
)

var (
	rdsClient *redis.Client
)

func init() {
	rdsClient = redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       1,
	})
	_, err := rdsClient.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer func() {
		err := rdsClient.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	NewYidianzixun().spiderColly("u144", "美食", "ydzx_food")

	// NewToutiao().spiderColly("news_game", "toutiao_game") // 游戏
	// NewToutiao().spiderColly("news_tech", "toutiao_tech") // 科技
	// NewToutiao().spiderColly("news_history", "toutiao_history") // 历史
	// NewToutiao().spiderColly("news_military", "toutiao_military") // 军事
	// NewToutiao().spiderColly("news_sports", "toutiao_sports") // 体育
	// NewToutiao().spiderColly("news_finance", "toutiao_finance") // 财经
	// NewToutiao().spiderColly("news_car", "toutiao_car") // 汽车
}

// GetMD5 获取md5值
func GetMD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Get16MD5 获取16位md5值
func Get16MD5(s string) string {
	return string([]byte(GetMD5(s))[8:24])
}
