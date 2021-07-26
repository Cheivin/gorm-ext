package dao

import (
	"github.com/cheivin/gorm-ext/criteria"
	"gorm.io/gorm"
)

func FindByID(db *gorm.DB, ID uint, target interface{}) (ok bool, err error) {
	return FindOne(db, criteria.Eq("id", ID), target)
}

func FindByIDs(db *gorm.DB, IDs []uint, target interface{}) error {
	if len(IDs) == 0 {
		return nil
	}
	return db.Find(target, IDs).Error
}

func FindOne(db *gorm.DB, cause criteria.Cause, target interface{}) (ok bool, err error) {
	if cause == nil {
		cause = criteria.New()
	}
	db = db.Scopes(cause.(scope).Query()).Limit(1).Find(target)
	return db.RowsAffected > 0, db.Error
}

func FindAll(db *gorm.DB, cause criteria.Cause, target interface{}) error {
	if cause == nil {
		cause = criteria.New()
	}
	return db.Scopes(cause.(scope).QueryAndOrder()).Find(target).Error
}

func Page(db *gorm.DB, page, size int, cause criteria.Cause, target interface{}) (total int64, err error) {
	if cause == nil {
		cause = criteria.New()
	}
	err = db.Scopes(cause.(scope).Query()).
		Count(&total).
		Scopes(cause.(scope).QueryAndOrder()).
		Offset(page * size).Limit(size).
		Find(target).
		Error
	return
}

func Delete(db *gorm.DB, model interface{}, cause criteria.Cause) (int64, error) {
	if cause == nil {
		db = db.Delete(model)
	} else {
		db = db.Scopes(cause.(scope).Query()).Delete(model)
	}
	return db.RowsAffected, db.Error
}

func Update(db *gorm.DB, update criteria.Update) (int64, error) {
	if update == nil {
		update = criteria.NewUpdate()
	}
	db = db.Scopes(update.(updates).Query()).Updates(update.(updates).Data())
	return db.RowsAffected, db.Error
}

func GetFieldMap(db *gorm.DB, field string, cause criteria.Cause, target interface{}) (err error) {
	var results []map[string]interface{}
	if cause != nil {
		err = db.Select("id", field).Scopes(cause.(scope).Query()).Find(&results).Error
	} else {
		err = db.Select("id", field).Find(&results).Error
	}
	if err != nil {
		return
	}
	accept := target.(map[uint]interface{})
	for _, result := range results {
		id := result["id"].(uint)
		val := result[field]
		accept[id] = val
	}
	return
}
