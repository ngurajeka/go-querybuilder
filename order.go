package querybuilder

import "fmt"

const (
	ASC  = "ASC"
	DESC = "DESC"

	DefaultOrder = ASC
)

type Order interface {
	Field() string
	Order() string
	String(bool) string
}

type order struct {
	field, order string
}

func Ascending(f string) Order {
	return &order{f, ASC}
}

func Descending(f string) Order {
	return &order{f, DESC}
}

func (o *order) Field() string {
	return o.field
}

func (o *order) Order() string {
	return o.order
}

func (o *order) String(useComma bool) (str string) {
	str = fmt.Sprintf("%s %s", o.field, o.order)
	if useComma {
		str = fmt.Sprintf(", %s", str)
	}
	return str
}
