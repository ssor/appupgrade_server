package upgrade

// type ParaList []*Para

// func (pl ParaList) Fit(paras map[string]string) bool {
// 	for k, v := range pl {
// 		value, exists := paras[k]
// 		if exists == false {
// 			return false
// 		}

// 		if v != value {
// 			return false
// 		}
// 	}

// 	return true
// }

type Rule struct {
	// Paras      ParaList
	Para *Para    `json:"para"`
	Op   Operator `json:"op"`
}

func NewRule(para *Para, op Operator) *Rule {
	return &Rule{
		Para: para,
		Op:   op,
	}
}

func (r *Rule) Test(para *Para) bool {
	if r.Para.Key != para.Key {
		return false
	}
	return r.Op.compareValue(r.Para.Value, para.Value)
}

type RuleList []*Rule

func NewRuleList(rules ...*Rule) RuleList {
	return append(RuleList{}, rules...)
}

func (rl RuleList) Add(rule *Rule) RuleList {
	return append(rl, rule)
}
func (rl RuleList) Test(para *Para) bool {
	for _, rule := range rl {
		if rule.Test(para) == true {
			return true
		}
	}
	return false
}
