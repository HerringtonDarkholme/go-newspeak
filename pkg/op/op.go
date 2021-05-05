package op

type LiteralOrColumn interface {
}

type GenericOp interface {
	ComparisonOp
	RangeOp
	InclusionOp
	ArithmeticOp
	BitwiseOp
	LogicalOp
	PatternOp
}

type CommonOp interface {
	ComparisonOp
	RangeOp
	InclusionOp
}

type IntLikeOp interface {
	CommonOp
	ArithmeticOp
}

type NumberLikeOp interface {
	IntLikeOp
	BitwiseOp
}

type BoolLikeOp interface {
	CommonOp
	LogicalOp
}

type StringLikeOp interface {
	CommonOp
	PatternOp
}
