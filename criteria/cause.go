package criteria

import (
	"gorm.io/gorm"
	"reflect"
	"strings"
)

type (
	cause struct {
		fragments []fragment
		args      []interface{}
		order     []string
	}
	fragment struct {
		and   bool
		query string
	}

	Cause interface {
		ExprIf(test bool, field string, val ...interface{}) Cause
		Expr(field string, val ...interface{}) Cause
		LikeIf(test bool, field string, val interface{}) Cause
		Like(field string, val interface{}) Cause
		EqIf(test bool, field string, val interface{}) Cause
		Eq(field string, val interface{}) Cause
		GtIf(test bool, field string, val interface{}) Cause
		Gt(field string, val interface{}) Cause
		GteIf(test bool, field string, val interface{}) Cause
		Gte(field string, val interface{}) Cause
		LtIf(test bool, field string, val interface{}) Cause
		Lt(field string, val interface{}) Cause
		LteIf(test bool, field string, val interface{}) Cause
		Lte(field string, val interface{}) Cause
		In(field string, val ...interface{}) Cause
		InIf(test bool, field string, val ...interface{}) Cause
		InSql(field string, sql string, val ...interface{}) Cause
		InSqlIf(test bool, field string, sql string, val ...interface{}) Cause
		And(causes ...Cause) Cause
		Or(cause Cause) Cause
		Asc(fields ...string) Cause
		Desc(fields ...string) Cause

		build() (string, []interface{}, string)
		query() func(db *gorm.DB) *gorm.DB
	}
)

func New() Cause {
	return new(cause)
}

func (c *cause) ExprIf(test bool, field string, val ...interface{}) Cause {
	if test {
		return c.Expr(field, val...)
	}
	return c
}

func (c *cause) Expr(field string, val ...interface{}) Cause {
	c.fragments = append(c.fragments, fragment{and: true, query: field})
	c.args = append(c.args, val...)
	return c
}

func (c *cause) LikeIf(test bool, field string, val interface{}) Cause {
	if test {
		return c.Like(field, val)
	}
	return c
}

func (c *cause) Like(field string, val interface{}) Cause {
	return c.Expr(field+" like ?", val)
}

func (c *cause) EqIf(test bool, field string, val interface{}) Cause {
	if test {
		return c.Eq(field, val)
	}
	return c
}

func (c *cause) Eq(field string, val interface{}) Cause {
	return c.Expr(field+" = ?", val)
}

func (c *cause) GtIf(test bool, field string, val interface{}) Cause {
	if test {
		return c.Gt(field, val)
	}
	return c
}

func (c *cause) Gt(field string, val interface{}) Cause {
	return c.Expr(field+" > ?", val)
}

func (c *cause) GteIf(test bool, field string, val interface{}) Cause {
	if test {
		return c.Gte(field, val)
	}
	return c
}

func (c *cause) Gte(field string, val interface{}) Cause {
	return c.Expr(field+" >= ?", val)
}

func (c *cause) LtIf(test bool, field string, val interface{}) Cause {
	if test {
		return c.Lt(field, val)
	}
	return c
}

func (c *cause) Lt(field string, val interface{}) Cause {
	return c.Expr(field+" < ?", val)
}

func (c *cause) LteIf(test bool, field string, val interface{}) Cause {
	if test {
		return c.Lte(field, val)
	}
	return c
}

func (c *cause) Lte(field string, val interface{}) Cause {
	return c.Expr(field+" <= ?", val)
}

func (c *cause) In(field string, val ...interface{}) Cause {
	if val == nil {
		return c
	}
	switch reflect.TypeOf(val).Kind() {
	case reflect.Array, reflect.Slice:
		s := reflect.ValueOf(val)
		if s.Len() == 0 {
			return c
		} else if s.Len() == 1 {
			v := s.Index(0)
			switch v.Elem().Kind() {
			case reflect.Array, reflect.Slice:
				if v.Elem().Len() == 0 {
					return c
				} else if v.Elem().Len() == 1 {
					return c.Eq(field, v.Elem().Index(0).Interface())
				}
				return c.Expr(field+" in ?", v.Interface())
			}
			return c.Eq(field, v.Interface())
		}
	}
	return c.Expr(field+" in ?", val)
}

