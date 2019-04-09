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
	txt = `
<div class="channel-list">
  
  <div class="item first-col js-channel-subscribe-item" data-index="0" data-fromid="m27214" data-channelid="undefined" data-channelname="装修不叫事">
    <div class="item-inner-box">
      <a href="/channel/m27214" target="_blank" class="channel-img">
        <img src="//i1.go2yd.com/image.php?url=http://s.go2yd.com/b/ieusdhtu.jpg&amp;type=thumbnail_300x300">
        
      </a>
      <div class="channel-info">
        <a href="/channel/m27214" target="_blank" class="channel-name">装修不叫事</a>
        <p class="channel-summary">北京江水平装修分享装修效果图、装修施工图、装修工艺、装修步骤、装修注意事项等装修知识。</p>
        <button class="subscribe js-subscribe-btn">关注</button>
        <button class="unsubscribe js-unsubscribe-btn">已关注</button>
      </div>
    </div>
  </div>

  
  <div class="item js-channel-subscribe-item" data-index="1" data-fromid="m948" data-channelid="undefined" data-channelname="设计馆">
    <div class="item-inner-box">
      <a href="/channel/m948" target="_blank" class="channel-img">
        <img src="//i1.go2yd.com/image.php?url=http://static.yidianzixun.com/media/63ec2700b1c0a05097ded32c41659abb.png&amp;type=thumbnail_300x300">
        
          <span class="icon-plus-round icon-plus-round-2"></span>
        
      </a>
      <div class="channel-info">
        <a href="/channel/m948" target="_blank" class="channel-name">设计馆</a>
        <p class="channel-summary">找装修灵感，学装修知识。</p>
        <button class="subscribe js-subscribe-btn">关注</button>
        <button class="unsubscribe js-unsubscribe-btn">已关注</button>
      </div>
    </div>
  </div>

  
  <div class="item js-channel-subscribe-item" data-index="2" data-fromid="m124966" data-channelid="undefined" data-channelname="HDD室内设计">
    <div class="item-inner-box">
      <a href="/channel/m124966" target="_blank" class="channel-img">
        <img src="//i1.go2yd.com/image.php?url=http://si1.go2yd.com/get-image/05WpmvyJPQ8&amp;type=thumbnail_300x300">
        
      </a>
      <div class="channel-info">
        <a href="/channel/m124966" target="_blank" class="channel-name">HDD室内设计</a>
        <p class="channel-summary">空间创意的梦想家,商业灵魂的缔造者，中国知名商业空间设计分享平台。</p>
        <button class="subscribe js-subscribe-btn">关注</button>
        <button class="unsubscribe js-unsubscribe-btn">已关注</button>
      </div>
    </div>
  </div>

  
  <div class="item last-col js-channel-subscribe-item" data-index="3" data-fromid="m115746" data-channelid="undefined" data-channelname="中装协设计网">
    <div class="item-inner-box">
      <a href="/channel/m115746" target="_blank" class="channel-img">
        <img src="//i1.go2yd.com/image.php?url=http://si1.go2yd.com/get-image/0EbH7NcbqEq&amp;type=thumbnail_300x300">
        
      </a>
      <div class="channel-info">
        <a href="/channel/m115746" target="_blank" class="channel-name">中装协设计网</a>
        <p class="channel-summary">中装协设计网是中国建筑装饰协会设计委员会官网；是中国建筑装饰行业权威的设计服务平台；是中国建筑装饰协会互联网媒体重要媒体之一；全国建筑装饰设计行业主流媒体之一。宗旨：关心设计师成长，关注设计师收入！</p>
        <button class="subscribe js-subscribe-btn">关注</button>
        <button class="unsubscribe js-unsubscribe-btn">已关注</button>
      </div>
    </div>
  </div>

  
  <div class="item first-col js-channel-subscribe-item" data-index="4" data-fromid="m121407" data-channelid="undefined" data-channelname="创意家居生活">
    <div class="item-inner-box">
      <a href="/channel/m121407" target="_blank" class="channel-img">
        <img src="//i1.go2yd.com/image.php?url=http://si1.go2yd.com/get-image/05LY76kIaEy&amp;type=thumbnail_300x300">
        
      </a>
      <div class="channel-info">
        <a href="/channel/m121407" target="_blank" class="channel-name">创意家居生活</a>
        <p class="channel-summary">家居融进创意，生活充满乐趣！</p>
        <button class="subscribe js-subscribe-btn">关注</button>
        <button class="unsubscribe js-unsubscribe-btn">已关注</button>
      </div>
    </div>
  </div>

  
  <div class="item js-channel-subscribe-item" data-index="5" data-fromid="m7666" data-channelid="undefined" data-channelname="住宅公园">
    <div class="item-inner-box">
      <a href="/channel/m7666" target="_blank" class="channel-img">
        <img src="//i1.go2yd.com/image.php?url=http://si1.go2yd.com/get-image/0IwZAObtdtw&amp;type=thumbnail_300x300">
        
      </a>
      <div class="channel-info">
        <a href="/channel/m7666" target="_blank" class="channel-name">住宅公园</a>
        <p class="channel-summary">用设计，让农村更美。</p>
        <button class="subscribe js-subscribe-btn">关注</button>
        <button class="unsubscribe js-unsubscribe-btn">已关注</button>
      </div>
    </div>
  </div>

  
  <div class="item js-channel-subscribe-item" data-index="6" data-fromid="m5516" data-channelid="undefined" data-channelname="室内设计师周小白">
    <div class="item-inner-box">
      <a href="/channel/m5516" target="_blank" class="channel-img">
        <img src="//i1.go2yd.com/image.php?url=http://si1.go2yd.com/get-image/0TH00msFzNI&amp;type=thumbnail_300x300">
        
          <span class="icon-plus-round icon-plus-round-2"></span>
        
      </a>
      <div class="channel-info">
        <a href="/channel/m5516" target="_blank" class="channel-name">室内设计师周小白</a>
        <p class="channel-summary">一个室内设计师的成长之路</p>
        <button class="subscribe js-subscribe-btn">关注</button>
        <button class="unsubscribe js-unsubscribe-btn">已关注</button>
      </div>
    </div>
  </div>

  
  <div class="item last-col js-channel-subscribe-item" data-index="7" data-fromid="m5341" data-channelid="undefined" data-channelname="韩胖说装修">
    <div class="item-inner-box">
      <a href="/channel/m5341" target="_blank" class="channel-img">
        <img src="//i1.go2yd.com/image.php?url=http://si1.go2yd.com/get-image/0L1ZlzRDr3A&amp;type=thumbnail_300x300">
        
          <span class="icon-plus-round icon-plus-round-2"></span>
        
      </a>
      <div class="channel-info">
        <a href="/channel/m5341" target="_blank" class="channel-name">韩胖说装修</a>
        <p class="channel-summary">每天为你分享精彩装修案例，装修知识，让你爱上生活，爱上家!</p>
        <button class="subscribe js-subscribe-btn">关注</button>
        <button class="unsubscribe js-unsubscribe-btn">已关注</button>
      </div>
    </div>
  </div>

  
  <div class="item first-col js-channel-subscribe-item" data-index="8" data-fromid="m11154" data-channelid="undefined" data-channelname="装修邦">
    <div class="item-inner-box">
      <a href="/channel/m11154" target="_blank" class="channel-img">
        <img src="//i1.go2yd.com/image.php?url=http://s.go2yd.com/b/iane8n4k.jpg&amp;type=thumbnail_300x300">
        
      </a>
      <div class="channel-info">
        <a href="/channel/m11154" target="_blank" class="channel-name">装修邦</a>
        <p class="channel-summary">为装修业主提供专业装修知识无缝墙布施工知识</p>
        <button class="subscribe js-subscribe-btn">关注</button>
        <button class="unsubscribe js-unsubscribe-btn">已关注</button>
      </div>
    </div>
  </div>

  
  <div class="item js-channel-subscribe-item" data-index="9" data-fromid="m4347" data-channelid="undefined" data-channelname="装修设计案例精选">
    <div class="item-inner-box">
      <a href="/channel/m4347" target="_blank" class="channel-img">
        <img src="//i1.go2yd.com/image.php?url=http://si1.go2yd.com/get-image/0MbrgAJpPU0&amp;type=thumbnail_300x300">
        
      </a>
      <div class="channel-info">
        <a href="/channel/m4347" target="_blank" class="channel-name">装修设计案例精选</a>
        <p class="channel-summary">分享最时尚的家居软装设计，最唯美的装修案例图片，最全面的风格设计精选，最接地气的装修方案推荐！</p>
        <button class="subscribe js-subscribe-btn">关注</button>
        <button class="unsubscribe js-unsubscribe-btn">已关注</button>
      </div>
    </div>
  </div>

  
  <div class="item js-channel-subscribe-item" data-index="10" data-fromid="m560" data-channelid="undefined" data-channelname="合肥飞墨设计">
    <div class="item-inner-box">
      <a href="/channel/m560" target="_blank" class="channel-img">
        <img src="//i1.go2yd.com/image.php?url=http://static.yidianzixun.com/media/b055b14c2b693d51aceadc05ad36780a.jpg&amp;type=thumbnail_300x300">
        
      </a>
      <div class="channel-info">
        <a href="/channel/m560" target="_blank" class="channel-name">合肥飞墨设计</a>
        <p class="channel-summary">飞墨设计理念为：设计以人为本。
因为喜欢，所以专业</p>
        <button class="subscribe js-subscribe-btn">关注</button>
        <button class="unsubscribe js-unsubscribe-btn">已关注</button>
      </div>
    </div>
  </div>

  
  <div class="item last-col js-channel-subscribe-item" data-index="11" data-fromid="m99346" data-channelid="undefined" data-channelname="建筑师的非建筑">
    <div class="item-inner-box">
      <a href="/channel/m99346" target="_blank" class="channel-img">
        <img src="//i1.go2yd.com/image.php?url=http://si1.go2yd.com/get-image/03qZSJg1hvE&amp;type=thumbnail_300x300">
        
      </a>
      <div class="channel-info">
        <a href="/channel/m99346" target="_blank" class="channel-name">建筑师的非建筑</a>
        <p class="channel-summary">感动你的，必先是感动我的</p>
        <button class="subscribe js-subscribe-btn">关注</button>
        <button class="unsubscribe js-unsubscribe-btn">已关注</button>
      </div>
    </div>
  </div>

  
  <div class="item first-col js-channel-subscribe-item" data-index="12" data-fromid="m12311" data-channelid="undefined" data-channelname="装修也疯狂">
    <div class="item-inner-box">
      <a href="/channel/m12311" target="_blank" class="channel-img">
        <img src="//i1.go2yd.com/image.php?url=http://s.go2yd.com/b/ibal5vpt.jpg&amp;type=thumbnail_300x300">
        
      </a>
      <div class="channel-info">
        <a href="/channel/m12311" target="_blank" class="channel-name">装修也疯狂</a>
        <p class="channel-summary">找最好的装修，做最好的设计，让想要的生活成为每天的日子！</p>
        <button class="subscribe js-subscribe-btn">关注</button>
        <button class="unsubscribe js-unsubscribe-btn">已关注</button>
      </div>
    </div>
  </div>

  
  <div class="item js-channel-subscribe-item" data-index="13" data-fromid="m133788" data-channelid="undefined" data-channelname="土巴兔装修家居">
    <div class="item-inner-box">
      <a href="/channel/m133788" target="_blank" class="channel-img">
        <img src="//i1.go2yd.com/image.php?url=http://si1.go2yd.com/get-image/0K9Evim8Nn6&amp;type=thumbnail_300x300">
        
      </a>
      <div class="channel-info">
        <a href="/channel/m133788" target="_blank" class="channel-name">土巴兔装修家居</a>
        <p class="channel-summary">服务中国1800万家庭的互联网装修平台，分享10000套装修案例！</p>
        <button class="subscribe js-subscribe-btn">关注</button>
        <button class="unsubscribe js-unsubscribe-btn">已关注</button>
      </div>
    </div>
  </div>

  
  <div class="item js-channel-subscribe-item" data-index="14" data-fromid="m68417" data-channelid="undefined" data-channelname="人民网美丽乡村">
    <div class="item-inner-box">
      <a href="/channel/m68417" target="_blank" class="channel-img">
        <img src="//i1.go2yd.com/image.php?url=http://si1.go2yd.com/get-image/03p5ar6a8Su&amp;type=thumbnail_300x300">
        
      </a>
      <div class="channel-info">
        <a href="/channel/m68417" target="_blank" class="channel-name">人民网美丽乡村</a>
        <p class="channel-summary">人民网美丽乡村报道集</p>
        <button class="subscribe js-subscribe-btn">关注</button>
        <button class="unsubscribe js-unsubscribe-btn">已关注</button>
      </div>
    </div>
  </div>

  
  <div class="item last-col js-channel-subscribe-item" data-index="15" data-fromid="m32596" data-channelid="undefined" data-channelname="设计之旅">
    <div class="item-inner-box">
      <a href="/channel/m32596" target="_blank" class="channel-img">
        <img src="//i1.go2yd.com/image.php?url=http://s.go2yd.com/b/ifkerm0f.jpg&amp;type=thumbnail_300x300">
        
      </a>
      <div class="channel-info">
        <a href="/channel/m32596" target="_blank" class="channel-name">设计之旅</a>
        <p class="channel-summary">设计之旅，专注设计师国际游学！</p>
        <button class="subscribe js-subscribe-btn">关注</button>
        <button class="unsubscribe js-unsubscribe-btn">已关注</button>
      </div>
    </div>
  </div>

  
  <div class="item first-col js-channel-subscribe-item" data-index="16" data-fromid="m99952" data-channelid="undefined" data-channelname="建E室内设计网">
    <div class="item-inner-box">
      <a href="/channel/m99952" target="_blank" class="channel-img">
        <img src="//i1.go2yd.com/image.php?url=http://si1.go2yd.com/get-image/09CXSW6UHlg&amp;type=thumbnail_300x300">
        
      </a>
      <div class="channel-info">
        <a href="/channel/m99952" target="_blank" class="channel-name">建E室内设计网</a>
        <p class="channel-summary">最美的家、最高端的样板房设计、最酷炫的设计作品，最欢神的生活方式都在这里！</p>
        <button class="subscribe js-subscribe-btn">关注</button>
        <button class="unsubscribe js-unsubscribe-btn">已关注</button>
      </div>
    </div>
  </div>

  
  <div class="item js-channel-subscribe-item" data-index="17" data-fromid="m4347" data-channelid="undefined" data-channelname="装修设计案例精选">
    <div class="item-inner-box">
      <a href="/channel/m4347" target="_blank" class="channel-img">
        <img src="//i1.go2yd.com/image.php?url=http://si1.go2yd.com/get-image/0MbrgAJpPU0&amp;type=thumbnail_300x300">
        
      </a>
      <div class="channel-info">
        <a href="/channel/m4347" target="_blank" class="channel-name">装修设计案例精选</a>
        <p class="channel-summary">分享最时尚的家居软装设计，最唯美的装修案例图片，最全面的风格设计精选，最接地气的装修方案推荐！</p>
        <button class="subscribe js-subscribe-btn">关注</button>
        <button class="unsubscribe js-unsubscribe-btn">已关注</button>
      </div>
    </div>
  </div>
</div>
`
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

	//ydzx := NewYidianzixun()
	//for i, mid := range ydzx.mediaList(txt) {
	//	log.Println(i, mid)
	//	ydzx.spiderColly(mid, "家居", "ydzx_home")
	//
	//	time.Sleep(time.Second * 2)
	//}

	// NewToutiao().spiderColly("news_game", "toutiao_game") // 游戏
	// NewToutiao().spiderColly("news_tech", "toutiao_tech") // 科技
	// NewToutiao().spiderColly("news_history", "toutiao_history") // 历史
	// NewToutiao().spiderColly("news_military", "toutiao_military") // 军事
	// NewToutiao().spiderColly("news_sports", "toutiao_sports") // 体育
	// NewToutiao().spiderColly("news_finance", "toutiao_finance") // 财经
	// NewToutiao().spiderColly("news_car", "toutiao_car") // 汽车
	// NewToutiao().spiderColly("news_food", "toutiao_food")

	NewSohu().spiderMPColly(100253173, 1, "sohu_comic")

	// 一点资讯, 头条
	// 动漫 1544 comic
	// 文化 1968 culture
	// 情感 3008 emotion
	// 美食 4335 food
	// 历史 3294 history
	// 家居 1006 home
	// 房产 963 house
	// 职场 1186 job
	// 军事 3110 military
	// 时政 474 politics
	// 科学 1271 science
	// 社会 896 society
	// 体育 3508 sports
	// 科技 4207 tech
	// 旅游 3341 travel

	// 育儿 6034 baby
	// 汽车 11687 car
	// 教育 5740 education
	// 娱乐 6135 entertainment
	// 时尚 5765 fashion
	// 财经 6636 finance
	// 星座 5812 fortune
	// 游戏 5067 game
	// 健康 6373 healthy
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
