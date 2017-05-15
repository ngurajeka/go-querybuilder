package extractor

import (
	"fmt"

	querybuilder "github.com/ngurajeka/go-querybuilder"
	db "upper.io/db.v3"
)

func AsUpperDBCondition(q querybuilder.QueryBuilder) (conds []interface{}) {
	for _, cond := range q.Conditions() {
		field := fmt.Sprintf("%s %s", cond.Field(), cond.Operator())
		conds = append(conds, db.Cond{field: cond.Value()})
	}
	return
}
