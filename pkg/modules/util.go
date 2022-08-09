package modules

import (
	"fmt"

	"github.com/razzie/gorzsony.com/assets"
)

func getContentTemplate(module string) string {
	t, err := assets.StaticAsset(fmt.Sprintf("template/%s.html", module))
	if err != nil {
		panic(err)
	}
	return string(t)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
