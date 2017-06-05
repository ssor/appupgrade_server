package upgrade

type ParaList map[string]string

func (pl ParaList) Fit(paras map[string]string) bool {
	for k, v := range pl {
		value, exists := paras[k]
		if exists == false {
			return false
		}

		if v != value {
			return false
		}
	}

	return true
}

type Rule struct {
	Paras ParaList
}

func (r *Rule) Test(paras map[string]string) bool {
	return r.Paras.Fit(paras)
}

type RuleList []*Rule
