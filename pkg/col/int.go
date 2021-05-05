package col

import (
	"github.com/HerringtonDarkholme/go-newspeak/pkg/op"
)

type IntColumn struct {
	name  string
	alias string
}

func (c *IntColumn) Eq(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *IntColumn) Ne(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *IntColumn) Gte(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *IntColumn) Gt(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *IntColumn) Lte(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *IntColumn) Lt(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *IntColumn) NullSafeEq(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *IntColumn) Not() op.BoolLikeOp {
	return nil
}

func (c *IntColumn) Is(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *IntColumn) And(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *IntColumn) Or(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *IntColumn) Xor(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *IntColumn) Between(start, end LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *IntColumn) NotBetween(start, end LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *IntColumn) In(others ...LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *IntColumn) NotIn(others ...LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *IntColumn) Like(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *IntColumn) NotLike(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *IntColumn) Regexp(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *IntColumn) NotRegexp(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *IntColumn) Mul(other LiteralOrColumn) op.NumberLikeOp {
	return nil
}

func (c *IntColumn) Add(other LiteralOrColumn) op.NumberLikeOp {
	return nil
}

func (c *IntColumn) Sub(other LiteralOrColumn) op.NumberLikeOp {
	return nil
}

func (c *IntColumn) Neg() op.NumberLikeOp {
	return nil
}

func (c *IntColumn) Div(other LiteralOrColumn) op.NumberLikeOp {
	return nil
}

func (c *IntColumn) Mod(other LiteralOrColumn) op.NumberLikeOp {
	return nil
}

func (c *IntColumn) FloorDiv(other LiteralOrColumn) op.NumberLikeOp {
	return nil
}

func (c *IntColumn) RightShift(other LiteralOrColumn) op.IntLikeOp {
	return nil
}

func (c *IntColumn) LeftShift(other LiteralOrColumn) op.IntLikeOp {
	return nil
}

func (c *IntColumn) BitXor(other LiteralOrColumn) op.IntLikeOp {
	return nil
}

func (c *IntColumn) BitAnd(other LiteralOrColumn) op.IntLikeOp {
	return nil
}

func (c *IntColumn) BitOr(other LiteralOrColumn) op.IntLikeOp {
	return nil
}

func (c *IntColumn) BitFlip() op.IntLikeOp {
	return nil
}

func (c *IntColumn) Case() *op.CaseStruct {
	return nil
}

func (c *IntColumn) As(alias string) *IntColumn {
	if c.alias != "" {
		panic("column alias cannot be reused!")
	}
	c.alias = alias
	return c
}

// Implement WhereCondition
func (c IntColumn) ToExprText() string {
	return ""
}

var _ op.IntLikeOp = (*IntColumn)(nil)
