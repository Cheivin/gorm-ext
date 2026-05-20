package wrapper

func Set(field string, val any) *Update {
	return U().Set(field, val)
}

func SetIf(test bool, field string, val any) *Update {
	return U().SetIf(test, field, val)
}

func SetExprIf(test bool, field string, Expr string, args ...any) *Update {
	return U().SetExprIf(test, field, Expr, args...)
}

func SetExpr(field string, Expr string, args ...any) *Update {
	return U().SetExpr(field, Expr, args...)
}
