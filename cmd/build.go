package cmd

import (
	"os"
	"path/filepath"

	"github.com/razzie/personal-website/internal"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(buildCmd)
}

var buildCmd = &cobra.Command{
	Use:  "build [flags]",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		dist, _ := cmd.Flags().GetString("dist")
		content := internal.LoadContent(".")
		navPages := []internal.Page{
			{
				ID:       "",
				Title:    "Hello",
				Template: "columns",
				Data:     content.Hello.ToHTML(),
			},
			{
				ID:       "skills",
				Title:    "Skills",
				Template: "columns",
				Data:     content.Skills.ToHTML(),
			},
			{
				ID:       "experience",
				Title:    "Experience",
				Template: "timeline",
				Data:     content.Experience.ToHTML(),
			},
			{
				ID:       "projects",
				Title:    "Projects",
				Template: "projects",
				Data: map[string]any{
					"Projects": content.Projects,
					"Tags":     content.ProjectTags,
				},
			},
		}
		render := internal.LoadTemplateRenderer(navPages)

		if err := os.MkdirAll(filepath.Join(dist, "projects"), 0770); err != nil {
			return err
		}

		staticFiles, _ := filepath.Glob("static/*")
		for _, filename := range staticFiles {
			if err := internal.CopyFile(filename, filepath.Join(dist, filepath.Base(filename))); err != nil {
				return err
			}
		}

		projectImages, _ := filepath.Glob("projects/*.webp")
		for _, filename := range projectImages {
			if err := internal.CopyFile(filename, filepath.Join(dist, "projects", filepath.Base(filename))); err != nil {
				return err
			}
		}

		for _, page := range navPages {
			if err := createPage(dist, page, render); err != nil {
				return err
			}
		}

		for _, p := range content.Projects {
			page := internal.Page{
				Title:    p.Name,
				ID:       "projects/id/" + p.ID,
				Template: "project",
				Data:     p,
			}
			if err := createPage(dist, page, render); err != nil {
				return err
			}
		}

		for _, tag := range content.ProjectTags {
			view := map[string]any{
				"Projects": internal.FilterProjectsByTag(content.Projects, tag),
				"Tags":     content.ProjectTags,
				"Tag":      tag,
			}
			page := internal.Page{
				Title:    "Projects (" + tag + ")",
				ID:       "projects/tag/" + tag,
				Template: "projects",
				Data:     view,
			}
			if err := createPage(dist, page, render); err != nil {
				return err
			}
		}

		return nil
	},
}

func createPage(root string, page internal.Page, render internal.TemplateRenderer) error {
	filename := filepath.Join(root, page.ID, "index.html")
	if err := os.MkdirAll(filepath.Dir(filename), 0770); err != nil {
		return err
	}
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return render(f, page)
}
