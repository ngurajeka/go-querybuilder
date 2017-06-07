package querybuilder

type QueryBuilder interface {
	AddCondition(Condition)
	AddOrder(Order)
	AddOrderIfNotExist(Order)
	Copy(QueryBuilder)
	SetOffset(int)
	SetLimit(int)
	SetNumber(int)

	Remove(string)
	RemoveAll()
	RemoveByPrefix(string)
	RemoveConditions()
	RemoveOrders()
	RemovePaging()

	Conditions() []Condition
	HasConditions() bool
	HasOrders() bool
	Limit() int
	Number() int
	Offset() int
	Orders() []Order
	Map() map[string]interface{}
	String(exclude ...string) string
	StringifyOrder() string
}

type querybuilder struct {
	conditions            []Condition
	orders                []Order
	limit, number, offset int
}

func NewQuerybuilder(offset, limit int) QueryBuilder {
	return &querybuilder{offset: offset, limit: limit}
}

func (qb *querybuilder) AddCondition(c Condition) {
	qb.conditions = append(qb.conditions, c)
}

func (qb *querybuilder) AddOrder(o Order) {
	qb.orders = append(qb.orders, o)
}

func (qb *querybuilder) AddOrderIfNotExist(o Order) {
	for _, order := range qb.orders {
		if order.Field() == o.Field() {
			return
		}
	}
	qb.AddOrder(o)
}

func (qb *querybuilder) Copy(another_qb QueryBuilder) {
	for _, order := range another_qb.Orders() {
		qb.AddOrderIfNotExist(order)
	}
	for _, condition := range another_qb.Conditions() {
		qb.AddCondition(condition)
	}
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

func (qb *querybuilder) Remove(f string) {
	var conditions []Condition
	for _, condition := range qb.Conditions() {
		if condition.Field() != f {
			conditions = append(conditions, condition)
		}
	}
	qb.conditions = conditions
}

func (qb *querybuilder) RemoveByPrefix(p string) {
	length := len(p)
	var conditions []Condition
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

func (qb *querybuilder) RemoveAll() {
	qb.RemoveConditions()
	qb.RemoveOrders()
	qb.RemovePaging()
}

func (qb *querybuilder) RemoveConditions() {
	qb.conditions = make([]Condition, 0)
}

func (qb *querybuilder) RemoveOrders() {
	qb.orders = make([]Order, 0)
}

func (qb *querybuilder) RemovePaging() {
	qb.limit = 0
	qb.number = 0
	qb.offset = 0
}

func (qb *querybuilder) Conditions() []Condition {
	return qb.conditions
}

func (qb *querybuilder) Orders() []Order {
	return qb.orders
}

func (qb *querybuilder) HasConditions() bool {
	return len(qb.conditions) > 0
}

func (qb *querybuilder) HasOrders() bool {
	return len(qb.orders) > 0
}

func (qb *querybuilder) Number() int {
	return qb.number
}

func (qb *querybuilder) Offset() int {
	return qb.offset
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
