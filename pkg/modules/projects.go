package modules

import (
	"math/rand"
	"strings"
	"time"

	"github.com/razzie/beepboop"
	"github.com/razzie/gorzsony.com/pkg/layout"
	"github.com/razzie/gorzsony.com/pkg/projects"
)

type projectsView struct {
	Tag      string
	Projects []projects.Project
}

func shuffleProjects(projects []projects.Project) []projects.Project {
	clone := append(projects[:0:0], projects...)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(clone), func(i, j int) { clone[i], clone[j] = clone[j], clone[i] })
	return clone
}

func filterProjects(projects []projects.Project, tag string) (results []projects.Project) {
	tag = strings.ToLower(tag)
	for _, proj := range projects {
		for _, t := range proj.Tags {
			if t == tag {
				results = append(results, proj)
				continue
			}
		}
	}
	return
}

// Projects returns the projects module
func Projects() *layout.Module {
	projectList, err := projects.LoadProjects()
	if err != nil {
		panic(err)
	}
	return &layout.Module{
		Name:            "Projects",
		ContentTemplate: getContentTemplate("projects"),
		Handler: func(pr *beepboop.PageRequest) interface{} {
			var v *projectsView
			tag := getTag(pr)
			if len(tag) > 0 {
				v = &projectsView{
					Tag:      tag,
					Projects: shuffleProjects(filterProjects(projectList, tag)),
				}
			} else {
				v = &projectsView{
					Projects: shuffleProjects(projectList),
				}
			}
			if len(v.Projects) == 0 {
				return nil
			}
			return v
		},
	}
}
