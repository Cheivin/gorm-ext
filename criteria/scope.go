package criteria

import "gorm.io/gorm"

func (c *cause) query() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		query, args, _ := c.build()
		if query == "" {
			return db
		}
		return db.Where(query, args...)
	}
}

func (c *cause) Query() func(db *gorm.DB) *gorm.DB {
	return c.query()
}

func (c *cause) Order() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		_, _, order := c.build()
		if order != "" {
			db.Order(order)
		}
		return db
	}
}

func (c *cause) QueryAndOrder() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		query, args, order := c.build()
		if order != "" {
			return db.Where(query, args...).Order(order)
		} else {
			return db.Where(query, args...)
		}
	}
}
