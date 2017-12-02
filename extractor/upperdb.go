package extractor

import (
	"fmt"

	querybuilder "github.com/ngurajeka/go-querybuilder"
	db "upper.io/db.v3"
)

func AsUpperDBCondition(q querybuilder.QueryBuilder) db.Cond {
	cond := db.Cond{}
	for _, condition := range q.Conditions() {
		field := fmt.Sprintf("%s %s", condition.Field(), condition.Operator())
		cond[field] = condition.Value()
	}
	return cond
}
