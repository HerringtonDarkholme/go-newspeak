package op

import (
	ns "github.com/HerringtonDarkholme/go-newspeak/pkg"
)

type OpInternal struct {
}

type CaseStruct struct {
}

func (c *CaseStruct) When(expr, result ns.Expression) *CaseStruct {
	return c
}

func (c *CaseStruct) Else(expr ns.Expression) ns.Expression {
	return nil
}
