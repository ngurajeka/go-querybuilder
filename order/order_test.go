package order_test

import (
	"testing"

	"github.com/ngurajeka/go-querybuilder"
	"github.com/ngurajeka/go-querybuilder/v2/order"
)

func TestSimpleOrder(t *testing.T) {
	f := "userId"
	v := order.Ascending(f)
	if v.String(false) != "userId ASC" {
		t.Fatal("result should be userId ASC")
	}
	v = order.Descending(f)
	if v.String(false) != "userId DESC" {
		t.Fatal("result should be userId DESC")
	}
}

func TestWithQuery(t *testing.T) {
	q := querybuilder.NewQuerybuilder(0, 10)
	f := "userId"
	v := order.Ascending(f)
	if v.String(false) != "userId ASC" {
		t.Fatal("result should be userId ASC")
	}
	q.AddOrderIfNotExist(v)
	if q.StringifyOrder() != "userId ASC" {
		t.Fatal("result should be userId ASC")
	}
	v = order.Descending(f)
	if v.String(false) != "userId DESC" {
		t.Fatal("result should be userId DESC")
	}
	if q.StringifyOrder() != "userId ASC" {
		t.Fatal("result should be userId ASC")
	}
}
