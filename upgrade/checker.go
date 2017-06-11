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
	return c.Rules.Test(paras)
	// for _, para := range paras {
	// 	if c.Rules.Test(para) == false {
	// 		return false
	// 	}
	// }
	// return true
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

func (ckl CheckerList) findByName(name string) *Checker {
	for _, checker := range ckl {
		if checker.Name == name {
			return checker
		}
	}
	return nil
}

func (ckl CheckerList) Update(checker *Checker) CheckerList {
	tempChecker := ckl.findByName(checker.Name)
	if tempChecker == nil {
		return append(ckl, checker)
	}

	tempChecker.Rules = checker.Rules
	tempChecker.Info = checker.Info

	return ckl
}
