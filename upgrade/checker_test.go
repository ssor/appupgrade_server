package upgrade

import (
	"testing"
)

func TestCheckers(t *testing.T) {
	para1 := NewPara("name", "android1")
	para2 := NewPara("version", "1.0")
	para3 := NewPara("name", "android2")

	android1Rules := RuleList{NewRule(para1, OpEqual), NewRule(para2, OpEqual)}

	android2Rules := RuleList{NewRule(para3, OpEqual), NewRule(para2, OpEqual)}

	android1UpgradeInfo := NewUpgradeInfo("1", "http1", "md555", "1.1")
	android2UpgradeInfo := NewUpgradeInfo("2", "http1", "md555", "1.2")

	android1Checker := NewChecker("1", android1Rules, android1UpgradeInfo)
	android2Checker := NewChecker("2", android2Rules, android2UpgradeInfo)

	androidCheckers := CheckerList{android1Checker, android2Checker}

	resultCheck := androidCheckers.Test(para2, para3)
	if resultCheck.Info.Version != "1.2" {
		t.Fatal("should be 1.2")
	}

	resultCheck = androidCheckers.Test(para1, para2)
	if resultCheck.Info.Version != "1.1" {
		t.Fatal("should be 1.1")
	}
}
