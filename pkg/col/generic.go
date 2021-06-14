package col

import (
	"github.com/HerringtonDarkholme/go-newspeak/pkg/conditions"
	"github.com/HerringtonDarkholme/go-newspeak/pkg/op"
)

type GenericColumn struct {
	name  string
	alias string
	op    op.Operator
	arg   interface{}
}

func (c *GenericColumn) Eq(other LiteralOrColumn) op.BoolLikeOp {
	return &GenericColumn{
		name:  c.name,
		alias: c.alias,
		op: op.Operator{
			OpString: "=",
			OpType:   op.Binary,
		},
		arg: other,
	}
}

func (c *GenericColumn) Ne(other LiteralOrColumn) op.BoolLikeOp {
	return &GenericColumn{}
}

func (c *GenericColumn) Gte(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *GenericColumn) Gt(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *GenericColumn) Lte(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *GenericColumn) Lt(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *GenericColumn) NullSafeEq(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *GenericColumn) Not() op.BoolLikeOp {
	return nil
}

func (c *GenericColumn) Is(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *GenericColumn) And(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *GenericColumn) Or(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *GenericColumn) Xor(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *GenericColumn) Between(start, end LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *GenericColumn) NotBetween(start, end LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *GenericColumn) In(others ...LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *GenericColumn) NotIn(others ...LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *GenericColumn) Like(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *GenericColumn) NotLike(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *GenericColumn) Regexp(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *GenericColumn) NotRegexp(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *GenericColumn) Mul(other LiteralOrColumn) op.NumberLikeOp {
	return nil
}

func (c *GenericColumn) Add(other LiteralOrColumn) op.NumberLikeOp {
	return nil
}

func (c *GenericColumn) Sub(other LiteralOrColumn) op.NumberLikeOp {
	return nil
}

func (c *GenericColumn) Neg() op.NumberLikeOp {
	return nil
}

func (c *GenericColumn) Div(other LiteralOrColumn) op.NumberLikeOp {
	return nil
}

func (c *GenericColumn) Mod(other LiteralOrColumn) op.NumberLikeOp {
	return nil
}

func (c *GenericColumn) FloorDiv(other LiteralOrColumn) op.NumberLikeOp {
	return nil
}

func (c *GenericColumn) RightShift(other LiteralOrColumn) op.IntLikeOp {
	return nil
}

func (c *GenericColumn) LeftShift(other LiteralOrColumn) op.IntLikeOp {
	return nil
}

func (c *GenericColumn) BitXor(other LiteralOrColumn) op.IntLikeOp {
	return nil
}

func (c *GenericColumn) BitAnd(other LiteralOrColumn) op.IntLikeOp {
	return nil
}

func (c *GenericColumn) BitOr(other LiteralOrColumn) op.IntLikeOp {
	return nil
}

func (c *GenericColumn) BitFlip() op.IntLikeOp {
	return nil
}

func (c *GenericColumn) Case() *op.CaseStruct {
	return nil
}

func (c *GenericColumn) As(alias string) *GenericColumn {
	if c.alias != "" {
		panic("column alias cannot be reused!")
	}
	c.alias = alias
	return c
}

// Implement WhereCondition
func (c *GenericColumn) ToCondition() conditions.Condition {
	return conditions.Condition{
		Query: c.name + c.op.OpString + "?",
		Args:  []interface{}{c.arg},
	}
}

// Implement WhereCondition
func (c GenericColumn) ToExprText() string {
	return c.name
}

var _ op.GenericOp = (*GenericColumn)(nil)
