package col

import (
	"github.com/HerringtonDarkholme/go-newspeak/pkg/op"
)

type StringColumn struct {
	name  string
	alias string
}

func (c *StringColumn) Eq(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *StringColumn) Ne(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *StringColumn) Gte(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *StringColumn) Gt(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *StringColumn) Lte(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *StringColumn) Lt(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *StringColumn) NullSafeEq(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *StringColumn) Not() op.BoolLikeOp {
	return nil
}

func (c *StringColumn) Is(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *StringColumn) Between(start, end LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *StringColumn) NotBetween(start, end LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *StringColumn) In(others ...LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *StringColumn) NotIn(others ...LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *StringColumn) Like(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *StringColumn) NotLike(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *StringColumn) Regexp(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *StringColumn) NotRegexp(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *StringColumn) Case() *op.CaseStruct {
	return nil
}

func (c *StringColumn) As(alias string) *StringColumn {
	if c.alias != "" {
		panic("column alias cannot be reused!")
	}
	c.alias = alias
	return c
}

// Implement Expr
func (c StringColumn) ToExprText() string {
	return ""
}

var _ op.StringLikeOp = (*StringColumn)(nil)
