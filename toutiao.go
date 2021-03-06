// project ml-crawler
// ide GoLand
// author Administrator
// create time 2019-04-02 11:01
// description 今日头条爬虫
// Copyright (c) 2019, fucaihe@gmail.com All Rights Reserved.

package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/runner"
	"github.com/gocolly/colly"
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	argsScript = `
function t(e, t) {
    var n = (65535 & e) + (65535 & t),
        r = (e >> 16) + (t >> 16) + (n >> 16);
    return r << 16 | 65535 & n
}

function n(e, t) {
    return e << t | e >>> 32 - t
}

function r(e, r, o, i, u, a) {
    return t(n(t(t(r, e), t(i, a)), u), o)
}

function o(e, t, n, o, i, u, a) {
    return r(t & n | ~t & o, e, t, i, u, a)
}

function i(e, t, n, o, i, u, a) {
    return r(t & o | n & ~o, e, t, i, u, a)
}

function u(e, t, n, o, i, u, a) {
    return r(t ^ n ^ o, e, t, i, u, a)
}

function a(e, t, n, o, i, u, a) {
    return r(n ^ (t | ~o), e, t, i, u, a)
}

function s(e, n) {
    e[n >> 5] |= 128 << n % 32,
        e[(n + 64 >>> 9 << 4) + 14] = n;
    var r,
        s,
        c,
        l,
        f,
        p = 1732584193,
        d = -271733879,
        h = -1732584194,
        m = 271733878;
    for (r = 0; r < e.length; r += 16)
        s = p, c = d, l = h, f = m, p = o(p, d, h, m, e[r], 7, -680876936), m = o(m, p, d, h, e[r + 1], 12, -389564586), h = o(h, m, p, d, e[r + 2], 17, 606105819), d = o(d, h, m, p, e[r + 3], 22, -1044525330), p = o(p, d, h, m, e[r + 4], 7, -176418897), m = o(m, p, d, h, e[r + 5], 12, 1200080426), h = o(h, m, p, d, e[r + 6], 17, -1473231341), d = o(d, h, m, p, e[r + 7], 22, -45705983), p = o(p, d, h, m, e[r + 8], 7, 1770035416), m = o(m, p, d, h, e[r + 9], 12, -1958414417), h = o(h, m, p, d, e[r + 10], 17, -42063), d = o(d, h, m, p, e[r + 11], 22, -1990404162), p = o(p, d, h, m, e[r + 12], 7, 1804603682), m = o(m, p, d, h, e[r + 13], 12, -40341101), h = o(h, m, p, d, e[r + 14], 17, -1502002290), d = o(d, h, m, p, e[r + 15], 22, 1236535329), p = i(p, d, h, m, e[r + 1], 5, -165796510), m = i(m, p, d, h, e[r + 6], 9, -1069501632), h = i(h, m, p, d, e[r + 11], 14, 643717713), d = i(d, h, m, p, e[r], 20, -373897302), p = i(p, d, h, m, e[r + 5], 5, -701558691), m = i(m, p, d, h, e[r + 10], 9, 38016083), h = i(h, m, p, d, e[r + 15], 14, -660478335), d = i(d, h, m, p, e[r + 4], 20, -405537848), p = i(p, d, h, m, e[r + 9], 5, 568446438), m = i(m, p, d, h, e[r + 14], 9, -1019803690), h = i(h, m, p, d, e[r + 3], 14, -187363961), d = i(d, h, m, p, e[r + 8], 20, 1163531501), p = i(p, d, h, m, e[r + 13], 5, -1444681467), m = i(m, p, d, h, e[r + 2], 9, -51403784), h = i(h, m, p, d, e[r + 7], 14, 1735328473), d = i(d, h, m, p, e[r + 12], 20, -1926607734), p = u(p, d, h, m, e[r + 5], 4, -378558), m = u(m, p, d, h, e[r + 8], 11, -2022574463), h = u(h, m, p, d, e[r + 11], 16, 1839030562), d = u(d, h, m, p, e[r + 14], 23, -35309556), p = u(p, d, h, m, e[r + 1], 4, -1530992060), m = u(m, p, d, h, e[r + 4], 11, 1272893353), h = u(h, m, p, d, e[r + 7], 16, -155497632), d = u(d, h, m, p, e[r + 10], 23, -1094730640), p = u(p, d, h, m, e[r + 13], 4, 681279174), m = u(m, p, d, h, e[r], 11, -358537222), h = u(h, m, p, d, e[r + 3], 16, -722521979), d = u(d, h, m, p, e[r + 6], 23, 76029189), p = u(p, d, h, m, e[r + 9], 4, -640364487), m = u(m, p, d, h, e[r + 12], 11, -421815835), h = u(h, m, p, d, e[r + 15], 16, 530742520), d = u(d, h, m, p, e[r + 2], 23, -995338651), p = a(p, d, h, m, e[r], 6, -198630844), m = a(m, p, d, h, e[r + 7], 10, 1126891415), h = a(h, m, p, d, e[r + 14], 15, -1416354905), d = a(d, h, m, p, e[r + 5], 21, -57434055), p = a(p, d, h, m, e[r + 12], 6, 1700485571), m = a(m, p, d, h, e[r + 3], 10, -1894986606), h = a(h, m, p, d, e[r + 10], 15, -1051523), d = a(d, h, m, p, e[r + 1], 21, -2054922799), p = a(p, d, h, m, e[r + 8], 6, 1873313359), m = a(m, p, d, h, e[r + 15], 10, -30611744), h = a(h, m, p, d, e[r + 6], 15, -1560198380), d = a(d, h, m, p, e[r + 13], 21, 1309151649), p = a(p, d, h, m, e[r + 4], 6, -145523070), m = a(m, p, d, h, e[r + 11], 10, -1120210379), h = a(h, m, p, d, e[r + 2], 15, 718787259), d = a(d, h, m, p, e[r + 9], 21, -343485551), p = t(p, s), d = t(d, c), h = t(h, l), m = t(m, f);
    return [p, d, h, m]
}

function c(e) {
    var t,
        n = "";
    for (t = 0; t < 32 * e.length; t += 8)
        n += String.fromCharCode(e[t >> 5] >>> t % 32 & 255);
    return n
}

function l(e) {
    var t,
        n = [];
    for (n[(e.length >> 2) - 1] = void 0, t = 0; t < n.length; t += 1)
        n[t] = 0;
    for (t = 0; t < 8 * e.length; t += 8)
        n[t >> 5] |= (255 & e.charCodeAt(t / 8)) << t % 32;
    return n
}

function f(e) {
    return c(s(l(e), 8 * e.length))
}

function p(e, t) {
    var n,
        r,
        o = l(e),
        i = [],
        u = [];
    for (i[15] = u[15] = void 0, o.length > 16 && (o = s(o, 8 * e.length)), n = 0; 16 > n; n += 1)
        i[n] = 909522486 ^ o[n], u[n] = 1549556828 ^ o[n];
    return r = s(i.concat(l(t)), 512 + 8 * t.length),
        c(s(u.concat(r), 640))
}

function d(e) {
    var t,
        n,
        r = "0123456789abcdef",
        o = "";
    for (n = 0; n < e.length; n += 1)
        t = e.charCodeAt(n), o += r.charAt(t >>> 4 & 15) + r.charAt(15 & t);
    return o
}

function h(e) {
    return unescape(encodeURIComponent(e))
}

function m(e) {
    return f(h(e))
}

function g(e) {
    return d(m(e))
}

function getHoney() {
    var t = Math.floor((new Date).getTime() / 1e3),
        e = t.toString(16).toUpperCase(),
        i = g(t).toString().toUpperCase();
    if (8 != e.length)
        return {
            as: "479BB4B7254C150",
            cp: "7E0AC8874BB0985"
        };
    for (var n = i.slice(0, 5), a = i.slice(-5), s = "", o = 0; 5 > o; o++)
        s += n[o] + e[o];
    for (var r = "", c = 0; 5 > c; c++)
        r += e[c + 3] + a[c];
    return {
        as: "A1" + s + e.slice(-3),
        cp: e.slice(0, 3) + r + "E1"
    }
}
`
	vm *otto.Otto
)

