package modules

import (
	"github.com/razzie/beepboop"
	"github.com/razzie/gorzsony.com/assets"
	"github.com/razzie/gorzsony.com/pkg/layout"
)

// Resume returns the resume module
func Resume(loader *assets.AssetLoader) *layout.Module {
	return &layout.Module{
		Name:            "Resume",
		ContentTemplate: getContentTemplate("resume"),
		Handler: func(pr *beepboop.PageRequest) interface{} {
			return loader.Content().Resume
		},
	}
}
