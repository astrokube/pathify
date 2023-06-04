package pathify

import (
	"github.com/astrokube/pathify/pathifier"
)

func Load[S pathifier.Type](content S, opts ...pathifier.PathifyOpt) pathifier.Pathifier[S] {
	return pathifier.Load[S](content, opts...)
}

func Array(opts ...pathifier.PathifyOpt) pathifier.Pathifier[[]any] {
	return pathifier.New[[]any](opts...)
}

func Map(opts ...pathifier.PathifyOpt) pathifier.Pathifier[map[string]any] {
	return pathifier.New[map[string]any](opts...)
}
