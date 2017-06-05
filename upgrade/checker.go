package upgrade

type Checker struct {
	Rules RuleList
	Info  *UpgradeInfo
}

func (c *Checker) Test(paras map[string]string) bool {
	for _, rule := range c.Rules {
		if rule.Test(paras) == true {
			return true
		}
	}

	return false
}

type CheckerList []*Checker

func (ckl CheckerList) Test(paras map[string]string) *Checker {
	for _, checker := range ckl {
		if checker.Test(paras) == true {
			return checker
		}
	}
	return nil
}
