package col

import (
	"github.com/HerringtonDarkholme/go-newspeak/pkg/op"
)

type StringColumn struct {
	name string
	op   op.OpInternal
}
