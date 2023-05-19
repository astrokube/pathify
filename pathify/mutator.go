package pathify

import "strconv"

type kind int32

const (
	node kind = iota
	array
)

type mutator struct {
	name  string
	index string
	child *mutator
	kind  kind
	value any
}

func (m *mutator) withValue(value any) {
	if m.child == nil {
		m.value = value
	} else {
		m.child.withValue(value)
	}
}

func (m *mutator) addToBottom(child *mutator) {
	if m.child == nil {
		m.child = child
	} else {
		m.child.addToBottom(child)
	}
}

func (m *mutator) toMap(content map[string]any) map[string]any {
	if content == nil {
		content = make(map[string]any)
	}
	if m.child == nil {
		content[m.name] = m.value
		return content
	}
	mt := *m.child
	switch mt.kind {
	case node:
		var childContent map[string]any
		c := content[m.name]
		if c == nil {
			childContent = make(map[string]any, 0)
		} else {
			childContent, _ = c.(map[string]any)
		}
		childContent[mt.name] = mt.toMap(childContent)
		content[m.name] = childContent[mt.name]
	case array:
		var childContent []any
		c := content[m.name]
		if c == nil {
			childContent = make([]any, 0)
		} else {
			childContent, _ = c.([]any)
		}
		content[m.name] = mt.toArray(childContent)
	}

	return content
}

func (m *mutator) toArray(content []any) []any {
	if content == nil {
		content = make([]any, 0)
	}
	content = ensureSizeOfArray(content, m.index)

	index, err := strconv.Atoi(m.index)
	if err != nil {
		// This means that we should apply a '*'
	}
	if m.child == nil {
		content[index] = m.value
		return content
	}
	c := content[index]
	switch m.kind {
	case array:
		var childContent []any
		if c == nil {
			childContent = make([]any, 0)
		} else {
			childContent, _ = c.([]any)
		}
		content[index] = m.child.toArray(childContent)
	case node:
		var childContent map[string]any
		if c == nil {
			childContent = make(map[string]any)
		} else {
			childContent, _ = c.(map[string]any)
		}

		content[index] = m.child.toMap(childContent)
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
