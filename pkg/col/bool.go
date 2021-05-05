package col

import "github.com/HerringtonDarkholme/go-newspeak/pkg/op"

type BoolColumn struct {
	alias string
	name  string
}

func (c *BoolColumn) Eq(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *BoolColumn) Ne(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *BoolColumn) NullSafeEq(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *BoolColumn) Not() op.BoolLikeOp {
	return nil
}

func (c *BoolColumn) Is(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *BoolColumn) And(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *BoolColumn) Or(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *BoolColumn) Xor(other LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *BoolColumn) Between(start, end LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *BoolColumn) NotBetween(start, end LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *BoolColumn) In(others ...LiteralOrColumn) op.BoolLikeOp {
	return nil
}

func (c *BoolColumn) NotIn(others ...LiteralOrColumn) op.BoolLikeOp {
	return nil
}
func (c *BoolColumn) Case() *op.CaseStruct {
	return nil
}

func (c *BoolColumn) As(alias string) *BoolColumn {
	if c.alias != "" {
		panic("column alias cannot be reused!")
	}
	c.alias = alias
	return c
}

// Implement WhereCondition
func (c *BoolColumn) ToConditionText() string {
	return c.name
}

// Implement WhereCondition
func (c BoolColumn) ToExprText() string {
	return ""
}

var _ op.BoolLikeOp = (*BoolColumn)(nil)
