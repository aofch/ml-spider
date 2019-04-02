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
		Addr:     "210.5.152.217:30703",
		Password: "lol_wf+hl.1211+1522",
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

	NewToutiao().foodSpiderColly()
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
