package controllers

import (
	"net/http"

	"customized_lp/upgrade_server/upgrade"

	"github.com/gin-gonic/gin"
)

// AddChecker update升级
func AddChecker(c *gin.Context) {
	var upload struct {
		Name         string            `json:"name"`
		Requirements map[string]string `json:"requirements"`

		Tip     string `json:"tip"`
		URL     string `json:"url"`
		MD5     string `json:"md5"`
		Version string `json:"version"`
	}

	err := c.BindJSON(&upload)
	if err != nil {
		c.JSON(http.StatusBadRequest, "format error")
		return
	}

	rule := upgrade.NewRule(upload.Requirements)
	upgradeInfo := upgrade.NewUpgradeInfo(upload.Tip, upload.URL, upload.MD5, upload.Version)
	newChecker := upgrade.NewChecker(upload.Name, rule, upgradeInfo)
	checkers = append(checkers, newChecker)
	err = saveCheckers(checkers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "system error")
		return
	}

	c.JSON(http.StatusOK, "OK")
}

// // UpdateIOS update IOS 的升级
// func UpdateIOS(c *gin.Context) {
// 	c.JSON(http.StatusOK, "OK")
// }

// func AllInfo(c *gin.Context) {
// 	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 	c.JSON(http.StatusOK, gin.H{
// 		"code":    0,
// 		"message": "",
// 		"data":    nil,
// 	})
// 	return
// }
