package op

type ComparativeOp interface {
	Gte(LiteralOrColumn) BoolLikeOp
	Gt(LiteralOrColumn) BoolLikeOp
	Lte(LiteralOrColumn) BoolLikeOp
	Lt(LiteralOrColumn) BoolLikeOp
}
