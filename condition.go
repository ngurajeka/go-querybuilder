package querybuilder

import (
	"fmt"
	"reflect"
)

const (
	AND = "AND"
	OR  = "OR"

	DefaultConjunction = AND
)

type Condition interface {
	Field() string
	Operator() string
	Value() interface{}
	Map() map[string]interface{}
	String(useConjunction bool) string
}

type filter struct {
	field, operator, conjunction string
	value                        interface{}
}

func New(f, o, c string, v interface{}) Condition {
	return &filter{
		field:       f,
		value:       v,
		operator:    o,
		conjunction: c,
	}
}

func Default(f string, v interface{}) Condition {
	return New(f, DefaultOperator, DefaultConjunction, v)
}

func (c *filter) Field() string {
	return c.field
}

func (c *filter) Operator() string {
	return c.operator
}

func (c *filter) Value() interface{} {
	return c.value
}

func (c *filter) Map() map[string]interface{} {
	return map[string]interface{}{
		"field":       c.field,
		"value":       c.value,
		"operator":    c.operator,
		"conjunction": c.conjunction,
	}
}

func (c *filter) stringify(str string, v interface{}) string {
	switch v.(type) {
	case int:
		str = fmt.Sprintf("%s %d", str, v.(int))
	case uint:
		str = fmt.Sprintf("%s %d", str, v.(uint))
	case uint64:
		str = fmt.Sprintf("%s %d", str, v.(uint64))
	case string:
		str = fmt.Sprintf("%s '%s'", str, v.(string))
	case nil:
		str = fmt.Sprintf("%s NULL")
	default:
		if reflect.ValueOf(v).Kind() == reflect.Slice {
			var newStr string
			for _, sv := range c.convertToSlice(v) {
				if newStr == "" {
					newStr = c.stringify(newStr, sv)
					continue
				}
				newStr = c.stringify(newStr+",", sv)
			}
			str = fmt.Sprintf("%s (%s )", str, newStr)
		}
	}
	return str
}

func (c *filter) convertToSlice(v interface{}) []interface{} {
	var newSlice []interface{}
	switch value := v.(type) {
	case []string:
		for _, s := range value {
			newSlice = append(newSlice, s)
		}
	case []uint:
		for _, u := range value {
			newSlice = append(newSlice, u)
		}
	case []uint64:
		for _, u := range value {
			newSlice = append(newSlice, u)
		}
	case []int:
		for _, i := range value {
			newSlice = append(newSlice, i)
		}
	}
	return newSlice
}

func (c *filter) String(useConjunction bool) (str string) {
	str = fmt.Sprintf("%s %s", c.field, c.operator)
	str = c.stringify(str, c.value)
	if useConjunction {
		str = fmt.Sprintf(" %s %s", c.conjunction, str)
	}
	return str
}
