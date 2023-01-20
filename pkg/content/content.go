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
	Personal         string        `yaml:"personal"`
	PersonalHTML     template.HTML `yaml:"-"`
	Professional     string        `yaml:"professional"`
	ProfessionalHTML template.HTML `yaml:"-"`
}

type Resume struct {
	Header     string        `yaml:"header"`
	HeaderHTML template.HTML `yaml:"-"`
	Sections   []struct {
		Title string `yaml:"title"`
		Items []struct {
			Title           string        `yaml:"title"`
			Year            string        `yaml:"year"`
			Company         string        `yaml:"company"`
			URL             string        `yaml:"url"`
			Description     string        `yaml:"description"`
			DescriptionHTML template.HTML `yaml:"-"`
			Faded           bool          `yaml:"faded,omitempty"`
		}
	} `yaml:"sections"`
}

type Project struct {
	ID              string        `yaml:"id"`
	Name            string        `yaml:"name"`
	Year            string        `yaml:"year"`
	Tags            []string      `yaml:"tags"`
	Description     string        `yaml:"description"`
	DescriptionHTML template.HTML `yaml:"-"`
	LinkGroups      []struct {
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
	convertMarkdownToHTML(content)
	return content, nil
}

func orderProjectsByYear(projects []Project) {
	getYear := func(i int) int {
		years := strings.Split(projects[i].Year, "-")
		year, _ := strconv.Atoi(strings.TrimSpace(years[len(years)-1]))
		return year
	}
	sort.SliceStable(projects, func(i, j int) bool {
		return getYear(i) > getYear(j)
	})
}

func convertMarkdownToHTML(content *Content) {
	htmlFlags := html.CommonFlags | html.HrefTargetBlank | html.UseXHTML
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	convert := func(input string, output *template.HTML) {
		extensions := parser.CommonExtensions | parser.NoEmptyLineBeforeBlock
		parser := parser.NewWithExtensions(extensions)

		html := markdown.ToHTML([]byte(input), parser, renderer)
		*output = template.HTML(html)
	}

	// bio
	convert(content.Bio.Personal, &content.Bio.PersonalHTML)
	convert(content.Bio.Professional, &content.Bio.ProfessionalHTML)

	// resume
	convert(content.Resume.Header, &content.Resume.HeaderHTML)
	for _, section := range content.Resume.Sections {
		for i, item := range section.Items {
			convert(item.Description, &section.Items[i].DescriptionHTML)
		}
	}

	// projects
	for i, project := range content.Projects {
		convert(project.Description, &content.Projects[i].DescriptionHTML)
	}
}
