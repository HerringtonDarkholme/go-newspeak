package op

type BitwiseOp interface {
	RightShift(LiteralOrColumn) IntLikeOp
	LeftShift(LiteralOrColumn) IntLikeOp
	BitXor(LiteralOrColumn) IntLikeOp
	BitOr(LiteralOrColumn) IntLikeOp
	BitAnd(LiteralOrColumn) IntLikeOp
	BitFlip() IntLikeOp
}
