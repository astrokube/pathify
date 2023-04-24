package pathify

import (
	"github.com/astrokube/pathify/internal"
)

type Node interface {
	Set(path string, value any) Node
	Map() map[string]any
}

type processor struct {
	root internal.Root
}

func New() Node {
	return &processor{
		root: internal.Root{},
	}
}

func (p *processor) Set(path string, value any) Node {
	p.root.Append(path, value)
	return p
}

func (p *processor) Map() map[string]any {
	return p.root.Get()
}
