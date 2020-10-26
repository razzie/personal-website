package modules

import (
	"github.com/razzie/gorzsony.com/pkg/layout"
)

// Hello returns the hello module
func Hello() *layout.Module {
	return &layout.Module{
		Name:            "Hello",
		ContentTemplate: getContentTemplate("hello"),
	}
}
