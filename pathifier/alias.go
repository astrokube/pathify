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

type SetterOpt = internal.SetterOpt

var (
	WithFuncPrefix   = internal.WithFuncPrefix
	WithStringPrefix = internal.WithStringPrefix
)
