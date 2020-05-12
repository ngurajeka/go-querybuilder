package condition_test

import (
	"testing"

	"github.com/ngurajeka/go-querybuilder/v2/condition"
)

func TestSimpleCondition(t *testing.T) {
	var (
		c                condition.Condition
		expected, result string
	)
	c = condition.Default("username", "ngurajeka")
	expected = "username = 'ngurajeka'"

	result = c.String(false)
	if result != expected {
		t.Errorf("invalid operation result, got %s\n", result)
	}

	c = condition.Default("id", 10)
	expected = "id = 10"

	result = c.String(false)
	if result != expected {
		t.Errorf("invalid operation result, got %s\n", result)
	}
}

func TestAdvancedCondition(t *testing.T) {
	var (
		c                condition.Condition
		expected, result string
	)

	c = condition.New(
		"username",
		condition.IN,
		condition.AND,
		[]string{"ady", "ngurajeka"},
	)
	expected = "username IN ( 'ady', 'ngurajeka' )"
	result = c.String(false)
	if result != expected {
		t.Errorf("invalid operation result, got %s\n", result)
	}

	c = condition.New(
		"id",
		condition.IN,
		condition.AND,
		[]int{10, 11},
	)
	expected = "id IN ( 10, 11 )"
	result = c.String(false)
	if result != expected {
		t.Errorf("invalid operation result, got %s\n", result)
	}

	c = condition.Default("id", []int{10, 11})
	result = c.String(false)
	if result != expected {
		t.Errorf("invalid operation result, got %s\n", result)
	}
}

func TestNilValue(t *testing.T) {
	var (
		c                condition.Condition
		expected, result string
	)

	c = condition.Nil("last_name")
	expected = "last_name IS NULL"
	result = c.String(false)
	if result != expected {
		t.Errorf("invalid operation result, got %s\n", result)
	}

	c = condition.Default("last_name", nil)
	if result != expected {
		t.Errorf("invalid operation result, got %s\n", result)
	}
}
