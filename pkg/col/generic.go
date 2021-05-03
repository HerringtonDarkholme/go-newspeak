package col

import (
	ns "github.com/HerringtonDarkholme/go-newspeak/pkg"
	"github.com/HerringtonDarkholme/go-newspeak/pkg/op"
)

type GenericColumn struct {
	name string
	op   op.OpInternal
}

func (c *GenericColumn) Eq(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) Ne(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) Gte(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) Gt(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) Lte(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) Lt(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) NullSafeEq(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) Not() ns.Expression {
	return nil
}

func (c *GenericColumn) Is(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) And(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) Or(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) Xor(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) Between(start, end LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) NotBetween(start, end LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) In(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) NotIn(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) Like(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) NotLike(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) Regexp(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) NotRegexp(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) Mul(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) Add(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) Sub(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) Neg() ns.Expression {
	return nil
}

func (c *GenericColumn) Div(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) Mod(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) FloorDiv(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) RightShift(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) LeftShift(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) BitXor(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) BitOr(other LiteralOrColumn) ns.Expression {
	return nil
}

func (c *GenericColumn) BitFlip() ns.Expression {
	return nil
}

func (c *GenericColumn) Case() *op.CaseStruct {
	return nil
}
