package modules

import (
	"math/rand"
	"sort"
	"strconv"
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

func orderProjectsByYear(projects []projects.Project) {
	getYear := func(i int) int {
		years := strings.Split(projects[i].Year, "-")
		year, _ := strconv.Atoi(years[len(years)-1])
		return year
	}
	sort.SliceStable(projects, func(i, j int) bool {
		return getYear(i) > getYear(j)
	})
}

// Projects returns the projects module
func Projects() *layout.Module {
	projectList, err := projects.LoadProjects()
	if err != nil {
		panic(err)
	}
	orderProjectsByYear(projectList)
	return &layout.Module{
		Name:            "Projects",
		ContentTemplate: getContentTemplate("projects"),
		Handler: func(pr *beepboop.PageRequest) interface{} {
			var v *projectsView
			tag := getTag(pr)
			if len(tag) > 0 {
				v = &projectsView{
					Tag:      tag,
					Projects: filterProjects(projectList, tag),
				}
			} else {
				v = &projectsView{
					Projects: projectList,
				}
			}
			if len(v.Projects) == 0 {
				return nil
			}
			return v
		},
	}
}
