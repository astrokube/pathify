package internal

import (
	"fmt"
	"regexp"
)

const (
	attribute kind = iota
	array
)

var pathRegExp = regexp.MustCompile(`^(?P<parent>([A-Za-z_]+[A-Za-z0-9_]*(\[([0-9]+|\*)\])*\.)*)(?P<attribute>[A-Za-z_]+[A-Za-z0-9_]*)(\[(?P<index>([0-9]+|\*))\])*$`)

type kind int

type node struct {
	path     string
	name     string
	value    any
	kind     kind
	children map[string]node
}

func parsePath(path string, value any) []node {
	match := pathRegExp.FindStringSubmatch(path)
	subMatchMap := map[string]string{}
	for i, name := range pathRegExp.SubexpNames() {
		if i != 0 {
			subMatchMap[name] = match[i]
		}
	}
	parent := subMatchMap["parent"]
	attr := subMatchMap["attribute"]
	index := subMatchMap["index"]
	var out = []node{{
		path: fmt.Sprintf("%s.%s", parent, attribute),
		name: attr,
		kind: attribute,
	}}
	if index != "" {
		out = append(out, node{
			path: fmt.Sprintf("%s.%s[%s}", parent, attribute, index),
			name: index,
			kind: attribute,
		})
		out[0].kind = array
	}
	out[len(out)-1].value = value

	if len(parent) > 0 {
		return append(parsePath(parent[:len(parent)-1], nil), out...)
	}
	return out
}
