package col

import (
	"github.com/HerringtonDarkholme/go-newspeak/pkg/op"
)

type BoolColumn struct {
	name string
	op   op.OpInternal
}
