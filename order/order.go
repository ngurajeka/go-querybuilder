package order

import "fmt"

const (
	// ASC order by Ascending
	ASC = "ASC"
	// DESC order by Descending
	DESC = "DESC"
	// DefaultOrder default to Ascending
	DefaultOrder = ASC
)

// Order store ordering config
type Order interface {
	Field() string
	Order() string
	String(bool) string
}

type order struct {
	field, order string
}

// Ascending create new ascending order with field as parameter
func Ascending(f string) Order {
	return &order{f, ASC}
}

// Descending create new descending order with field as parameter
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
