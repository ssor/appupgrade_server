package upgrade

type Para struct {
	Key   string
	Value string
}

func NewPara(key, value string) *Para {
	return &Para{
		Key:   key,
		Value: value,
	}
}

type ParaList []*Para
