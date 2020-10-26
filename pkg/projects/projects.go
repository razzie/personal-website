package projects

import (
	"encoding/xml"
	"html/template"
	"strings"

	"github.com/razzie/gorzsony.com/internal"
)

// Project contains data about one of my hobby projects
type Project struct {
	ID          string
	Name        string
	Year        string
	ImageURL    string
	Tags        []string
	Description template.HTML
}

type xmlProjects struct {
	XMLName  xml.Name     `xml:"projects"`
	Projects []xmlProject `xml:"project"`
}

type xmlProject struct {
	XMLName     xml.Name       `xml:"project"`
	ID          string         `xml:"id"`
	Name        string         `xml:"name"`
	Year        string         `xml:"year"`
	ImageURL    string         `xml:"img"`
	Tags        string         `xml:"tags"`
	Description xmlDescription `xml:"description"`
}

type xmlDescription struct {
	XMLName  xml.Name `xml:"description"`
	InnerXML string   `xml:",innerxml"`
}

func newProject(proj xmlProject) Project {
	return Project{
		ID:          proj.ID,
		Name:        proj.Name,
		Year:        proj.Year,
		ImageURL:    proj.ImageURL,
		Tags:        strings.Fields(proj.Tags),
		Description: template.HTML(proj.Description.InnerXML),
	}
}

// LoadProjects parses projects.xml and returns the projects in a slice
func LoadProjects() ([]Project, error) {
	data, err := internal.Asset("xml/projects.xml")
	if err != nil {
		return nil, err
	}

	var projects xmlProjects
	err = xml.Unmarshal(data, &projects)
	if err != nil {
		return nil, err
	}

	var result []Project
	for _, proj := range projects.Projects {
		result = append(result, newProject(proj))
	}
	return result, nil
}
