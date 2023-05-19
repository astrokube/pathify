package pathify

import (
	"github.com/astrokube/pathify/pathify"
)

func Load[S pathify.Type](content S, opts ...pathify.PathifyOpt) pathify.Pathifier[S] {
	return pathify.Load[S](content, opts...)
}

func New(opts ...pathify.PathifyOpt) pathify.Pathifier[map[string]any] {
	return pathify.New(opts...)
}
