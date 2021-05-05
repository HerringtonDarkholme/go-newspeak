package newspeak

type WhereCondition interface {
	ToConditionText() string
}

type TableSpecifier interface {
	ToTableText() string
}

type Expression interface {
}

type OrderExpression interface {
}

type WindowSpec interface {
}
