package pathify

import (
	"github.com/astrokube/pathify/pathifier"
)

func Load[S pathifier.Type](content S, opts ...pathifier.PathifyOpt) pathifier.Pathifier[S] {
	return pathifier.Load[S](content, opts...)
}

func New(opts ...pathifier.PathifyOpt) pathifier.Pathifier[map[string]any] {
	return pathifier.New(opts...)
}
