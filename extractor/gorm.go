package extractor

import (
	"github.com/jinzhu/gorm"
	querybuilder "github.com/ngurajeka/go-querybuilder/v2"
)

// GormQueryset appending gorm queryset into gorm database from querybuilder
func GormQueryset(db *gorm.DB, qb querybuilder.QueryBuilder) *gorm.DB {
	if qb.HasConditions() {
		stmt, args := qb.PrepareStatement()
		db = db.Where(stmt, args...)
	}
	if qb.HasOrders() {
		db = db.Order(qb.StringifyOrder())
	}
	if qb.Limit() > 0 {
		db = db.Limit(qb.Limit())
	}
	if qb.Offset() > 0 {
		db = db.Offset(qb.Offset())
	}
	return db
}
