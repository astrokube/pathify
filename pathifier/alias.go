package pathifier

import "github.com/astrokube/pathify/pathifier/internal"

const (
	YAML = internal.YAML
	JSON = internal.JSON
)

var (
	AsJSON = internal.WithOutputFormat(JSON)
	AsYAML = internal.WithOutputFormat(YAML)
)
