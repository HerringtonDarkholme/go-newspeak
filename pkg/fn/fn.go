package fn

import (
	"fmt"
	"strings"

	ns "github.com/HerringtonDarkholme/go-newspeak/pkg/conditions"
)

type AggregateExpr interface {
}

type sqlFuncCall struct {
	fnName string
	alias  string
	args   []string
}

func (fn *sqlFuncCall) ToExprText() string {
	args := strings.Join(fn.args, ", ")
	return fmt.Sprintf("%s(%s)", fn.fnName, args)
}

func (fn *sqlFuncCall) ToSelectText() string {
	expr := fn.ToExprText()
	if fn.alias != "" {
		return fmt.Sprintf("%s as %s", expr, fn.alias)
	} else {
		return expr
	}
}

func Fn(name string, args ...interface{}) ns.SelectExpr {
	argStrings := make([]string, 0, len(args))
	for _, arg := range args {
		switch arg.(type) {
		case int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64,
			float32, float64, bool:
			argStrings = append(argStrings, fmt.Sprintf("%v", arg))
		default:
			if argStr, ok := arg.(ns.Expression); ok {
				argStrings = append(argStrings, argStr.ToExprText())
			}
		}
	}
	return &sqlFuncCall{
		fnName: name,
		args:   argStrings,
	}
}
