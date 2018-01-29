package grifts

import (
	"github.com/dnnrly/gostars/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
