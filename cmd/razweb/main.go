package main

import (
	"net/http"
	"os"

	"github.com/razzie/beepboop"
	"github.com/razzie/gorzsony.com/assets"
	"github.com/razzie/gorzsony.com/pkg/layout"
	"github.com/razzie/gorzsony.com/pkg/modules"
)

func envOrDefault(envVar, defaultVal string) string {
	if val := os.Getenv(envVar); len(val) > 0 {
		return val
	}
	return defaultVal
}

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	cacheDir := envOrDefault("CACHE_DIR", "./cache")
	remoteDir := envOrDefault("REMOTE_DIR", "https://raw.githubusercontent.com/razzie/gorzsony.com/master/assets/")
	os.MkdirAll(cacheDir, os.ModeDir|os.ModePerm)

	loader := assets.NewAssetLoader(cacheDir, remoteDir)
	bio := modules.Hello(loader)
	projs := modules.Projects(loader)
	repos, stars := modules.Github(token)
	resume := modules.Resume(loader)

	mainPage := layout.CombineModules("/", "Gábor Görzsöny", bio, projs, repos, stars)
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
	srv.FaviconPNG, _ = assets.StaticAsset("img/favicon.png")
	srv.AddPages(
		beepboop.FSPage("/css/", assets.StaticFS()),
		beepboop.FSPage("/img/", loader),
		beepboop.FSPage("/js/", assets.StaticFS()),
		mainPage,
		layout.CombineModules("/tag/", "Gábor Görzsöny", projs, repos),
		layout.CombineModules("/project/", "Gábor Görzsöny", projs),
		layout.CombineModules("/resume", "Gábor Görzsöny - Resume", resume),
	)
	srv.Header = header

	http.ListenAndServe(":8080", srv)
}
