package upgrade

type Checker struct {
	Name  string       `json:"name"`
	Rules RuleList     `json:"rules"`
	Info  *UpgradeInfo `json:"info"`
}

func NewChecker(name string, rules RuleList, info *UpgradeInfo) *Checker {
	return &Checker{
		Name:  name,
		Rules: rules,
		Info:  info,
	}
}

func (c *Checker) Test(paras ...*Para) bool {
	for _, para := range paras {
		if c.Rules.Test(para) == false {
			return false
		}
	}
	return true
}

type CheckerList []*Checker

func (ckl CheckerList) Test(paras ...*Para) *Checker {
	for _, checker := range ckl {
		if checker.Test(paras...) == true {
			return checker
		}
	}
	return nil
}
