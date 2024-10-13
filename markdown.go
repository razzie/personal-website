package main

import (
	"html/template"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Markdown string

func (md Markdown) ToHTML() template.HTML {
	htmlRenderer := html.NewRenderer(html.RendererOptions{
		Flags: html.CommonFlags | html.HrefTargetBlank | html.UseXHTML,
	})
	mdParser := parser.NewWithExtensions(parser.CommonExtensions | parser.NoEmptyLineBeforeBlock)
	html := markdown.ToHTML([]byte(md), mdParser, htmlRenderer)
	return template.HTML(html)
}
