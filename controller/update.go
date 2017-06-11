package controller

import (
	"net/http"

	"customized_lp/upgrade_server/upgrade"

	"github.com/gin-gonic/gin"
)

type Rule struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	Operator string `json:"operator"`
}

type UpgradeRule struct {
	Name  string  `json:"name"`
	Rules []*Rule `json:"rules"`

	Tip     string `json:"tip"`
	URL     string `json:"url"`
	MD5     string `json:"md5"`
	Version string `json:"version"`
}

// AddChecker update升级
func AddChecker(c *gin.Context) {
	var upload UpgradeRule

	err := c.BindJSON(&upload)
	if err != nil {
		c.JSON(http.StatusBadRequest, "format error")
		return
	}

	newChecker := parseToChecker(&upload)
	checkers = checkers.Update(newChecker)
	err = saveCheckers(checkers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "system error")
		return
	}

	c.JSON(http.StatusOK, "OK")
}

func parseToChecker(upgradeRule *UpgradeRule) *upgrade.Checker {
	rules := upgrade.RuleList{}
	for _, rule := range upgradeRule.Rules {
		rules = append(rules, upgrade.NewRule(upgrade.NewPara(rule.Key, rule.Value), upgrade.Operator(rule.Operator)))
	}
	upgradeInfo := upgrade.NewUpgradeInfo(upgradeRule.Tip, upgradeRule.URL, upgradeRule.MD5, upgradeRule.Version)
	newChecker := upgrade.NewChecker(upgradeRule.Name, rules, upgradeInfo)
	return newChecker
}

// Checkers return all checkers
func Checkers(c *gin.Context) {
	c.JSON(http.StatusOK, checkers)
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
