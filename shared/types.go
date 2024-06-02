package shared

import "github.com/catppuccin/cli/query"

type Context struct {
	Interactive bool
	Datastore   query.Datastore
}
