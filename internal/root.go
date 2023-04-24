package internal

import (
	"strconv"
)

type Root map[string]node

func (r Root) Get() map[string]any {
	if len(r) == 0 {
		return map[string]any{}
	}
	out := make(map[string]any)
	for _, v := range r {
		if v.kind == array {
			out[v.name] = Root(v.children).getItems()
		} else if v.value != nil {
			out[v.name] = v.value
		} else {
			out[v.name] = Root(v.children).Get()
		}
	}
	return out
}

func (r Root) getItems() []any {
	out := []any{}
	if len(r) == 0 {
		return out
	}
	for _, v := range r {
		index, err := strconv.Atoi(v.name)
		increment := index - len(out) + 1
		if increment > 0 {
			appendedArray := make([]any, increment)
			appendedArray[len(appendedArray)-1] = v
			out = append(out, appendedArray...)
		}
		if err == nil {
			out[index] = Root(v.children).Get()
		}
	}
	return out
}

func (r Root) appendToRoot(nodes []node) {
	n, ok := r[nodes[0].name]
	if !ok {
		n = nodes[0]
		n.children = make(map[string]node)
	}
	r[nodes[0].name] = n
	if len(nodes) > 1 {
		Root(r[nodes[0].name].children).appendToRoot(nodes[1:])
	}
}

func (r Root) Append(path string, value any) {
	nodes := parsePath(path, value)
	r.appendToRoot(nodes)
}
