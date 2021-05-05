package op

type ArithmeticOp interface {
	Add(LiteralOrColumn) NumberLikeOp
	Sub(LiteralOrColumn) NumberLikeOp
	Mul(LiteralOrColumn) NumberLikeOp
	Div(LiteralOrColumn) NumberLikeOp
	Mod(LiteralOrColumn) NumberLikeOp
	FloorDiv(LiteralOrColumn) NumberLikeOp
	Neg() NumberLikeOp
}
