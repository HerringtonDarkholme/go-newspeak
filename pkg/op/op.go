package op

import (
	ns "github.com/HerringtonDarkholme/go-newspeak/pkg"
)

type OpType uint8

const (
	Binary OpType = iota + 1
	Prefix
)

type Operator struct {
	OpString string
	OpType   OpType
}

type LiteralOrColumn interface {
}

type GenericOp interface {
	EqualOp
	ComparativeOp
	RangeOp
	InclusiveOp
	ArithmeticOp
	BitwiseOp
	LogicalOp
	PatternOp
	ns.WhereCondition
	ns.Expression
}

type CommonOp interface {
	EqualOp
	RangeOp
	InclusiveOp
	ns.Expression
}

type IntLikeOp interface {
	NumberLikeOp
	BitwiseOp
}

type NumberLikeOp interface {
	CommonOp
	ComparativeOp
	ArithmeticOp
}

type BoolLikeOp interface {
	CommonOp
	LogicalOp
	ns.WhereCondition
}

type StringLikeOp interface {
	CommonOp
	ComparativeOp
	PatternOp
}

type DateLikeOp interface {
	CommonOp
	ComparativeOp
}
