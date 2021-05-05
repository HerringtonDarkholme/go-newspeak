package op

type EqualOp interface {
	Eq(LiteralOrColumn) BoolLikeOp
	Ne(LiteralOrColumn) BoolLikeOp
	NullSafeEq(LiteralOrColumn) BoolLikeOp
}
