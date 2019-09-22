package main

import (
	"encoding/xml"
	"html/template"
)

type Project struct {
	ID          string
	Name        string
	ImageUrl    string
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
	ImageUrl    string         `xml:"img"`
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
		ImageUrl:    proj.ImageUrl,
		Description: template.HTML(proj.Description.InnerXML)}
}

func LoadProjects() ([]Project, error) {
	data, err := Asset("projects.xml")
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
