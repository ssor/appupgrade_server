package controllers

import (
	"customized_lp/upgrade_server/upgrade"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpgradeAndroid 检测安卓的升级
func UpgradeAndroid(c *gin.Context) {
	paras := make(map[string]string)
	for _, param := range c.Params {
		paras[param.Key] = param.Value
	}

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var ui *upgrade.UpgradeInfo
	checker := andoidCheckers.Test(paras)
	if checker != nil {
		ui = checker.Info
	}
	c.JSON(http.StatusOK, ui)
}

// UpgradeIOS 检测 IOS 的升级
func UpgradeIOS(c *gin.Context) {
	paras := make(map[string]string)
	for _, param := range c.Params {
		paras[param.Key] = param.Value
	}

	var ui *upgrade.UpgradeInfo
	checker := andoidCheckers.Test(paras)
	if checker != nil {
		ui = checker.Info
	}

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, ui)
}

// func AllInfo(c *gin.Context) {
// 	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 	c.JSON(http.StatusOK, gin.H{
// 		"code":    0,
// 		"message": "",
// 		"data":    nil,
// 	})
// 	return
// }
