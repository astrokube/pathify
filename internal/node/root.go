package node

import "strconv"

type Root map[string]*Node

func (r Root) PrettyPrint() string {
	out := ""
	for _, node := range r {
		out += node.Tree() + "\n"
	}
	return out
}

func (r Root) Add(path string, value any) {
	partialNode := pathToNode(path, value)
	r.merge(partialNode)
}

func (r Root) merge(n partial) {
	if _, ok := r[n.name]; !ok {
		r[n.name] = &Node{
			attributes: attributes{
				path:  n.path,
				name:  n.name,
				value: n.value,
				kind:  n.kind,
			},
			children: map[string]*Node{},
		}
	}
	if n.child != nil {
		r[n.name].children.merge(*n.child)
	}
}

func (r Root) asArray() []any {
	out := make([]any, 0)
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
			out[index] = v.children.AsMap()
		}
	}
	return out
}

func (r Root) AsMap() map[string]any {
	if len(r) == 0 {
		return map[string]any{}
	}
	out := make(map[string]any)
	for _, v := range r {
		if v.kind == array {
			out[v.name] = v.children.asArray()
		} else if v.value != nil {
			out[v.name] = v.value
		} else {
			out[v.name] = v.children.AsMap()
		}
	}
	return out
}
