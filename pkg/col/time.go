package col

import "github.com/HerringtonDarkholme/go-newspeak/pkg/op"

type TimeColumn struct {
	name  string
	alias string
}

func (c *TimeColumn) Eq(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *TimeColumn) Ne(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *TimeColumn) Gte(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *TimeColumn) Gt(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *TimeColumn) Lte(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *TimeColumn) Lt(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *TimeColumn) NullSafeEq(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *TimeColumn) Not() op.BoolLikeOp {
	return nil
}

func (c *TimeColumn) Is(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *TimeColumn) And(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *TimeColumn) Or(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *TimeColumn) Xor(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *TimeColumn) Between(start, end LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *TimeColumn) NotBetween(start, end LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *TimeColumn) In(others ...LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *TimeColumn) NotIn(others ...LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *TimeColumn) Case() *op.CaseStruct {
	return nil
}

func (c *TimeColumn) As(alias string) *TimeColumn {
	if c.alias != "" {
		panic("column alias cannot be reused!")
	}
	c.alias = alias
	return c
}

// Implement WhereCondition
func (c TimeColumn) ToExprText() string {
	return ""
}

var _ op.DateLikeOp = (*TimeColumn)(nil)
