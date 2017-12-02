package extractor

import (
	"fmt"

	querybuilder "github.com/ngurajeka/go-querybuilder"
	db "upper.io/db.v3"
)

func AsUpperDBCondition(q querybuilder.QueryBuilder) db.Cond {
	var cond db.Cond
	for _, cond := range q.Conditions() {
		field := fmt.Sprintf("%s %s", cond.Field(), cond.Operator())
		cond[field] = cond.Value()
	}
	return cond
}