// A Toutiao spider
type Toutiao struct{}

func NewToutiao() *Toutiao {
	return new(Toutiao)
}

// 美食板块数据结构
type Food struct {
	HasMore bool   `json:"has_more"` // 是否还有更多
	Message string `json:"message"`  // success
	Next    struct {
		MaxBehotTime int64 `json:"max_behot_time"`
	} `json:"next"`
	Data []struct {
		Title        string `json:"title"`         // 标题
		SourceURL    string `json:"source_url"`    // 这个都是头条的链接
		Tag          string `json:"tag"`           // 数据标签
		ArticleGenre string `json:"article_genre"` // 文章类型 article, wenda, gallery
	} `json:"data"` // 内容数据
}

func (s *Toutiao) spiderColly(tag, rdsTag string) {
	vm = otto.New()
	vm.Run(argsScript)

	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36"

	c.OnRequest(func(r *colly.Request) {
		log.Printf("-------------------- 头条 -- %s --------------------\n", tag)

		log.Println(r.URL)
		r.Headers.Set("accept", "text/javascript, text/html, application/xml, text/xml, */*")
		r.Headers.Set("content-type", "application/x-www-form-urlencoded")
		r.Headers.Set("referer", fmt.Sprintf("https://www.toutiao.com/ch/%s/", tag))
		r.Headers.Set("cookie", "tt_webid=6675124015910422029; Domain=.toutiao.com; expires=Tue, 09-Jul-2019 14:25:19 GMT; Max-Age=7804800; Path=/")
	})

	c.OnError(func(r *colly.Response, e error) {
		log.Fatalln(r.StatusCode, e)
	})

	c.OnScraped(func(r *colly.Response) {

		if r.StatusCode != 200 {
			log.Println("failed statuscode =", r.StatusCode)
			return
		}
		var food Food
		if err := json.Unmarshal(r.Body, &food); err != nil {
			log.Println(err)
			return
		}

		if food.Message != "success" {
			log.Fatal("获取数据错误", food.Message)
		}

		cnt := 0
		for i, data := range food.Data {
			data.SourceURL = "https://www.toutiao.com" + data.SourceURL
			if data.Tag != tag {
				log.Println(data.Tag, "这条数据不是该板块数据", data.SourceURL)
				continue
			}
			if data.ArticleGenre != "article" {
				log.Println(data.Tag, "这条数据不是文章类型", data.ArticleGenre)
				continue
			}
			temp := struct {
				Title string
				Link  string
			}{
				Title: strings.TrimSpace(data.Title),
				Link:  strings.TrimSpace(data.SourceURL),
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
			cnt++
		}

		log.Println("count", cnt, "has more", food.HasMore)

		time.Sleep(500 * time.Millisecond)

		// 下一次抓取
		v, err := vm.Call("getHoney", nil)
		if err != nil {
			log.Fatal(err)
		}
		asCP, _ := v.Export()
		asCPMap := asCP.(map[string]interface{})
		cp := asCPMap["cp"].(string)
		as := asCPMap["as"].(string)

		l := fmt.Sprintf("https://www.toutiao.com/api/pc/feed/?category=%s&utm_source=toutiao&widen=1&max_behot_time=%d&max_behot_time_tmp=%d&tadrequire=true&as=%s&cp=%s",
			tag, food.Next.MaxBehotTime, food.Next.MaxBehotTime, as, cp)

		r.Request.Visit(l)
	})

	v, err := vm.Call("getHoney", nil)
	if err != nil {
		log.Fatal(err)
	}
	asCP, _ := v.Export()
	asCPMap := asCP.(map[string]interface{})
	cp := asCPMap["cp"].(string)
	as := asCPMap["as"].(string)

	l := fmt.Sprintf("https://www.toutiao.com/api/pc/feed/?category=%s&utm_source=toutiao&widen=1&max_behot_time=0&max_behot_time_tmp=0&tadrequire=true&as=%s&cp=%s&_signature=X.g8yQAAA1JWtEbc-Ecn0V.4PN",
		tag, as, cp)

	c.Visit(l)
}

// 头条号
func (s *Toutiao) spiderMPColly(mid int64, tag, rdsTag string) {

	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:66.0) Gecko/20100101 Firefox/66.0"

	c.OnRequest(func(r *colly.Request) {
		log.Printf("-------------------- 头条 -- %s --------------------\n", tag)

		log.Println(r.URL)
		r.Headers.Set("Accept", "application/json, text/javascript")
		r.Headers.Set("Content-Type", "application/x-www-form-urlencoded")
		r.Headers.Set("Host", "www.toutiao.com")
		r.Headers.Set("Referer", fmt.Sprintf("https://www.toutiao.com/c/user/%d/", mid))
	})

	c.OnError(func(r *colly.Response, e error) {
		log.Fatalln(r.StatusCode, e)
	})

	c.OnScraped(func(r *colly.Response) {
		if r.StatusCode != 200 {
			log.Println("failed statuscode =", r.StatusCode)
			return
		}
		var food Food
		if err := json.Unmarshal(r.Body, &food); err != nil {
			log.Println(err)
			return
		}

		if food.Message != "success" {
			log.Fatal("获取数据错误", food.Message)
		}

		cnt := 0
		for i, data := range food.Data {
			data.SourceURL = "https://www.toutiao.com" + data.SourceURL
			if data.Tag != tag {
				log.Println(data.Tag, "这条数据不是该板块数据", data.SourceURL)
				continue
			}
			if data.ArticleGenre != "article" {
				log.Println(data.Tag, "这条数据不是文章类型", data.ArticleGenre)
				continue
			}
			temp := struct {
				Title string
				Link  string
			}{
				Title: strings.TrimSpace(data.Title),
				Link:  strings.TrimSpace(data.SourceURL),
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
			cnt++
		}
		if !food.HasMore {
			log.Println("数据抓取完成")
			os.Exit(0)
		}

		log.Println("count", cnt, "has more", food.HasMore)

		time.Sleep(500 * time.Millisecond)

		// 下一次抓取
		signature := strings.TrimSpace(s.getSign(fmt.Sprintf("%d%d", mid, food.Next.MaxBehotTime)))
		l := fmt.Sprintf("https://www.toutiao.com/c/user/article/?page_type=1&user_id=%d&max_behot_time=%d&count=20&_signature=%s", mid, food.Next.MaxBehotTime, signature)
		if err := r.Request.Visit(l); err != nil {
			log.Fatal(err)
		}
	})

	signature := strings.TrimSpace(s.getSign(fmt.Sprintf("%d%d", mid, time.Now().UnixNano()/1e6)))
	log.Println(signature)

	//signature := "XMFl2xAdAG2wZAq9GpWQIFzBZc"
	l := fmt.Sprintf("https://www.toutiao.com/c/user/article/?page_type=1&user_id=%d&max_behot_time=0&count=20&_signature=%s", mid, signature)

	if err := c.Visit(l); err != nil {
		log.Fatal(err)
	}
}

func (s *Toutiao) getSign(tmp string) string {
	cmd := exec.Command("node", "toutiao.js", fmt.Sprintf("%d", tmp))
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	// 保证关闭输出流
	defer stdout.Close()
	// 运行命令
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	// 读取输出结果
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	}

	return string(opBytes)
}

