package col

import (
	"github.com/HerringtonDarkholme/go-newspeak/pkg/op"
)

type FloatColumn struct {
	name string
	op   op.OpInternal
}
