package op

type InclusiveOp interface {
	In(...LiteralOrColumn) BoolLikeOp
	NotIn(...LiteralOrColumn) BoolLikeOp
}
