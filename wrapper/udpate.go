package wrapper

import (
	"gorm.io/gorm"
)

type (
	Update struct {
		query  *Query
		update map[string]any
	}
)

func U(querys ...*Query) *Update {
	return &Update{
		query:  And(querys...),
		update: map[string]any{},
	}
}

func (u *Update) Query() *Query {
	return u.query
}

func (u *Update) Data() map[string]any {
	return u.update
}

func (u *Update) Set(field string, val any) *Update {
	u.update[field] = val
	return u
}

func (u *Update) SetIf(test bool, field string, val any) *Update {
	if test {
		return u.Set(field, val)
	}
	return u
}

func (u *Update) SetExpr(field string, expr string, args ...any) *Update {
	u.update[field] = gorm.Expr(expr, args...)
	return u
}

func (u *Update) SetExprIf(test bool, field string, Expr string, args ...any) *Update {
	if test {
		return u.SetExpr(field, Expr, args...)
	}
	return u
}

// 查询条件部分

func (u *Update) ExprIf(test bool, field string, val ...any) *Update {
	if test {
		return u.Expr(field, val...)
	}
	return u
}

func (u *Update) Expr(field string, val ...any) *Update {
	u.query.Expr(field, val...)
	return u
}

func (u *Update) LikeIf(test bool, field string, val any) *Update {
	if test {
		return u.Like(field, val)
	}
	return u
}

func (u *Update) Like(field string, val any) *Update {
	u.query.Like(field, val)
	return u
}

func (u *Update) EqIf(test bool, field string, val any) *Update {
	if test {
		return u.Eq(field, val)
	}
	return u
}

func (u *Update) Eq(field string, val any) *Update {
	u.query.Eq(field, val)
	return u
}

func (u *Update) GtIf(test bool, field string, val any) *Update {
	if test {
		return u.Gt(field, val)
	}
	return u
}

func (u *Update) Gt(field string, val any) *Update {
	u.query.Gt(field, val)
	return u
}

func (u *Update) GteIf(test bool, field string, val any) *Update {
	if test {
		return u.Gte(field, val)
	}
	return u
}

func (u *Update) Gte(field string, val any) *Update {
	u.query.Gte(field, val)
	return u
}

func (u *Update) LtIf(test bool, field string, val any) *Update {
	if test {
		return u.Lt(field, val)
	}
	return u
}

func (u *Update) Lt(field string, val any) *Update {
	u.query.Lt(field, val)
	return u
}

func (u *Update) LteIf(test bool, field string, val any) *Update {
	if test {
		return u.Lte(field, val)
	}
	return u
}

func (u *Update) Lte(field string, val any) *Update {
	u.query.Lte(field, val)
	return u
}

func (u *Update) In(field string, val ...any) *Update {
	u.query.In(field, val...)
	return u
}

func (u *Update) InIf(test bool, field string, val ...any) *Update {
	if test {
		return u.In(field, val...)
	}
	return u
}

func (u *Update) InSql(field string, sql string, val ...any) *Update {
	u.query.InSql(field, sql, val...)
	return u
}

func (u *Update) InSqlIf(test bool, field string, sql string, val ...any) *Update {
	if test {
		return u.InSql(field, sql, val...)
	}
	return u
}

func (u *Update) And(causes ...*Query) *Update {
	u.query.And(causes...)
	return u
}

func (u *Update) Or(cause *Query) *Update {
	u.query.Or(cause)
	return u
}
