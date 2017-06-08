package upgrade

type Checker struct {
	Name string
	Rule *Rule
	Info *UpgradeInfo
}

func NewChecker(name string, rule *Rule, info *UpgradeInfo) *Checker {
	return &Checker{
		Name: name,
		Rule: rule,
		Info: info,
	}
}

func (c *Checker) Test(paras map[string]string) bool {
	return c.Rule.Test(paras)
	// for _, rule := range c.Rules {
	// 	if rule.Test(paras) == true {
	// 		return true
	// 	}
	// }

	// return false
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
