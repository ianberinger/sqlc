package golang

import (
	"strings"

	"github.com/ianberinger/sqlc/internal/config"
	"github.com/ianberinger/sqlc/internal/core"
)

type Struct struct {
	Table   core.FQN
	Name    string
	Fields  []Field
	Comment string
}

func StructName(name string, settings config.CombinedSettings) string {
	if rename := settings.Rename[name]; rename != "" {
		return rename
	}
	out := ""
	for _, p := range strings.Split(name, "_") {
		if p == "id" {
			out += "ID"
		} else {
			out += strings.Title(p)
		}
	}
	return out
}

func JSONTagName(name string, settings config.CombinedSettings) string {
	if settings.Go.JSONCamelCase {
		parts := strings.Split(name, "_")
		for i := range parts[1:] {
			parts[i+1] = strings.Title(parts[i+1])
		}
		return strings.Join(parts, "")
	}
	return name
}