// 头条号抓取, 有头浏览器
func (s *Toutiao) spiderChrome(mid int64, tag, rdsTag string) {
	// create context
	//ctxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	runnerOps := chromedp.WithRunnerOptions(
		// 启动chrome的时候不检查默认浏览器
		//runner.Flag("no-default-browser-check", true),
		// 启动chrome 不适用沙盒, 性能优先
		runner.Flag("no-sandbox", true),
		// 设置浏览器窗口尺寸,
		//runner.WindowSize(1280, 1024),
		// 设置浏览器的userage
		runner.UserAgent(`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
		// 启动headless
		//runner.Flag("headless", true),
	)

	// 在普通模式的情况下启动chrome程序,并且建立共代码和chrome程序的之间的连接(https://127.0.0.1:9222)
	//chromedp.WithLog(log.Printf)
	cdp, err := chromedp.New(ctxt, runnerOps)
	if err != nil {
		log.Fatal(err)
	}

	// 获取第一页的标签ID
	homeTID := cdp.ListTargets()[0]
	log.Println("home target id", homeTID)

	log.Println("open home")
	homeURL := fmt.Sprintf("https://www.toutiao.com/c/user/%d/#mid=%d", mid, mid)
	handleErr(cdp.Run(ctxt, chromedp.Tasks{
		chromedp.Navigate(homeURL),
		chromedp.WaitReady(".relatedFeed", chromedp.ByQuery),
	}))
	var behotTime int64
	pageIndex := 0
	for {
		log.Println("get sign & honey")
		handleErr(cdp.SetHandlerByID(homeTID))
		log.Println("change handle home", homeTID)
		var sign string
		var honey interface{}
		handleErr(cdp.Run(ctxt, actionParam(mid, behotTime, &sign, &honey)))
		honeyMap := honey.(map[string]interface{})
		log.Println(sign, honeyMap)

		log.Println("open new data target page")
		var dataTID string
		handleErr(cdp.NewTarget(&dataTID).Do(ctxt, cdp.GetHandlerByID(homeTID)))

		log.Println("change handle data", dataTID)
		handleErr(cdp.SetHandlerByID(dataTID))

		dataURL := fmt.Sprintf("https://www.toutiao.com/c/user/article/?page_type=1&user_id=%d&max_behot_time=%d&count=20&as=%s&cp=%s&_signature=%s",
			mid, behotTime, honeyMap["as"], honeyMap["cp"], sign)
		log.Println("open data url", dataURL)

		var bodyStr string
		handleErr(cdp.Run(ctxt, chromedp.Tasks{
			chromedp.Navigate(dataURL),
			chromedp.WaitReady("body", chromedp.ByQuery),
			chromedp.Text("body", &bodyStr),
		}))

		// 解析数据
		var food Food
		if err := json.Unmarshal([]byte(bodyStr), &food); err != nil {
			log.Fatalln(err)
		}
		if err = processData(&food, tag, rdsTag); err != nil {
			log.Fatalln(err)
		}

		behotTime = food.Next.MaxBehotTime

		pageIndex++

		log.Printf("------------------------------ %d ------------------------------\n", pageIndex)
		_, err = target.CloseTarget(target.ID(dataTID)).Do(ctxt, cdp.GetHandlerByID(dataTID))
		handleErr(err)

		if !food.HasMore {
			break
		}

		time.Sleep(200 * time.Millisecond)
	}

	log.Println("crawler done", mid)

	if err = cdp.Shutdown(ctxt); err != nil {
		log.Fatalln(err)
	}
}

func actionCrawlerData(res *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitReady("body", chromedp.BySearch),
		chromedp.Text("body", res, chromedp.BySearch),
	}
}

func actionParam(mid, t int64, sign *string, honey interface{}) chromedp.Tasks {
	scriptSign := fmt.Sprintf("TAC.sign('%d%d')", mid, t)
	return chromedp.Tasks{
		chromedp.EvaluateAsDevTools(scriptSign, sign),
		chromedp.EvaluateAsDevTools("ascp.getHoney()", honey),
	}
}

func processData(food *Food, tag, rdsTag string) error {
	if food == nil {
		return fmt.Errorf("failed to nil")
	}
	if food.Message != "success" {
		return fmt.Errorf("error: %s", food.Message)
	}

	for i, data := range food.Data {
		data.Title = strings.TrimSpace(data.Title)
		data.SourceURL = strings.TrimSpace(data.SourceURL)
		data.SourceURL = "https://www.toutiao.com" + data.SourceURL
		if data.Tag != tag {
			log.Println(i, data.Tag, "这条数据不是该板块数据", data.SourceURL)
			continue
		}
		if data.ArticleGenre != "article" {
			log.Println(i, data.Tag, "这条数据不是文章类型", data.ArticleGenre)
			continue
		}

		temp := struct {
			Title string
			Link  string
		}{
			Title: strings.TrimSpace(data.Title),
			Link:  strings.TrimSpace(data.SourceURL),
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
	return nil
}
