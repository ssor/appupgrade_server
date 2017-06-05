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

	controllers.InitSys(conf)
	router := gin.Default()

	router.Static("/api/v1/upgrade1UPGRADE/css", "./static/css")
	router.Static("/api/v1/upgrade1UPGRADE/img", "./static/img")
	router.Static("/api/v1/upgrade1UPGRADE/js", "./static/js")
	router.Static("/api/v1/upgrade1UPGRADE/fonts", "./static/fonts")
	router.Static("/api/v1/upgrade1UPGRADE/bootstrap", "./static/bootstrap")

	router.LoadHTMLGlob("views/*.html")

	//************* api *****************

	//android api
	router.GET("/upgrade_android", controllers.UpgradeAndroid)
	router.OPTIONS("/api/v1/upgrade1UPGRADE/updateVersion", controllers.AddCrossHeader)
	router.POST("/api/v1/upgrade1UPGRADE/updateVersion", controllers.UpdateVersion)

	//ios api
	router.GET("/upgradeIOS", controllers.UpgradeIOS)
	router.OPTIONS("/api/v1/upgrade1UPGRADE/updateVersionIOS", controllers.AddCrossHeader)
	router.POST("/api/v1/upgrade1UPGRADE/updateVersionIOS", controllers.UpdateVersionIOS)

	router.GET("/api/v1/upgrade1UPGRADE/allInfo", controllers.AllInfo)

	//************* views *****************
	router.GET("/weixinscan", controllers.Weixinscan)
	router.GET("/downloadApp", controllers.DownloadApp)

	router.Run(":" + conf.GetListeningPort())
}
