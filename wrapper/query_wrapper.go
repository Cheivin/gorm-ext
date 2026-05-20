package wrapper

func ExprIf(test bool, field string, val ...any) *Query {
	return Q().ExprIf(test, field, val...)
}

func Expr(field string, val ...any) *Query {
	return Q().Expr(field, val...)
}

func LikeIf(test bool, field string, val any) *Query {
	return Q().LikeIf(test, field, val)
}

func Like(field string, val any) *Query {
	return Q().Like(field, val)
}

func EqIf(test bool, field string, val any) *Query {
	return Q().EqIf(test, field, val)
}

func Eq(field string, val any) *Query {
	return Q().Eq(field, val)
}

func GtIf(test bool, field string, val any) *Query {
	return Q().GtIf(test, field, val)
}

func Gt(field string, val any) *Query {
	return Q().Gt(field, val)
}

func GteIf(test bool, field string, val any) *Query {
	return Q().GteIf(test, field, val)
}

func Gte(field string, val any) *Query {
	return Q().Gte(field, val)
}

func LtIf(test bool, field string, val any) *Query {
	return Q().LtIf(test, field, val)
}

func Lt(field string, val any) *Query {
	return Q().Lt(field, val)
}

func LteIf(test bool, field string, val any) *Query {
	return Q().LteIf(test, field, val)
}

func Lte(field string, val any) *Query {
	return Q().Lte(field, val)
}

func In(field string, val ...any) *Query {
	return Q().In(field, val...)
}

func InIf(test bool, field string, val ...any) *Query {
	return Q().InIf(test, field, val...)
}

func InSql(field string, sql string, val ...any) *Query {
	return Q().InSql(field, sql, val...)
}

func InSqlIf(test bool, field string, sql string, val ...any) *Query {
	return Q().InSqlIf(test, field, sql, val...)
}

func And(causes ...*Query) (q *Query) {
	q = Q()
	if len(causes) > 0 {
		q.And(causes...)
	}
	return q
}

func Or(causes ...*Query) (q *Query) {
	q = Q()
	if len(causes) > 0 {
		for i := range causes {
			q = q.Or(causes[i])
		}
	}
	return q
}
