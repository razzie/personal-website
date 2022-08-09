package content

import (
	"html/template"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"gopkg.in/yaml.v2"
)

type Bio struct {
	Personal     template.HTML `yaml:"personal"`
	Professional template.HTML `yaml:"professional"`
}

type Resume struct {
	Title     string        `yaml:"title"`
	Objective template.HTML `yaml:"objective"`
	Sections  []struct {
		Title string `yaml:"title"`
		Items []struct {
			Title       string        `yaml:"title"`
			Year        string        `yaml:"year"`
			Company     string        `yaml:"company"`
			URL         string        `yaml:"url"`
			Description template.HTML `yaml:"description"`
			Faded       bool          `yaml:"faded,omitempty"`
		}
	} `yaml:"sections"`
}

type Project struct {
	ID          string        `yaml:"id"`
	Name        string        `yaml:"name"`
	Year        string        `yaml:"year"`
	Tags        []string      `yaml:"tags"`
	Description template.HTML `yaml:"description"`
	LinkGroups  []struct {
		Name  string `yaml:"name"`
		Links []struct {
			Name string `yaml:"name"`
			URL  string `yaml:"url"`
		}
	} `yaml:"linkGroups"`
}

type Content struct {
	Bio      Bio       `yaml:"bio"`
	Resume   Resume    `yaml:"resume"`
	Projects []Project `yaml:"projects"`
}

func LoadContent(r io.Reader) (*Content, error) {
	bytes, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	content := &Content{}
	if err := yaml.Unmarshal(bytes, content); err != nil {
		return nil, err
	}
	orderProjectsByYear(content.Projects)
	convertMarkdown(content)
	return content, nil
}

func orderProjectsByYear(projects []Project) {
	getYear := func(i int) int {
		years := strings.Split(projects[i].Year, "-")
		year, _ := strconv.Atoi(years[len(years)-1])
		return year
	}
	sort.SliceStable(projects, func(i, j int) bool {
		return getYear(i) > getYear(j)
	})
}

func convertMarkdown(content *Content) {
	htmlFlags := html.CommonFlags | html.HrefTargetBlank | html.UseXHTML
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	convert := func(text *template.HTML) {
		extensions := parser.CommonExtensions | parser.NoEmptyLineBeforeBlock
		parser := parser.NewWithExtensions(extensions)

		html := markdown.ToHTML([]byte(*text), parser, renderer)
		*text = template.HTML(html)
	}

	// bio
	convert(&content.Bio.Personal)
	convert(&content.Bio.Professional)

	// resume
	convert(&content.Resume.Objective)
	for _, section := range content.Resume.Sections {
		for i := range section.Items {
			convert(&section.Items[i].Description)
		}
	}

	// projects
	for i := range content.Projects {
		convert(&content.Projects[i].Description)
	}
}
