package spider

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

var (
	Client = CreateClient()
)

var RefererUrl string

const (
	ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36"
)

type RequestInfo struct {
	Uri    string      `json:"uri"`    // 请求url地址
	Params url.Values  `json:"param"`  // 请求参数
	Header http.Header `json:"header"` // 请求头数据
	Resp   []byte      `json:"resp"`   // 响应结果数据
	Err    string      `json:"err"`    // 错误信息
	Body   interface{} `json:"body"`
}

func CreateClient() *colly.Collector {
	collector := colly.NewCollector(
		colly.UserAgent(ua),
		colly.AllowURLRevisit(),
	)
	return collector
}

func Get(r *RequestInfo) {
	if r.Header != nil {
		if t, err := strconv.Atoi(r.Header.Get("timeout")); err != nil && t > 0 {
			Client.SetRequestTimeout(time.Duration(t) * time.Second)
		}
	}
	extensions.RandomUserAgent(Client)
	Client.OnResponse(func(response *colly.Response) {
		if (response.StatusCode == 200 || (response.StatusCode >= 300 && response.StatusCode <= 399)) && len(response.Body) > 0 {
			r.Resp = response.Body
		} else {
			r.Resp = []byte{}
		}
		RefererUrl = response.Request.URL.String()
	})
	err := Client.Visit(fmt.Sprintf("%s?%s", r.Uri, r.Params.Encode()))
	if err != nil {
		r.Err = err.Error()
		log.Println("获取数据失败: ", err)
	}
}

func Post(r *RequestInfo) {

	da, err := json.Marshal(r.Body)

	if err != nil {
		fmt.Println(err)
	}
	Client.OnResponse(func(response *colly.Response) {
		r.Resp = response.Body
	})
	Client.OnRequest(func(r *colly.Request) {
		fmt.Println(r)
		fmt.Println(r.Method)
		r.Headers.Set("Content-Type", "application/json;charset=UTF-8")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.113 Safari/537.36")
	})
	Client.OnError(func(response *colly.Response, e error) {
		fmt.Println(e)
	})
	Client.PostRaw(r.Uri, da)
	//c.Visit("http://www.××××.com:×××/baseDevice/getUserInfo")

}
