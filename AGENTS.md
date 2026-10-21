# Working on this repository

## Scope and collaboration

This is a Go static-site generator for a personal website. Keep changes focused
on the requested task and follow the existing implementation. Layout and styling
may be changing independently; inspect the current files rather than treating
these notes as a design specification.

Check `git status --short` before editing. Preserve unrelated changes, including
work from other agents. Do not reset, revert, or overwrite someone else's work.
Keep documentation focused on architecture, commands, and content conventions
rather than detailed descriptions of the current appearance.

## Architecture

- `main.go` calls the Cobra CLI in `cmd/`.
- `cmd/root.go` defines the shared `--output` / `-o` flag (default: `dist`).
- `cmd/build.go` loads content from the repository root, constructs pages, renders
  them as `<output>/<page ID>/index.html`, and copies assets.
- `cmd/serve.go` serves the output directory using Go's `http.FileServer`.
  It does not generate pages or watch source files.
- `internal/content.go` defines and loads the main YAML content model.
- `internal/projects.go` loads project records, derives their IDs and tags,
  filters by tag, and groups projects by starting year, newest first.
- `internal/markdown.go` converts Markdown to HTML with gomarkdown.
- `internal/templates.go` parses `templates/*.html`, registers `CallTemplate`,
  and renders pages through the `layout` template.
- `internal/copy.go` copies assets, preferring hard links when possible.

The build generates the home page, skills and experience pages, a projects index,
individual pages under `/projects/id/<id>/`, and filtered pages under
`/projects/tag/<tag>/`. Navigation pages are defined in `cmd/build.go`.
There is no database or application backend in production: the output is static.

## Build and local development

Use the Go version declared in `go.mod`; the Docker builder also pins a Go
version. Dependencies are checked into `vendor/`. There is no npm installation
or JavaScript bundling step.

Run commands from the repository root because content, template, and asset paths
are relative to the working directory:

```sh
make build
make run
```

`make build` compiles `./personal-website` using vendored dependencies and then
generates `dist/`. `make run` performs that build before serving on `:8080`.
After source edits, rebuild and refresh the browser; there is no live reload.
If the preview server is already running, rebuilding its output is sufficient.

Equivalent commands, with explicit output and listen-address examples:

```sh
go build -mod=vendor -o personal-website .
./personal-website build -o dist
./personal-website serve -o dist --addr 127.0.0.1:8080
```

`build` writes into the output directory without removing stale files. To verify
removals or renamed routes, generate into a fresh temporary output directory.
Build output and the compiled executable are ignored by Git; edit their source
files instead. Output assets may be hard-linked to sources, so editing an asset
inside `dist/` can also alter its source.

## Content and images

- Edit general site text in `content.yaml`. Its schema is defined by the structs
  and YAML tags in `internal/content.go`; update consumers when changing fields.
- Add projects as `projects/<id>.yaml`. The filename stem becomes the project ID
  and public URL, so renaming it changes links.
- Project fields are `name`, `year`, `tags`, `description`, and `linkGroups`.
  Use existing records as examples for nested link groups and YAML block text.
- Pair each project record with `projects/<id>.webp`. Images are referenced by
  convention, not by an image field in the YAML.
- Generate WebP images with `cwebp` at quality **60**:

  ```sh
  cwebp -q 60 input.png -o projects/<id>.webp
  ```

- Years can be a single year or a range such as `2023 - 2024`. Ordering and
  grouping use the first year; tags are collected and sorted automatically.
- Markdown fields can contain HTML, and Markdown output is marked trusted HTML.
  Treat repository content as trusted author input; do not reuse that rendering
  path for untrusted user submissions without reviewing sanitization.

## Templates and browser code

Keep presentation in `templates/` and browser assets in `static/`. Go templates
use named `define` blocks, with `layout` as the shared entry point. Follow existing
template data contracts when changing Go structs or page construction.

The generator copies only top-level `static/*` files and `projects/*.webp`
images; adding nested asset directories requires a corresponding build change.
Links and asset URLs assume hosting at the domain root.

HTMX boosts navigation. Browser behavior that needs initialization after page
changes should handle both `DOMContentLoaded` and `htmx:afterSwap`, as the
existing scripts do. Avoid duplicate listeners and clean up observers when
reinitializing. Keep directly loaded pages and ordinary links functional.

Do not edit vendored Go packages or bundled third-party browser libraries for
routine feature changes. For an intentional Go dependency update, keep `go.mod`,
`go.sum`, and `vendor/` consistent using the Go tooling.

## Validation

For content, template, or generator changes, run `make build` and inspect the
generated pages relevant to the change. Check build logs too: malformed project
YAML can be logged and skipped without failing the entire build.

For Go changes, run `gofmt` on modified Go files and use:

```sh
go test -mod=vendor ./...
go vet -mod=vendor ./...
```

`internal/projects_test.go` covers executable-link detection and project card
links, including checks against repository content. Successful `go test` alone
does not verify generated HTML or browser behavior. Add focused tests when
changing nontrivial generator logic. For browser changes, preview the affected
routes, check direct loading and HTMX navigation, and inspect the browser console.
For documentation-only edits, verify commands and paths against the source;
no site rebuild is necessary.

## Deployment

`.github/workflows/deploy.yml` runs on pushes to `master`. It uses the Go version
from `go.mod`, compiles the generator, builds `dist/`, adds a `CNAME` file for
`gorzsony.com`, and publishes the output to the `gh-pages` branch.
Keep source changes separate from generated deployment artifacts.
