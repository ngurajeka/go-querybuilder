package condition

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
	ILIKE   = "ILIKE"
	NULL    = "NULL"

	DefaultOperator = EQUALS
)

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
	"ilike":  ILIKE,
	"null":   NULL,
}

func Operator(opr string) string {
	if operator, ok := operators[opr]; ok {
		return operator
	}

	return DefaultOperator
}
