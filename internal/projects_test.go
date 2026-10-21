package internal

import "testing"

func TestExecutableLinks(t *testing.T) {
	for _, tc := range []struct {
		url  string
		want bool
	}{
		{"https://example.com/app.exe", true},
		{"https://example.com/App.EXE?download=1#file", true},
		{"https://github.com/owner/repo/blob/main/Process%20Manager.exe", true},
		{"https://example.com/tool.zip", false},
		{"https://example.com/?file=app.exe", false},
		{"https://example.com/app.exe/releases", false},
	} {
		if got := (ProjectLink{URL: tc.url}).IsExecutable(); got != tc.want {
			t.Errorf("IsExecutable(%q) = %v, want %v", tc.url, got, tc.want)
		}
	}
}

func TestCardLinks(t *testing.T) {
	content := LoadContent("..")
	byID := map[string]Project{}
	for _, project := range content.Projects {
		byID[project.ID] = project
		links := project.CardLinks()
		if len(links) > 2 {
			t.Errorf("%s has %d card links", project.ID, len(links))
		}
		for _, link := range links {
			if link.IsExecutable() {
				t.Errorf("%s exposes an executable on its card: %s", project.ID, link.URL)
			}
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
}
