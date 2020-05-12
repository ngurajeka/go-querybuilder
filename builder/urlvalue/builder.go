package urlvalue

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"

	querybuilder "github.com/ngurajeka/go-querybuilder/v2"
	"github.com/ngurajeka/go-querybuilder/v2/condition"
	"github.com/ngurajeka/go-querybuilder/v2/order"
)

type builder struct {
	values url.Values
}

// New create new Builder from url.Values
func New(values url.Values) querybuilder.Builder {
	return &builder{values: values}
}

func (b *builder) Compile() querybuilder.QueryBuilder {
	var qb = querybuilder.NewQuerybuilder(0, 10)
	qb = b.filter(qb)
	qb = b.order(qb)
	qb = b.pagination(qb)
	return qb
}

func (b *builder) filter(qb querybuilder.QueryBuilder) querybuilder.QueryBuilder {
	rf := regexp.MustCompile(`filter`)
	rw := regexp.MustCompile(`\W`)
	for q, v := range b.values {
		if rf.MatchString(q) != true {
			continue
		}
		qSplit := rf.Split(q, -1)
		fSplit := rw.Split(qSplit[1], -1)

		opr := condition.DefaultOperator
		conj := condition.DefaultConjunction
		for i, val := range fSplit {
			switch {
			case i == 3 && val != "":
				opr = condition.Operator(val)
			case i == 5 && val == "or":
				conj = condition.OR
			}
		}
		if opr == condition.IN || opr == condition.NOTIN {
			filter := condition.New(fSplit[1], opr, conj, strings.Split(v[0], ","))
			qb.AddCondition(filter)
			continue
		}
		filter := condition.New(fSplit[1], opr, conj, v[0])
		qb.AddCondition(filter)
	}
	return qb
}

func (b *builder) pagination(qb querybuilder.QueryBuilder) querybuilder.QueryBuilder {
	var (
		limit  = 10
		number = 1
		offset = 0
	)
	pageNumber := b.values.Get("pageNumber")
	pageSize := b.values.Get("pageSize")
	if pageSize != "" {
		_limit, err := strconv.Atoi(pageSize)
		if err == nil {
			limit = _limit
		}
	}
	if pageNumber != "" {
		_number, err := strconv.Atoi(pageNumber)
		if err == nil {
			number = _number
			offset = (number - 1) * limit
		}
	}

	qb.SetLimit(limit)
	qb.SetNumber(number)
	qb.SetOffset(offset)
	return qb
}

func (b *builder) order(qb querybuilder.QueryBuilder) querybuilder.QueryBuilder {
	var (
		orderType order.Order
		sort      string
	)
	sort = b.values.Get("sort")
	if sort == "" {
		return qb
	}

	switch sort[:1] {
	case "+":
		orderType = order.Ascending(sort[1:])
	case "-":
		orderType = order.Descending(sort[1:])
	default:
		orderType = order.Ascending(sort)
	}

	qb.AddOrder(orderType)
	return qb
}
