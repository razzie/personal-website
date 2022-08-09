package modules

import (
	"github.com/razzie/beepboop"
	"github.com/razzie/gorzsony.com/assets"
	"github.com/razzie/gorzsony.com/pkg/layout"
)

// Hello returns the hello module
func Hello(loader *assets.AssetLoader) *layout.Module {
	return &layout.Module{
		Name:            "Hello",
		ContentTemplate: getContentTemplate("hello"),
		Handler: func(pr *beepboop.PageRequest) interface{} {
			return loader.Content().Bio
		},
	}
}
