package spider

import (
	"encoding/json"
	"go-film-demo/dao"
	"go-film-demo/model/news"
	"log"
	"time"
)

var NewsDao = dao.NewNewsRepository()

func CollectNews() {
	r := &RequestInfo{
		Uri:    "https://i.news.qq.com/web_feed/getHotModuleList",
		Params: nil,
		Header: nil,
		Resp:   nil,
		Err:    "",
		Body: Request{
			BaseReq: BaseReq{
				From: "pc",
			},
			Forward:   "2",
			FlushNum:  1,
			ChannelID: "news_news_top",
			ItemCount: 20,
		},
	}
	Post(r)
	var newsResp NewsResponse
	err := json.Unmarshal(r.Resp, &newsResp)
	if err != nil {
		log.Fatal(err.Error())
	}

	newsList := make([]*news.News, 0)
	for _, data := range newsResp.Data {
		publishTime, err := time.ParseInLocation("2006-01-02 15:04:05", data.PublishTime, time.Local)
		if err != nil {
			continue
		}
		newItem := &news.News{
			NewsID:      data.ID,
			Title:       data.Title,
			Creator:     data.MediaInfo.ChlName,
			Source:      data.LinkInfo.URL,
			Content:     data.Desc,
			PublishTime: publishTime,
			Cover:       data.PicInfo.BigImg[0],
			Desc:        data.Desc,
			Link:        data.LinkInfo.URL,
			CommentNum:  data.InteractionInfo.CommetNum,
			ReadNum:     data.InteractionInfo.ReadNum,
			LikeNum:     data.InteractionInfo.LikeNum,
			CollectNum:  data.InteractionInfo.CollectNum,
			ShareNum:    data.InteractionInfo.ShareNum,
		}
		newsList = append(newsList, newItem)
	}
	err = NewsDao.CreateBatch(newsList)
	if err != nil {
		log.Fatal(err.Error())
	}
}
