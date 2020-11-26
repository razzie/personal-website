package main

import (
	"net/http"

	"github.com/razzie/beepboop"
	"github.com/razzie/gorzsony.com/pkg/assets"
	"github.com/razzie/gorzsony.com/pkg/layout"
	"github.com/razzie/gorzsony.com/pkg/modules"
)

func main() {
	token, _ := assets.Asset("github.token")
	projs := modules.Projects()
	repos, stars := modules.Github(string(token))
	fs := assets.FS("")

	mainPage := layout.CombineModules("/", "Gábor Görzsöny", modules.Hello(), projs, repos, stars)
	mainPageHandler := mainPage.Handler
	mainPage.Handler = func(pr *beepboop.PageRequest) *beepboop.View {
		if len(pr.RelPath) > 0 {
			return pr.RedirectView("/")
		}
		return mainPageHandler(pr)
	}

	header := http.Header{
		"Content-Security-Policy": {
			"script-src 'self' 'nonce-apply-scrollreveal'",
			"style-src 'self' 'unsafe-inline'",
			"img-src 'self' data: image/svg+xml",
			"frame-ancestors 'none'",
			"base-uri 'self'",
			"form-action 'self'",
		},
		"Strict-Transport-Security": {"max-age=63072000"},
		"X-Content-Type-Options":    {"nosniff"},
		"X-Frame-Options":           {"DENY"},
		"X-XSS-Protection":          {"1", "mode=block"},
	}

	srv := beepboop.NewServer()
	srv.Layout = layout.Layout
	srv.FaviconPNG, _ = assets.Asset("img/favicon.png")
	srv.AddPages(
		beepboop.AssetFSPage("/css/", fs),
		beepboop.AssetFSPage("/img/", fs),
		beepboop.AssetFSPage("/js/", fs),
		mainPage,
		layout.CombineModules("/tag/", "Gábor Görzsöny", projs, repos),
		layout.CombineModules("/timeline", "Gábor Görzsöny - Project timeline", modules.Timeline()),
		layout.CombineModules("/resume", "Gábor Görzsöny - Resume", modules.Resume()),
	)
	srv.Header = header

	http.ListenAndServe("localhost:8080", srv)
}
