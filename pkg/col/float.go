package col

import "github.com/HerringtonDarkholme/go-newspeak/pkg/op"

type FloatColumn struct {
	name  string
	alias string
}

func (c *FloatColumn) Eq(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *FloatColumn) Ne(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *FloatColumn) Gte(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *FloatColumn) Gt(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *FloatColumn) Lte(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *FloatColumn) Lt(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *FloatColumn) NullSafeEq(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *FloatColumn) Not() op.BoolLikeOp {
	return nil
}

func (c *FloatColumn) Is(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *FloatColumn) And(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *FloatColumn) Or(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *FloatColumn) Xor(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *FloatColumn) Between(start, end LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *FloatColumn) NotBetween(start, end LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *FloatColumn) In(others ...LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *FloatColumn) NotIn(others ...LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *FloatColumn) Mul(other LiteralOrColumn) op.NumberLikeOp {
	return nil
}

func (c *FloatColumn) Add(other LiteralOrColumn) op.NumberLikeOp {
	return nil
}

func (c *FloatColumn) Sub(other LiteralOrColumn) op.NumberLikeOp {
	return nil
}

func (c *FloatColumn) Neg() op.NumberLikeOp {
	return nil
}

func (c *FloatColumn) Div(other LiteralOrColumn) op.NumberLikeOp {
	return nil
}

func (c *FloatColumn) Mod(other LiteralOrColumn) op.NumberLikeOp {
	return nil
}

func (c *FloatColumn) FloorDiv(other LiteralOrColumn) op.NumberLikeOp {
	return nil
}

func (c *FloatColumn) Case() *op.CaseStruct {
	return nil
}

func (c *FloatColumn) As(alias string) *FloatColumn {
	if c.alias != "" {
		panic("column alias cannot be reused!")
	}
	c.alias = alias
	return c
}

// Implement WhereCondition
func (c FloatColumn) ToExprText() string {
	return c.name
}

var _ op.NumberLikeOp = (*FloatColumn)(nil)
