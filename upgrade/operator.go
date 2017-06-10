package upgrade

type Operator string

const (
	OpEqual            = "equal"
	OpGreaterThan      = ">"
	OpGreaterEqualThan = ">="
	OpLessThan         = "<"
	OpLessEqualThan    = "<="
)

func (op Operator) compareValue(dest, src string) bool {
	switch op {
	case OpEqual:
		return dest == src
	case OpGreaterThan:
		return dest > src
	case OpGreaterEqualThan:
		return dest >= src
	case OpLessThan:
		return dest < src
	case OpLessEqualThan:
		return dest <= src
	}
	return false
}
