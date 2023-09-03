package controller

import (
	"TinyTikTok/models"
	"TinyTikTok/service/impl"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
type FeedResponse struct {
	models.Response
	VideoList []models.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, FeedResponse{
		Response:  models.Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
*/

func Feed(c *gin.Context) {

	//获取请求中的时间参数，如果没有，就将当前时间赋值给查询条件
	lastTime := c.Query("latest_time")
	latestTime, err := impl.FeedServiceImpl{}.GetFeedLatestTime(lastTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 1,
			"status_msg":  "Failed to get videos",
		})
		return
	}

	//查询视频video及user信息
	feedResponseVideoInfos, nextTime, err := impl.FeedServiceImpl{}.GetFeedByLatestTime(latestTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 1,
			"status_msg":  "Failed to get videos",
		})
		return
	}

	//返回视频结果
	c.JSON(http.StatusOK, models.FeedResponse{
		NextTime:   nextTime,
		StatusCode: 0,
		StatusMsg:  "Success",
		VideoList:  feedResponseVideoInfos,
	})

}
