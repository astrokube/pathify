package internal

import (
	"fmt"
	"reflect"
	"strconv"
)

type Kind int32

const (
	Node Kind = iota
	Array
)

func (k Kind) String() string {
	if k == Node {
		return "Node"
	}
	return "Array"
}

type Mutator struct {
	name  string
	index string
	child *Mutator
	kind  Kind
	value any
}

func (m *Mutator) Child() *Mutator {
	return m.child
}

func (m *Mutator) Kind() Kind {
	return m.kind
}

func (m *Mutator) applyValue(in any) any {
	val := reflect.ValueOf(m.value)
	switch val.Kind() {
	case reflect.Struct:
		return nil
	case reflect.Func:
		x := reflect.TypeOf(m.value)
		if x.NumIn() != 1 || x.NumOut() != 1 {
			return nil
		}
		inVal := reflect.ValueOf(in)
		if x.In(0).Kind() != inVal.Kind() {
			return nil
		}
		out := val.Call([]reflect.Value{inVal})
		return out[0].String()
	default:
		return m.value
	}
}

func (m *Mutator) String() string {
	return m.pretty("")
}

func (m *Mutator) pretty(prefix string) string {
	out := fmt.Sprintf("name: %s index: %s Kind: %v Value: %v", m.name, m.index, m.kind, m.value)
	if m.child != nil {
		return fmt.Sprintf("%s \n%s %s ", out, prefix, m.child.pretty(prefix+"\t"))
	}
	return out
}

func (m *Mutator) WithValue(value any) {
	if m.child == nil {
		m.value = value
	} else {
		m.child.WithValue(value)
	}
}

func (m *Mutator) addToBottom(child *Mutator) {
	if m.child == nil {
		m.child = child
	} else {
		m.child.addToBottom(child)
	}
}

func (m *Mutator) ToMap(content map[string]any) map[string]any {
	if content == nil {
		content = make(map[string]any)
	}
	if m.child == nil {
		if m.value != nil {
			content[m.name] = m.applyValue(content[m.name])
		}
		return content
	}
	mt := *m.child
	switch m.kind {
	case Node:
		var childContent map[string]any
		c := content[m.name]
		if c == nil {
			childContent = make(map[string]any, 0)
		} else {
			childContent, _ = c.(map[string]any)
		}
		mt.ToMap(childContent)
		content[m.name] = childContent
	case Array:
		var childContent []any
		c := content[m.name]
		if c == nil {
			childContent = make([]any, 0)
		} else {
			childContent, _ = c.([]any)
		}
		content[m.name] = mt.ToArray(childContent)
	}

	return content
}

func (m *Mutator) ToArray(content []any) []any {
	if content == nil {
		content = make([]any, 0)
	}
	content = ensureSizeOfArray(content, m.index)

	index, err := strconv.Atoi(m.index)
	if err != nil {
		// This means that we should apply a '*'
	}
	if m.child == nil {
		if m.value != nil {
			content[index] = m.applyValue(content[index])
		}
		return content
	}
	c := content[index]
	switch m.kind {
	case Array:
		var childContent []any
		if c == nil {
			childContent = make([]any, 0)
		} else {
			childContent, _ = c.([]any)
		}
		content[index] = m.child.ToArray(childContent)
	case Node:
		var childContent map[string]any
		if c == nil {
			childContent = make(map[string]any)
		} else {
			childContent, _ = c.(map[string]any)
		}

		content[index] = m.child.ToMap(childContent)
	}
	return content
}

func ensureSizeOfArray(arrayContent []any, indexStr string) []any {
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		return arrayContent
	}
	increment := index - len(arrayContent) + 1
	if increment > 0 {
		appendedArray := make([]any, increment)
		arrayContent = append(arrayContent, appendedArray...)
	}
	return arrayContent
}
