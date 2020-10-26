package main

import (
	"net/http"

	"github.com/razzie/beepboop"
	"github.com/razzie/gorzsony.com/internal"
	"github.com/razzie/gorzsony.com/pkg/layout"
	"github.com/razzie/gorzsony.com/pkg/modules"
)

func main() {
	token, _ := internal.Asset("github.token")
	projs := modules.Projects()
	repos, stars := modules.Github(string(token))
	fs := internal.FS("")

	mainPage := layout.CombineModules("/", "Gábor Görzsöny", modules.Hello(), projs, repos, stars)
	mainPageHandler := mainPage.Handler
	mainPage.Handler = func(pr *beepboop.PageRequest) *beepboop.View {
		if len(pr.RelPath) > 0 {
			return pr.RedirectView("/")
		}
		return mainPageHandler(pr)
	}

	srv := beepboop.NewServer()
	srv.Layout = layout.Layout
	srv.FaviconPNG, _ = internal.Asset("img/favicon.png")
	srv.AddPages(
		beepboop.AssetFSPage("/css/", fs),
		beepboop.AssetFSPage("/img/", fs),
		beepboop.AssetFSPage("/js/", fs),
		mainPage,
		layout.CombineModules("/tag/", "Gábor Görzsöny", projs, repos),
		layout.CombineModules("/timeline", "Gábor Görzsöny - Project timeline", modules.Timeline()),
		layout.CombineModules("/resume", "Gábor Görzsöny - Resume", modules.Resume()),
	)

	http.ListenAndServe("localhost:8080", srv)
}
