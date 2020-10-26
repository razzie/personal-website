package projects

import (
	"encoding/xml"

	"github.com/razzie/gorzsony.com/internal"
)

// TimelineYear ...
type TimelineYear struct {
	Year     string            `xml:"year"`
	Projects []TimelineProject `xml:"projects"`
}

// TimelineProject ...
type TimelineProject struct {
	Name string `xml:"name"`
	Link string `xml:"link"`
}

type xmlTimeline struct {
	XMLName xml.Name          `xml:"timeline"`
	Years   []xmlTimelineYear `xml:"year"`
}

type xmlTimelineYear struct {
	XMLName  xml.Name            `xml:"year"`
	Year     string              `xml:"year"`
	Projects xmlTimelineProjects `xml:"projects"`
}

type xmlTimelineProjects struct {
	XMLName  xml.Name          `xml:"projects"`
	Projects []TimelineProject `xml:"project"`
}

// LoadTimeline parses project_timeline.xml and returns the years in a slice
func LoadTimeline() ([]TimelineYear, error) {
	data, err := internal.Asset("xml/project_timeline.xml")
	if err != nil {
		return nil, err
	}

	var timeline xmlTimeline
	err = xml.Unmarshal(data, &timeline)
	if err != nil {
		return nil, err
	}

	var result []TimelineYear
	for _, year := range timeline.Years {
		result = append(result, TimelineYear{
			Year:     year.Year,
			Projects: year.Projects.Projects,
		})
	}
	return result, nil
}
