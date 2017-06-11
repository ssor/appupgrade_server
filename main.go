package main

import (
	"customized_lp/upgrade_server/controller"
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/ssor/config"
)

var (
	configFile = flag.String("config", "conf/config.json", "config file for system")
	conf       config.IConfigInfo
)

func main() {
	var err error
	conf, err = config.LoadConfig(*configFile)
	if err != nil {
		panic("配置文件加载错误: " + err.Error())
	}

	router := gin.Default()

	router.Static("/css", "./static/css")
	router.Static("/img", "./static/img")
	router.Static("/js", "./static/js")
	router.Static("/fonts", "./static/fonts")
	router.Static("/bootstrap", "./static/bootstrap")

	router.LoadHTMLGlob("views/*.html")

	//************* api *****************
	router.POST("/api/v1/upgrade/addChecker", controller.AddChecker)

	// //android api
	// router.GET("/upgrade_android", controller.UpgradeAndroid)
	// router.OPTIONS("/api/v1/upgrade1UPGRADE/updateVersion", addCrossHeader)
	// router.POST("/api/v1/upgrade/updateVersion", controller.UpdateVersion)

	// //ios api
	// router.GET("/upgradeIOS", controller.UpgradeIOS)
	// router.OPTIONS("/api/v1/upgrade1UPGRADE/updateVersionIOS", addCrossHeader)
	// router.POST("/api/v1/upgrade/updateVersionIOS", controller.UpdateVersionIOS)

	// router.GET("/api/v1/upgrade1UPGRADE/allInfo", controller.AllInfo)

	//************* views *****************
	router.GET("/admin", controller.Admin)
	router.GET("/weixinscan", controller.Weixinscan)
	router.GET("/downloadApp", controller.DownloadApp)

	router.Run(":" + conf.Get("listeningPort").(string))
}

func addCrossHeader(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "X-Requested-With,X_Requested_With,Content-Type")
}
