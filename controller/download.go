package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Weixinscan(c *gin.Context) {
	c.HTML(http.StatusOK, "weixinScan.html", nil)
}
func DownloadApp(c *gin.Context) {
	c.HTML(http.StatusOK, "appDownload.html", nil)
}

func AddCrossHeader(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "X-Requested-With,X_Requested_With,Content-Type")
}