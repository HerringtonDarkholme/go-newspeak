package op

type RangeOp interface {
	Between(LiteralOrColumn, LiteralOrColumn) BoolLikeOp
	NotBetween(LiteralOrColumn, LiteralOrColumn) BoolLikeOp
}
