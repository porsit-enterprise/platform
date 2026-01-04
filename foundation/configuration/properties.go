package configuration

import "github.com/porsit-enterprise/platform/foundation/configuration/entities"

//──────────────────────────────────────────────────────────────────────────────────────────────────

type Properties struct {
	Version        string                  `yaml:"version"`
	Infrastructure entities.Infrastructure `yaml:"infrastructure"`
	Provider       entities.Provider       `yaml:"provider"`
}
