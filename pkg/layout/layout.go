package layout

import (
	"html/template"
	"net/http"

	"github.com/razzie/beepboop"
	"github.com/razzie/gorzsony.com/pkg/assets"
)

// Layout ...
var Layout beepboop.Layout

func init() {
	t, err := assets.Asset("template/layout.html")
	if err != nil {
		panic(err)
	}

	Layout = (*layout)(template.Must(template.New("layout").Parse(string(t))))
}

type layout template.Template

// BindTemplate creates a layout renderer function from a page template
func (l *layout) BindTemplate(pageTemplate string, stylesheets, scripts []string, meta map[string]string) (beepboop.LayoutRenderer, error) {
	cloneLayout, _ := (*template.Template)(l).Clone()
	tmpl, err := cloneLayout.New("page").Parse(pageTemplate)
	if err != nil {
		return nil, err
	}

	return func(w http.ResponseWriter, r *http.Request, title string, data interface{}, statusCode int) {
		view := struct {
			Title       string
			Base        string
			Stylesheets []string
			Scripts     []string
			Meta        map[string]string
			Data        interface{}
		}{
			Title:       title,
			Base:        beepboop.GetBase(r),
			Stylesheets: stylesheets,
			Scripts:     scripts,
			Meta:        meta,
			Data:        data,
		}

		w.WriteHeader(statusCode)
		tmpl.ExecuteTemplate(w, "layout", &view)
	}, nil
}
