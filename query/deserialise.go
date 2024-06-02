package query

import (
	_ "embed"
	"strings"

	"github.com/charmbracelet/log"

	yaml "gopkg.in/yaml.v3"
)

//go:embed ports.yml
var yamlString string

func parseYaml(yml string) PortsYAML {
	var portsYml PortsYAML

	err := yaml.Unmarshal([]byte(yml), &portsYml)
	if err != nil {
		log.Fatal("Could not parse static YML", "err", err)
	}

	portsYml.Maintainers = make(map[string]Maintainer)

	for _, v := range portsYml.RawMaintainers {
		var maintainer Maintainer
		err = v.Decode(&maintainer)
		if err != nil {
			log.Fatal("Could not parse maintainer", "err", err)
		}

		if maintainer.Name == nil {
			newName := strings.Clone(v.Anchor)
			maintainer.Name = &newName
		}
		portsYml.Maintainers[v.Anchor] = maintainer
	}

	for i := range portsYml.Ports {
		// NOTE: `v` takes a copy
		v := portsYml.Ports[i]
		currentMaintainers := make([]Maintainer, 0, len(v.RawCurrentMaintainers))

		for _, maintainer := range v.RawCurrentMaintainers {
			newMaintainer := portsYml.Maintainers[maintainer.Value]
			currentMaintainers = append(currentMaintainers, newMaintainer)
		}

		v.CurrentMaintainers = currentMaintainers

		// NOTE: Because `v` is a copy, we need to write it back
		portsYml.Ports[i] = v
	}

	return portsYml
}

var staticYAML = parseYaml(yamlString)

type Port struct {
	Name                  string      `yaml:"name"`
	Categories            []string    `yaml:"categories"`
	Platform              interface{} `yaml:"platform"`
	Color                 string      `yaml:"color"`
	RawCurrentMaintainers []yaml.Node `yaml:"current-maintainers"`
	CurrentMaintainers    []Maintainer
}

func (p Port) Maintainers() []Maintainer {
	return p.CurrentMaintainers
}

func (p Port) FilterValue() string {
	return p.Name
}

func (p Port) Render() string {
	return p.Name
}

type Maintainer struct {
	URL  string  `yaml:"url"`
	Name *string `yaml:"name"`
}

type Category struct {
	Key         string `yaml:"key"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Emoji       string `yaml:"emoji"`
}

type ShowcaseEntry struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Link        string `yaml:"link"`
}

type PortsYAML struct {
	RawMaintainers []yaml.Node `yaml:"collaborators"`
	Maintainers    map[string]Maintainer
	Ports          map[string]Port `yaml:"ports"`
	Categories     []Category      `yaml:"categories"`
	Showcases      []ShowcaseEntry `yaml:"showcases"`
}
