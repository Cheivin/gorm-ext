package criteria

import "gorm.io/gorm"

type update struct {
	update map[string]interface{}
	//Cause  *Cause
	*cause
}

type Update interface {
	Set(field string, val interface{}) Update
	SetIf(test bool, field string, val interface{}) Update
	SetExpr(field string, expr string, args ...interface{}) Update
	SetExprIf(test bool, field string, expr string, args ...interface{}) Update
	Cause
}

func NewUpdate(causes ...Cause) Update {
	c := new(cause)
	if len(causes) > 0 {
		c.And(causes...)
	}
	return &update{
		update: map[string]interface{}{},
		cause:  c,
	}
}

func (u *update) Set(field string, val interface{}) Update {
	u.update[field] = val
	return u
}

func (u *update) SetIf(test bool, field string, val interface{}) Update {
	if test {
		return u.Set(field, val)
	}
	return u
}

func (u *update) SetExprIf(test bool, field string, expr string, args ...interface{}) Update {
	if test {
		return u.SetExpr(field, expr, args...)
	}
	return u
}

func (u *update) SetExpr(field string, expr string, args ...interface{}) Update {
	u.update[field] = gorm.Expr(expr, args...)
	return u
}

func (u *update) Data() map[string]interface{} {
	return u.update
}

func (u *update) Query() func(db *gorm.DB) *gorm.DB {
	return u.query()
}

func Set(field string, val interface{}) Update {
	return NewUpdate().Set(field, val)
}

func SetIf(test bool, field string, val interface{}) Update {
	return NewUpdate().SetIf(test, field, val)
}

func SetExprIf(test bool, field string, expr string, args ...interface{}) Update {
	return NewUpdate().SetExprIf(test, field, expr, args...)
}

func SetExpr(field string, expr string, args ...interface{}) Update {
	return NewUpdate().SetExpr(field, expr, args...)
}
