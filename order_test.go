package querybuilder_test

import (
	"testing"

	querybuilder "github.com/ngurajeka/go-querybuilder"
)

func TestSimpleOrder(t *testing.T) {
	f := "userId"
	order := querybuilder.Ascending(f)
	if order.String(false) != "userId ASC" {
		t.Fatal("result should be userId ASC")
	}
	order = querybuilder.Descending(f)
	if order.String(false) != "userId DESC" {
		t.Fatal("result should be userId DESC")
	}
}

func TestWithQuery(t *testing.T) {
	q := querybuilder.NewQuerybuilder(0, 10)
	f := "userId"
	order := querybuilder.Ascending(f)
	if order.String(false) != "userId ASC" {
		t.Fatal("result should be userId ASC")
	}
	q.AddOrderIfNotExist(order)
	if q.StringifyOrder() != "userId ASC" {
		t.Fatal("result should be userId ASC")
	}
	order = querybuilder.Descending(f)
	if order.String(false) != "userId DESC" {
		t.Fatal("result should be userId DESC")
	}
	if q.StringifyOrder() != "userId ASC" {
		t.Fatal("result should be userId ASC")
	}
}
