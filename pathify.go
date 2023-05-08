package pathify

import (
	"github.com/astrokube/pathify/internal/node"
)

type Node interface {
	Set(path string, value any) Node
	Map() map[string]any
	PrettyPrint() string
}

type processor struct {
	root node.Root
}

func New() Node {
	return &processor{
		root: node.Root{},
	}
}

func (p *processor) Set(path string, value any) Node {
	p.root.Add(path, value)
	return p
}

func (p *processor) Map() map[string]any {
	return p.root.AsMap()
}

func (p *processor) PrettyPrint() string {
	return p.root.PrettyPrint()
}
