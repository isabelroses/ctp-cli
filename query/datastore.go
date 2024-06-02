package query

type Datastore struct {
	yml PortsYAML
}

func NewDatastore() Datastore {
	return Datastore{
		yml: staticYAML,
	}
}

func (d *Datastore) Ports() []Port {
	return mapValues(d.yml.Ports)
}

func FilterMaintained[T HasMaintainers](input []T) []T {
	return executeFilter(input, func(item T) bool {
		return len(item.Maintainers()) > 0
	})
}

func FilterUnmaintained[T HasMaintainers](input []T) []T {
	return executeFilter(input, func(item T) bool {
		return len(item.Maintainers()) == 0
	})
}

type HasMaintainers interface {
	Maintainers() []*Maintainer
}
