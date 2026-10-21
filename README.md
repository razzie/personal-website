# Personal website

Gábor Görzsöny's personal website, built with a small Go static-site generator.
It reads YAML content, renders Go HTML templates, and copies assets into `dist/`.
The generated site uses Pico CSS and HTMX, with no database or Node.js build step.

## Build and preview

Install the Go version specified in `go.mod` and Make, then run from the repository root:

```sh
make build  # Compile the generator and generate dist/
make run    # Build, then serve at http://localhost:8080
```

Rebuild after editing sources; the preview server has no live reload.

## Where things live

- `content.yaml`: main site content, including Markdown text.
- `projects/`: individual project YAML files and matching WebP thumbnails.
- `templates/`: Go HTML templates.
- `static/`: CSS, JavaScript, and other assets copied into the output.
- `cmd/` and `internal/`: generator, preview server, and content processing.

Project thumbnails use the same filename stem as their YAML file. WebP images
are generated with `cwebp` at quality 60 (`cwebp -q 60 input.png -o output.webp`).

Pushes to `master` trigger GitHub Actions to build the site and publish `dist/`
to GitHub Pages at [gorzsony.com](https://gorzsony.com).

See [AGENTS.md](AGENTS.md) for development guidance.
