package memberships

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

func NewReporitory(db *gorm.DB) *repository {
	return &repository{db: db}
}
