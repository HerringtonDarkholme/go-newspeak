package op

type ComparisonOp interface {
	Eq(LiteralOrColumn) BoolLikeOp
	Ne(LiteralOrColumn) BoolLikeOp
	Gte(LiteralOrColumn) BoolLikeOp
	Gt(LiteralOrColumn) BoolLikeOp
	Lte(LiteralOrColumn) BoolLikeOp
	Lt(LiteralOrColumn) BoolLikeOp
	NullSafeEq(LiteralOrColumn) BoolLikeOp
}
