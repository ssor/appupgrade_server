package upgrade

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

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

func (r *Rule) Test(paras []*Para) bool {
	for _, para := range paras {
		if para.Key == r.Para.Key {
			return r.Op.compareValue(para.Value, r.Para.Value)
		}
	}
	return false
	// if r.Para.Key != para.Key {
	// 	return false
	// }
}

type RuleList []*Rule

func NewRuleList(rules ...*Rule) RuleList {
	return append(RuleList{}, rules...)
}

func (rl RuleList) Add(rule *Rule) RuleList {
	return append(rl, rule)
}
func (rl RuleList) Test(paras []*Para) bool {
	for _, rule := range rl {
		if rule.Test(paras) == false {
			fmt.Println("rule test failed: ")
			spew.Dump(rule)
			// fmt.Println("paras: ", paras)
			return false
		}
	}
	return true
}
