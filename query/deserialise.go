package query

import (
	_ "embed"

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

	return portsYml
}

var staticYAML = parseYaml(yamlString)

type Port struct {
	Name               string        `yaml:"name"`
	Categories         []string      `yaml:"categories"`
	Platform           interface{}   `yaml:"platform"`
	Color              string        `yaml:"color"`
	CurrentMaintainers []*Maintainer `yaml:"current-maintainers"`
}

func (p Port) Maintainers() []*Maintainer {
	return p.CurrentMaintainers
}

func (p Port) Render() string {
	return p.Name
}

type Maintainer struct {
	URL string `yaml:"url"`
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
	Maintainers []Maintainer    `yaml:"collaborators"`
	Ports       map[string]Port `yaml:"ports"`
	Categories  []Category      `yaml:"categories"`
	Showcases   []ShowcaseEntry `yaml:"showcases"`
}
