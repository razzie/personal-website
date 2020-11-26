package modules

import (
	"fmt"
	"strings"

	"github.com/razzie/beepboop"
	"github.com/razzie/gorzsony.com/pkg/assets"
)

func getContentTemplate(module string) string {
	t, err := assets.Asset(fmt.Sprintf("template/%s.html", module))
	if err != nil {
		panic(err)
	}
	return string(t)
}

func getTag(pr *beepboop.PageRequest) string {
	if strings.HasPrefix(pr.Request.URL.Path, "/tag/") {
		return pr.RelPath
	}
	return ""
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
