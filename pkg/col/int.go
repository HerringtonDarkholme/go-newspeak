package col

import (
	"github.com/HerringtonDarkholme/go-newspeak/pkg/op"
)

type IntColumn struct {
	name string
	op   op.OpInternal
}
