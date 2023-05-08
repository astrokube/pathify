package node

import "fmt"

type kind int

const (
	attribute kind = iota
	array
)

type attributes struct {
	path  string
	name  string
	value any
	kind  kind
}

type Node struct {
	attributes
	children Root
}

type partial struct {
	attributes
	child *partial
}

func (n *partial) EqualTo(other partial) bool {
	fail := n.path != other.path || n.name != other.name || n.value != other.value || (n.child == nil && other.child != nil) || (n.child != nil && other.child == nil)
	if fail {
		return false
	}
	if n.child != nil && other.child != nil {
		return n.child.EqualTo(*other.child)
	}
	return true
}

func (n *partial) addToBottom(child *partial) {
	if n.child != nil {
		n.child.addToBottom(child)
		return
	}
	n.child = child
}

func (n *partial) String() string {
	out := fmt.Sprintf("Path: %s, Name: %s", n.path, n.name)
	if n.value != nil {
		out += fmt.Sprintf(", Value:%v", n.value)
	}
	if n.child != nil {
		out += fmt.Sprintf(", Child:{ %s }", n.child.String())
	}
	return out
}
