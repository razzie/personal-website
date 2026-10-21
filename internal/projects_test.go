package internal

import (
	"bytes"
	"html/template"
	"reflect"
	"strings"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestCardLinkVisibility(t *testing.T) {
	for _, url := range []string{
		"https://example.com/App.EXE?download=1#file",
		"https://example.com/tool-linux-amd64",
		"https://example.com/tool.zip",
		"https://example.com/tool.jar",
		"https://github.com/owner/repo/tree/main/binaries",
		"https://example.com/demo.mp4",
		"https://example.com/model.stl",
	} {
		t.Run(url, func(t *testing.T) {
			for _, label := range []string{"", "Open"} {
				project := Project{LinkGroups: []ProjectLinkGroup{{Links: []ProjectLink{
					{Name: "Full title", URL: url, CardLabel: label},
				}}}}
				links := project.CardLinks()
				if label == "" && len(links) != 0 || label != "" && len(links) != 1 {
					t.Errorf("cardLabel=%q: got %#v", label, links)
				}
			}
		})
	}
}

func TestCardLinksMetadata(t *testing.T) {
	var project Project
	err := yaml.Unmarshal([]byte(`
tags: [game]
detailsLabel: Explore the collection
linkGroups:
- name: Source code
  links:
  - name: Unlabeled source
    url: https://example.com/source
- name: Arbitrary heading
  links:
  - name: Windows download
    url: https://example.com/App.EXE?download=1
  - name: Full guide title
    url: https://example.com/guide
    cardLabel: Read guide
  - name: Another full title
    url: https://example.com/demo
    cardLabel: Open demo
- name: More links
  links:
  - name: Third destination
    url: https://example.com/extra
    cardLabel: Extra
`), &project)
	if err != nil {
		t.Fatal(err)
	}
	want := []ProjectLink{
		{Name: "Read guide", URL: "https://example.com/guide"},
		{Name: "Open demo", URL: "https://example.com/demo"},
	}
	if got := project.CardLinks(); !reflect.DeepEqual(got, want) {
		t.Errorf("CardLinks() = %#v, want %#v", got, want)
	}
	if got := project.LinkGroups[1].Links[1].Name; got != "Full guide title" {
		t.Errorf("CardLinks changed detail link name to %q", got)
	}
	if got := project.DetailsLabel(); got != "Explore the collection" {
		t.Errorf("DetailsLabel() = %q, want explicit label", got)
	}

	// Omitting metadata leaves links in details, regardless of headings or tags.
	project = Project{LinkGroups: project.LinkGroups[:1]}
	if got := project.CardLinks(); len(got) != 0 {
		t.Errorf("unlabeled card links = %#v, want none", got)
	}
	if got := project.DetailsLabel(); got != "Details" {
		t.Errorf("DetailsLabel() = %q, want Details", got)
	}
}

func TestRepositoryCardLinks(t *testing.T) {
	content := LoadContent("..")
	byID := map[string]Project{}
	for _, project := range content.Projects {
		byID[project.ID] = project
		links := project.CardLinks()
		if len(links) > 2 {
			t.Errorf("%s has %d card links", project.ID, len(links))
		}
	}
	// Riddle Solver's two executables and secondary references belong in
	// details, while the source stays directly accessible from its card.
	riddle := byID["riddle_solver"].CardLinks()
	if len(riddle) != 1 || riddle[0].Name != "Source" {
		t.Fatalf("Riddle Solver card links = %#v, want only Source", riddle)
	}
	// A repository collection should not arbitrarily promote two entries.
	utils := byID["golang_utils"]
	if len(utils.CardLinks()) != 0 || utils.DetailsLabel() != "All 6 repositories" {
		t.Errorf("Golang utils should expose its complete collection through details")
	}
	chess := byID["razchess"].CardLinks()
	if len(chess) != 2 || chess[0].Name != "Play" || chess[1].Name != "Source" {
		t.Errorf("RazChess should promote Play and Source, got %#v", chess)
	}
	for _, tc := range []struct {
		id         string
		wantLabels []string
	}{
		{"logic_circuit_simulator", []string{"Source"}},
		{"potato_game", []string{"Video", "Source"}},
		{"prepi", []string{"Source"}},
		{"process_manager", nil},
		{"raztalk", []string{"Source"}},
		{"razzgravitas", []string{"Source"}},
		{"razzie_messenger", nil},
		{"razziecraft_vanilla", nil},
		{"riddle_solver", []string{"Source"}},
		{"rts", nil},
		{"server_client_application", nil},
		{"windows_manager", nil},
		{"3d_prints", []string{"Download"}},
		{"magnetic_spinner", []string{"Download"}},
		{"labyrinth", []string{"Download"}},
		{"floating_islands", []string{"Source & downloads"}},
		{"powerman", []string{"Download"}},
		{"spyland", []string{"Download"}},
		{"ground_tool", []string{"Source & downloads"}},
		{"hexagon", []string{"Source & downloads"}},
		{"minecraft_clone", []string{"Source & downloads"}},
	} {
		t.Run(tc.id, func(t *testing.T) {
			project, ok := byID[tc.id]
			if !ok {
				t.Fatalf("missing project %q", tc.id)
			}
			var labels []string
			for _, link := range project.CardLinks() {
				labels = append(labels, link.Name)
			}
			if !reflect.DeepEqual(labels, tc.wantLabels) {
				t.Errorf("card labels = %v, want %v", labels, tc.wantLabels)
			}
		})
	}
}

func TestUnlabeledLinkRemainsInDetails(t *testing.T) {
	tmpl, err := template.ParseFiles("../templates/project.html")
	if err != nil {
		t.Fatal(err)
	}
	project := Project{LinkGroups: []ProjectLinkGroup{{Links: []ProjectLink{
		{Name: "Linux download", URL: "https://example.com/tool-linux-amd64"},
		{Name: "Source code", URL: "https://example.com/source", CardLabel: "Source"},
	}}}}
	for _, name := range []string{"project", "project-detail"} {
		var buf bytes.Buffer
		if err := tmpl.ExecuteTemplate(&buf, name, project); err != nil {
			t.Fatal(err)
		}
		if got := strings.Contains(buf.String(), `href="https://example.com/tool-linux-amd64"`); got != (name == "project-detail") {
			t.Errorf("%s: hidden download visibility = %v", name, got)
		}
		if !strings.Contains(buf.String(), `href="https://example.com/source"`) {
			t.Errorf("%s: missing source link", name)
		}
	}
}
