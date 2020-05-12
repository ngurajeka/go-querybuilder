package extractor

import (
	"fmt"

	querybuilder "github.com/ngurajeka/go-querybuilder"
	db "upper.io/db.v3"
)

// AsUpperDBCondition create upper.db condition usingfrom querybuilder
func AsUpperDBCondition(q querybuilder.QueryBuilder) db.Cond {
	cond := db.Cond{}
	for _, condition := range q.Conditions() {
		if condition.Operator() != "=" {
			field := fmt.Sprintf("%s %s", condition.Field(), condition.Operator())
			cond[field] = condition.Value()
		} else {
			cond[condition.Field()] = condition.Value()
		}
	}
	return cond
}
