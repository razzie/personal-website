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

	header := http.Header{
		"Content-Security-Policy": {
			"default-src 'none'",
			"script-src 'self' 'nonce-apply-scrollreveal'",
			"style-src 'self' 'unsafe-inline'",
			"img-src 'self' data:",
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
	srv.Header = header

	http.ListenAndServe("localhost:8080", srv)
}
