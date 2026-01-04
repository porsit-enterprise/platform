package foundation

import (
	"github.com/porsit-enterprise/platform/foundation/configuration"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

type Foundation struct {
	Configuration configuration.Properties
	Settings      any
}
