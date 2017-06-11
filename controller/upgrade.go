package controller

import (
	"customized_lp/upgrade_server/upgrade"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
)

// Upgrade 检测升级
func Upgrade(c *gin.Context) {
	paras := upgrade.ParaList{}

	// values, _ := url.ParseQuery(c.Request.URL.RawQuery)
	values := c.Request.URL.Query()
	for key, values := range values {
		value := ""
		if len(values) > 0 {
			value = values[0]
		}
		paras = append(paras, upgrade.NewPara(key, value))
	}
	spew.Dump(paras)
	var ui *upgrade.UpgradeInfo
	checker := checkers.Test(paras...)
	if checker != nil {
		ui = checker.Info
	}
	c.JSON(http.StatusOK, ui)
}

// // UpgradeAndroid 检测安卓的升级
// func UpgradeAndroid(c *gin.Context) {
// 	paras := make(map[string]string)
// 	for _, param := range c.Params {
// 		paras[param.Key] = param.Value
// 	}

// 	var ui *upgrade.UpgradeInfo
// 	checker := andoidCheckers.Test(paras)
// 	if checker != nil {
// 		ui = checker.Info
// 	}
// 	c.JSON(http.StatusOK, ui)
// }

// // UpgradeIOS 检测 IOS 的升级
// func UpgradeIOS(c *gin.Context) {
// 	paras := make(map[string]string)
// 	for _, param := range c.Params {
// 		paras[param.Key] = param.Value
// 	}

// 	var ui *upgrade.UpgradeInfo
// 	checker := andoidCheckers.Test(paras)
// 	if checker != nil {
// 		ui = checker.Info
// 	}

// 	c.JSON(http.StatusOK, ui)
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