func (c *cause) InIf(test bool, field string, val ...interface{}) Cause {
	if test {
		return c.In(field, val...)
	}
	return c
}

func (c *cause) InSql(field string, sql string, val ...interface{}) Cause {
	return c.Expr(field+" in ("+sql+")", val...)
}

func (c *cause) InSqlIf(test bool, field string, sql string, val ...interface{}) Cause {
	if test {
		return c.InSql(field, sql, val...)
	}
	return c
}

func (c *cause) And(causes ...Cause) Cause {
	if len(causes) == 0 {
		return c
	}
	for _, cause := range causes {
		fragmentCause, args, _ := cause.build()
		if fragmentCause == "" {
			return c
		}
		c.fragments = append(c.fragments, fragment{and: true, query: "( " + fragmentCause + " )"})
		c.args = append(c.args, args...)
	}
	return c
}

func (c *cause) Or(cause Cause) Cause {
	fragmentCause, args, _ := cause.build()
	if fragmentCause == "" {
		return c
	}
	c.fragments = append(c.fragments, fragment{and: false, query: "( " + fragmentCause + " )"})
	c.args = append(c.args, args...)
	return c
}

func (c *cause) Asc(fields ...string) Cause {
	c.order = append(c.order, fields...)
	return c
}

func (c *cause) Desc(fields ...string) Cause {
	for _, field := range fields {
		c.order = append(c.order, field+" desc")
	}
	return c
}

func (c *cause) build() (string, []interface{}, string) {
	orderStr := strings.Join(c.order, ", ")
	elems := c.fragments
	andSep := " and "
	orSep := " or "
	switch len(elems) {
	case 0:
		return "", c.args, orderStr
	case 1:
		return elems[0].query, c.args, orderStr
	}
	n := 0
	for i := 0; i < len(elems); i++ {
		n += len(elems[i].query)
		if elems[i].and {
			n += len(andSep)
		} else {
			n += len(orSep)
		}
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(elems[0].query)
	for _, s := range elems[1:] {
		if s.and {
			b.WriteString(andSep)
			b.WriteString(s.query)
		} else {
			b.WriteString(orSep)
			b.WriteString(s.query)
		}
	}
	return b.String(), c.args, orderStr
}

func ExprIf(test bool, field string, val ...interface{}) Cause {
	return new(cause).ExprIf(test, field, val...)
}

func Expr(field string, val ...interface{}) Cause {
	return new(cause).Expr(field, val...)
}

func LikeIf(test bool, field string, val interface{}) Cause {
	return new(cause).LikeIf(test, field, val)
}

func Like(field string, val interface{}) Cause {
	return new(cause).Like(field, val)
}

func EqIf(test bool, field string, val interface{}) Cause {
	return new(cause).EqIf(test, field, val)
}

func Eq(field string, val interface{}) Cause {
	return new(cause).Eq(field, val)
}

func GtIf(test bool, field string, val interface{}) Cause {
	return new(cause).GtIf(test, field, val)
}

func Gt(field string, val interface{}) Cause {
	return new(cause).Gt(field, val)
}

func GteIf(test bool, field string, val interface{}) Cause {
	return new(cause).GteIf(test, field, val)
}

func Gte(field string, val interface{}) Cause {
	return new(cause).Gte(field, val)
}

func LtIf(test bool, field string, val interface{}) Cause {
	return new(cause).LtIf(test, field, val)
}

func Lt(field string, val interface{}) Cause {
	return new(cause).Lt(field, val)
}

func LteIf(test bool, field string, val interface{}) Cause {
	return new(cause).LteIf(test, field, val)
}

func Lte(field string, val interface{}) Cause {
	return new(cause).Lte(field, val)
}

func In(field string, val ...interface{}) Cause {
	return new(cause).In(field, val...)
}

func InIf(test bool, field string, val ...interface{}) Cause {
	return new(cause).InIf(test, field, val...)
}

func And(causes ...Cause) Cause {
	return New().And(causes...)
}

func Or(causes ...Cause) Cause {
	switch len(causes) {
	case 0:
		return new(cause)
	case 1:
		return causes[0]
	default:
		cause := new(cause)
		for i := range causes {
			cause.Or(causes[i])
		}
		return cause
	}
}
