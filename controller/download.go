package controller

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
