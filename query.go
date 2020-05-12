package querybuilder

import (
	"fmt"
	"strings"

	"github.com/ngurajeka/go-querybuilder/v2/condition"
	"github.com/ngurajeka/go-querybuilder/v2/order"
)

// QueryBuilder store all filtering
type QueryBuilder interface {
	AddCondition(condition.Condition)
	AddOrder(order.Order)
	Conditions() []condition.Condition
	Copy(QueryBuilder)
	HasConditions() bool
	HasOrders() bool
	Limit() int
	Map() map[string]interface{}
	Number() int
	Offset() int
	Orders() []order.Order
	PrepareStatement() (string, []interface{})
	Remove(string)
	RemoveByPrefix(string)
	RemoveConditions()
	RemoveOrders()
	RemovePaging()
	Reset()
	SetLimit(int)
	SetNumber(int)
	SetOffset(int)
	String(exclude ...string) string
	StringifyOrder() string
}

type querybuilder struct {
	conditions            []condition.Condition
	orders                []order.Order
	limit, number, offset int
}

// NewQuerybuilder create new querybuilder with offset and limit option
func NewQuerybuilder(offset, limit int) QueryBuilder {
	return &querybuilder{offset: offset, limit: limit}
}

func (qb *querybuilder) AddCondition(c condition.Condition) {
	qb.conditions = append(qb.conditions, c)
}

func (qb *querybuilder) AddOrder(o order.Order) {
	for _, order := range qb.orders {
		if order.Field() == o.Field() {
			return
		}
	}
	qb.orders = append(qb.orders, o)
}

func (qb *querybuilder) Conditions() []condition.Condition {
	return qb.conditions
}

func (qb *querybuilder) Copy(anotherQb QueryBuilder) {
	for _, order := range anotherQb.Orders() {
		qb.AddOrder(order)
	}
	for _, condition := range anotherQb.Conditions() {
		qb.AddCondition(condition)
	}
}

func (qb *querybuilder) HasConditions() bool {
	return len(qb.conditions) > 0
}

func (qb *querybuilder) HasOrders() bool {
	return len(qb.orders) > 0
}

func (qb *querybuilder) Limit() int {
	return qb.limit
}

func (qb *querybuilder) Map() map[string]interface{} {
	return map[string]interface{}{
		"conditions": qb.conditions,
		"orders":     qb.orders,
		"limit":      qb.limit,
		"offset":     qb.offset,
	}
}

func (qb *querybuilder) Number() int {
	return qb.number
}

func (qb *querybuilder) Offset() int {
	return qb.offset
}

func (qb *querybuilder) Orders() []order.Order {
	return qb.orders
}

func (qb *querybuilder) PrepareStatement() (string, []interface{}) {
	var (
		q      []string
		values []interface{}
	)
	for _, v := range qb.conditions {
		if v.Operator() == condition.IN || v.Operator() == condition.NOTIN {
			q = append(q, fmt.Sprintf("%s %s (?)", v.Field(), v.Operator()))
		} else {
			q = append(q, fmt.Sprintf("%s %s ?", v.Field(), v.Operator()))
		}
		values = append(values, v.Value())
	}

	return strings.Join(q, " "), values
}

func (qb *querybuilder) Remove(f string) {
	var conditions []condition.Condition
	for _, condition := range qb.Conditions() {
		if condition.Field() != f {
			conditions = append(conditions, condition)
		}
	}
	qb.conditions = conditions
}

func (qb *querybuilder) RemoveByPrefix(p string) {
	length := len(p)
	var conditions []condition.Condition
	for _, condition := range qb.Conditions() {
		switch {
		case len(condition.Field()) < length:
			conditions = append(conditions, condition)
		case condition.Field()[:length] != p:
			conditions = append(conditions, condition)
		}
	}
	qb.conditions = conditions
}

func (qb *querybuilder) RemoveConditions() {
	qb.conditions = make([]condition.Condition, 0)
}

func (qb *querybuilder) RemoveOrders() {
	qb.orders = make([]order.Order, 0)
}

func (qb *querybuilder) RemovePaging() {
	qb.limit = 0
	qb.number = 0
	qb.offset = 0
}

func (qb *querybuilder) Reset() {
	qb.RemoveConditions()
	qb.RemoveOrders()
	qb.RemovePaging()
}

func (qb *querybuilder) SetLimit(limit int) {
	qb.limit = limit
}

func (qb *querybuilder) SetNumber(number int) {
	qb.number = number
}

func (qb *querybuilder) SetOffset(offset int) {
	qb.offset = offset
}

func (qb *querybuilder) String(exclude ...string) string {
	var str string
	for _, condition := range qb.conditions {
		if !isExist(condition.Field(), exclude) {
			str += condition.String((str != ""))
		}
	}
	return str
}

func (qb *querybuilder) StringifyOrder() string {
	var str string
	for _, order := range qb.orders {
		str += order.String((str != ""))
	}
	return str
}

func isExist(f string, v []string) bool {
	for _, s := range v {
		if s == f {
			return true
		}
	}
	return false
}
