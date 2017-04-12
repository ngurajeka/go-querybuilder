package querybuilder_test

import (
	"testing"

	querybuilder "github.com/ngurajeka/go-querybuilder"
)

func TestSimpleCondition(t *testing.T) {
	var (
		c                querybuilder.Condition
		expected, result string
	)
	c = querybuilder.Default("username", "ngurajeka")
	expected = "username = 'ngurajeka'"

	result = c.String(false)
	if result != expected {
		t.Errorf("invalid operation result, got %s\n", result)
	}

	c = querybuilder.Default("id", 10)
	expected = "id = 10"

	result = c.String(false)
	if result != expected {
		t.Errorf("invalid operation result, got %s\n", result)
	}
}

func TestAdvancedCondition(t *testing.T) {
	var (
		c                querybuilder.Condition
		expected, result string
	)
	c = querybuilder.New(
		"username",
		querybuilder.IN,
		querybuilder.AND,
		[]string{"ady", "ngurajeka"},
	)
	expected = "username IN ( 'ady', 'ngurajeka' )"

	result = c.String(false)
	if result != expected {
		t.Errorf("invalid operation result, got %s\n", result)
	}

	c = querybuilder.New(
		"id",
		querybuilder.IN,
		querybuilder.AND,
		[]int{10, 11},
	)
	expected = "id IN ( 10, 11 )"

	result = c.String(false)
	if result != expected {
		t.Errorf("invalid operation result, got %s\n", result)
	}
}
