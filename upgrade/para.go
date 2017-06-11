package upgrade

type Para struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewPara(key, value string) *Para {
	return &Para{
		Key:   key,
		Value: value,
	}
}

type ParaList []*Para
