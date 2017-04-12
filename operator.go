package querybuilder

const (
	BETWEEN = "BETWEEN"
	EQUALS  = "="
	GT      = ">"
	GTE     = ">="
	LT      = "<"
	LTE     = "<="
	IN      = "IN"
	IS      = "IS"
	NOT     = "<>"
	NOTIN   = "NOT IN"
	LIKE    = "LIKE"

	DefaultOperator = EQUALS
)

func GetOperators() map[string]string {
	var operators = map[string]string{
		"equals": EQUALS,
		"gt":     GT,
		"gte":    GTE,
		"lt":     LT,
		"lte":    LTE,
		"in":     IN,
		"nin":    NOTIN,
		"not":    NOT,
		"like":   LIKE,
	}
	return operators
}

func Operator(opr string) string {
	var operators = GetOperators()
	if operator, ok := operators[opr]; ok {
		return operator
	}
	return DefaultOperator
}
