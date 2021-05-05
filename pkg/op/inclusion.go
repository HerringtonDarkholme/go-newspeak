package op

type InclusionOp interface {
	In(...LiteralOrColumn) BoolLikeOp
	NotIn(...LiteralOrColumn) BoolLikeOp
}
