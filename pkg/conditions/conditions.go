package conditions

type Condition struct {
	Query string
	Args  []interface{}
}

// where_condition is an expression that evaluates to true for each row to be selected
// cannot reference aggregate
type WhereCondition interface {
	ToCondition() Condition
}

type TableSpecifier interface {
	// table name
	// subquery as table_alias
	ToTableText() string
}

type Expression interface {
	ToExprText() string
}

type SelectExpr interface {
	// expr | expr as alias
	ToSelectText() string
	Expression
}

type OrderExpression interface {
	// col_name | expr | position
}

type WindowSpec interface {
}
