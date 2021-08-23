package dao

import "gorm.io/gorm"

type scope interface {
	Query() func(db *gorm.DB) *gorm.DB
	Order() func(db *gorm.DB) *gorm.DB
	QueryAndOrder() func(db *gorm.DB) *gorm.DB
}

type updates interface {
	Query() func(db *gorm.DB) *gorm.DB
	Data() map[string]interface{}
}
