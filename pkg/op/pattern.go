package op

type PatternOp interface {
	Like(LiteralOrColumn) BoolLikeOp
	NotLike(LiteralOrColumn) BoolLikeOp
	Regexp(LiteralOrColumn) BoolLikeOp
	NotRegexp(LiteralOrColumn) BoolLikeOp
}
