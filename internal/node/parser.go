package node

import (
	"fmt"
	"regexp"
)

const (
	attributeRegExpStr  = `([A-Za-z_]+[A-Za-z0-9_]*|\"[A-Za-z_]+[A-Za-z0-9_./-]*\")`
	arrayIndexRegExpStr = `([0-9]+|\*)`
)

var (
	pathRegExpStr = fmt.Sprintf(`^(?P<parent>(%s(\[%s\])*\.)*)(?P<attribute>%s)(\[(?P<index>%s)\])*$`, attributeRegExpStr, arrayIndexRegExpStr, attributeRegExpStr, arrayIndexRegExpStr)
	pathRegExp    = regexp.MustCompile(pathRegExpStr)
)

func pathToNode(exp string, value any) partial {
	match := pathRegExp.FindStringSubmatch(exp)
	subMatchMap := map[string]string{}
	for i, name := range pathRegExp.SubexpNames() {
		if i != 0 {
			subMatchMap[name] = match[i]
		}
	}
	parent := subMatchMap["parent"]
	attr := subMatchMap["attribute"]
	index := subMatchMap["index"]
	var path = attr
	if parent != "" {
		parent = parent[:len(parent)-1]
		path = fmt.Sprintf("%s.%s", parent, attr)
	}
	n := partial{
		attributes: attributes{
			value: value,
			path:  path,
			name:  normalizeAttributeName(attr, index),
		},
	}

	if index != "" {

		n.path = fmt.Sprintf("%s[%v]", n.path, index)
		var parentPath = attr
		if parent != "" {
			parentPath = parent + "." + parentPath
		}
		parentNode := pathToNode(parentPath, nil)
		parentNode.addToBottom(&n)
		parentNode.kind = array
		return parentNode
	}

	if len(parent) > 0 {

		parentNode := pathToNode(parent, nil)
		parentNode.addToBottom(&n)
		return parentNode
	}

	return n
}

func normalizeAttributeName(value, index string) string {
	if index != "" {
		return index
	}
	if value[0] == '"' && value[len(value)-1] == '"' {
		return value[1 : len(value)-1]
	}
	return value
}
