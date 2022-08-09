package modules

import (
	"strings"

	"github.com/razzie/beepboop"
	"github.com/razzie/gorzsony.com/assets"
	"github.com/razzie/gorzsony.com/pkg/content"
	"github.com/razzie/gorzsony.com/pkg/layout"
)

type projectsView struct {
	Tag      string
	Projects []content.Project
}

/*func shuffleProjects(projects []content.Project) []content.Project {
	clone := append(projects[:0:0], projects...)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(clone), func(i, j int) { clone[i], clone[j] = clone[j], clone[i] })
	return clone
}*/

func filterProjects(projects []content.Project, tag string) (results []content.Project) {
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

func getTag(pr *beepboop.PageRequest) string {
	if strings.HasPrefix(pr.Request.URL.Path, "/tag/") {
		return pr.RelPath
	}
	return ""
}

func getProjectID(pr *beepboop.PageRequest) string {
	if strings.HasPrefix(pr.Request.URL.Path, "/project/") {
		return pr.RelPath
	}
	return ""
}

// Projects returns the projects module
func Projects(loader *assets.AssetLoader) *layout.Module {
	return &layout.Module{
		Name:            "Projects",
		ContentTemplate: getContentTemplate("projects"),
		Handler: func(pr *beepboop.PageRequest) interface{} {
			var v *projectsView
			projects := loader.Content().Projects
			tag := getTag(pr)
			projectID := getProjectID(pr)
			switch {
			case len(tag) > 0:
				v = &projectsView{
					Tag:      tag,
					Projects: filterProjects(projects, tag),
				}
			case len(projectID) > 0:
				for _, project := range projects {
					if project.ID == projectID {
						v = &projectsView{
							Projects: []content.Project{project},
						}
						break
					}
				}
			default:
				v = &projectsView{
					Projects: projects,
				}
			}
			if v == nil || len(v.Projects) == 0 {
				return nil
			}
			return v
		},
	}
}
