package op

type LogicalOp interface {
	Not() BoolLikeOp
	Is(LiteralOrColumn) BoolLikeOp
	And(LiteralOrColumn) BoolLikeOp
	Or(LiteralOrColumn) BoolLikeOp
	Xor(LiteralOrColumn) BoolLikeOp
}
