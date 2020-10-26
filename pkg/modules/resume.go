package modules

import (
	"github.com/razzie/gorzsony.com/pkg/layout"
)

// Resume returns the resume module
func Resume() *layout.Module {
	return &layout.Module{
		Name:            "Resume",
		ContentTemplate: getContentTemplate("resume"),
	}
}
