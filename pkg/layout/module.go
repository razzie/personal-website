package layout

import (
	"html/template"
	"strings"

	"github.com/razzie/beepboop"
	"github.com/razzie/gorzsony.com/pkg/assets"
)

var modulePageT string

func init() {
	t, err := assets.Asset("template/layout_modules.html")
	if err != nil {
		panic(err)
	}
	modulePageT = string(t)
}

// Module ...
type Module struct {
	Name            string
	ContentTemplate string
	Handler         func(pb *beepboop.PageRequest) interface{}
}

// CombineModules ...
func CombineModules(path, title string, modules ...*Module) *beepboop.Page {
	templates := make([]*template.Template, len(modules))
	for i, module := range modules {
		templates[i] = template.Must(template.New("").Parse(module.ContentTemplate))
	}
	handler := func(pr *beepboop.PageRequest) *beepboop.View {
		var results []template.HTML
		for i, module := range modules {
			var w strings.Builder
			var data interface{}
			if module.Handler != nil {
				data = module.Handler(pr)
				if data == nil {
					pr.Logf("skipping module: %s", module.Name)
					continue
				}
			}
			if err := templates[i].Execute(&w, data); err != nil {
				pr.Logf("module %s error: %s", module.Name, err.Error())
				continue
			}
			results = append(results, template.HTML(w.String()))
		}
		return pr.Respond(results)
	}
	return &beepboop.Page{
		Path:            path,
		Title:           title,
		ContentTemplate: modulePageT,
		Handler:         handler,
	}
}
