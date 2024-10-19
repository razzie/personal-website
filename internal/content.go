package internal

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Content struct {
	Hello       Markdown      `yaml:"hello"`
	Skills      Markdown      `yaml:"skills"`
	Experience  MarkdownSlice `yaml:"experience"`
	Projects    []Project     `yaml:"-"`
	ProjectTags []string      `yaml:"-"`
}

func LoadContent(dir string) (content Content) {
	contentRaw, err := os.ReadFile(filepath.Join(dir, "content.yaml"))
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(contentRaw, &content); err != nil {
		panic(err)
	}
	content.Projects, content.ProjectTags = loadProjects(filepath.Join(dir, "projects"))
	return
}
