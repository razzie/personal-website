package main

import (
	"encoding/xml"
)

type Project struct {
	ID          string `xml:"id"`
	Name        string `xml:"name"`
	ImageUrl    string `xml:"img"`
	Description string `xml:"description"`
}

type xmlProjects struct {
	XMLName  xml.Name  `xml:"projects"`
	Projects []Project `xml:"project"`
}

type xmlProject struct {
	Project
	XMLName xml.Name `xml:"project"`
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

	return projects.Projects, nil
}
