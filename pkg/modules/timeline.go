package modules

import (
	"github.com/razzie/beepboop"
	"github.com/razzie/gorzsony.com/pkg/layout"
	"github.com/razzie/gorzsony.com/pkg/projects"
)

// Timeline returns the timeline module
func Timeline() *layout.Module {
	timeline, err := projects.LoadTimeline()
	if err != nil {
		panic(err)
	}
	return &layout.Module{
		Name:            "Timeline",
		ContentTemplate: getContentTemplate("timeline"),
		Handler: func(pr *beepboop.PageRequest) interface{} {
			return timeline
		},
	}
}
