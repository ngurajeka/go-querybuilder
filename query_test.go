package querybuilder_test

import (
	"testing"

	querybuilder "github.com/ngurajeka/go-querybuilder/v2"
	"github.com/ngurajeka/go-querybuilder/v2/condition"
)

func TestSimpleQuery(t *testing.T) {
	var (
		qb querybuilder.QueryBuilder
	)

	qb = querybuilder.NewQuerybuilder(0, 0)

	if qb.HasConditions() {
		t.Error("empty query should not have any conditions")
	}

	if qb.HasOrders() {
		t.Error("empty query should not have any orders")
	}

	if qb.Limit() != 0 {
		t.Error("invalid limit option")
	}

	if qb.Offset() != 0 {
		t.Error("invalid offset option")
	}
}

func TestPrepareStatement(t *testing.T) {
	var (
		qb querybuilder.QueryBuilder
	)

	qb = querybuilder.NewQuerybuilder(0, 0)
	qb.AddCondition(condition.Default("username", "ngurajeka"))

	stmt, args := qb.PrepareStatement()
	if stmt != "username = ?" {
		t.Errorf("invalid statement, got: %s\n", stmt)
	}
	if args[0] != "ngurajeka" {
		t.Errorf("invalid arguments, got: %s\n", args[0])
	}

	qb.AddCondition(condition.Default("gender", "m"))
	stmt, args = qb.PrepareStatement()
	if stmt != "username = ? AND gender = ?" {
		t.Errorf("invalid statement, got: %s\n", stmt)
	}
	if args[0] != "ngurajeka" {
		t.Errorf("invalid arguments, got: %s\n", args[0])
	}
	if args[1] != "m" {
		t.Errorf("invalid arguments, got: %s\n", args[1])
	}
}
