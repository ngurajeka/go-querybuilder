package querybuilder

type Condition interface {
	Field() string
	Map() map[string]interface{}
	String(useConjunction bool) string
}
