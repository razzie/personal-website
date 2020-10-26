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
	hello := modules.Hello()
	projs := modules.Projects()
	repos, stars := modules.Github(string(token))
	fs := internal.FS("")

	srv := beepboop.NewServer()
	srv.Layout = layout.Layout
	srv.FaviconPNG, _ = internal.Asset("img/favicon.png")
	srv.AddPages(
		beepboop.AssetFSPage("/css/", fs),
		beepboop.AssetFSPage("/img/", fs),
		beepboop.AssetFSPage("/js/", fs),
		layout.CombineModules("/", "Gábor Görzsöny", hello, projs, repos, stars),
		layout.CombineModules("/tag/", "Gábor Görzsöny", projs, repos),
		layout.CombineModules("/resume", "Gábor Görzsöny - Resume", modules.Resume()),
	)

	http.ListenAndServe("localhost:8080", srv)
}
