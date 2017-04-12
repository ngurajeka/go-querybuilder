package urlvalue_test

import (
	"net/url"
	"testing"

	querybuilder "github.com/ngurajeka/go-querybuilder"
	"github.com/ngurajeka/go-querybuilder/builder/urlvalue"
)

func TestUrlValueBuilder(t *testing.T) {
	var (
		b querybuilder.Builder
		// expected string
		qb     querybuilder.QueryBuilder
		values = url.Values{}
	)
	values.Add("pageSize", "20")
	values.Add("pageNumber", "3")
	values.Add("sort", "-id")
	b = urlvalue.New(values)
	qb = b.Compile()

	if qb.Offset() != 40 {
		t.Errorf("invalid offset, got %d\n", qb.Offset())
	}

	if qb.Limit() != 20 {
		t.Errorf("invalid limit, got %d\n", qb.Limit())
	}

	if qb.StringifyOrder() != "id DESC" {
		t.Errorf("invalid order, got %s\n", qb.StringifyOrder())
	}
}

func TestAdvancedUrlValueBuilder(t *testing.T) {
	var (
		b                querybuilder.Builder
		expected, result string
		qb               querybuilder.QueryBuilder
		values           = url.Values{}
	)
	values.Add("filter[username]", "ady")
	b = urlvalue.New(values)
	expected = "username = 'ady'"
	qb = b.Compile()

	result = qb.String()
	if result != expected {
		t.Errorf("invalid result, got %s\n", result)
	}

	values = url.Values{}
	values.Add("filter[username][in]", "ady,ngurajeka")
	values.Add("filter[id][in]", "10,11")
	values.Add("filter[age][gt]", "10")
	b = urlvalue.New(values)
	qb = b.Compile()

	result = qb.String()
	switch result {
	case "username IN ( 'ady', 'ngurajeka' ) AND id IN ( '10', '11' ) AND age > '10'":
	case "username IN ( 'ady', 'ngurajeka' ) AND age > '10' AND id IN ( '10', '11' )":
	case "age > '10' AND username IN ( 'ady', 'ngurajeka' ) AND id IN ( '10', '11' )":
	case "age > '10' AND id IN ( '10', '11' ) AND username IN ( 'ady', 'ngurajeka' )":
	case "id IN ( '10', '11' ) AND age > '10' AND username IN ( 'ady', 'ngurajeka' )":
	case "id IN ( '10', '11' ) AND username IN ( 'ady', 'ngurajeka' ) AND age > '10'":
		break
	default:
		t.Errorf("invalid result, got %s\n", result)
	}

	values = url.Values{}
	values.Add("filter[username][in]", "ady,ngurajeka")
	values.Add("filter[age]", "10")
	b = urlvalue.New(values)
	qb = b.Compile()

	result = qb.String()
	switch result {
	case "username IN ( 'ady', 'ngurajeka' ) AND age = '10'":
	case "age = '10' AND username IN ( 'ady', 'ngurajeka' )":
		break
	default:
		t.Errorf("invalid result, got %s\n", result)
	}
}
