package upgrade

import (
	"testing"
)

func TestCheckers(t *testing.T) {
	para1 := map[string]string{
		"name":    "android1",
		"version": "1.0",
	}
	para2 := map[string]string{
		"name":    "android2",
		"version": "1.0",
	}
	android1Rule := NewRule(para1)
	android2Rule := NewRule(para2)

	android1UpgradeInfo := NewUpgradeInfo("1", "http1", "md555", "1.1")
	android2UpgradeInfo := NewUpgradeInfo("2", "http1", "md555", "1.2")

	android1Checker := NewChecker("1", android1Rule, android1UpgradeInfo)
	android2Checker := NewChecker("2", android2Rule, android2UpgradeInfo)

	androidCheckers := CheckerList{android1Checker, android2Checker}

	resultCheck := androidCheckers.Test(para2)
	if resultCheck.Info.Version != "1.2" {
		t.Fatal("should be 1.2")
	}

	resultCheck = androidCheckers.Test(para1)
	if resultCheck.Info.Version != "1.1" {
		t.Fatal("should be 1.1")
	}
}
