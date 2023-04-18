package pathify

import (
	"regexp"
	"strings"
)

var fieldRegExp = regexp.MustCompile(`^[A-Za-z_]+[A-Za-z0-9_]*(\[[0-9]+\])*$`)

func newPath(value string) path {
	return strings.Split(value, ".")
}

type path []string

var emptyPath = path{}

func (p path) Parent() string {
	return p[0]
}

func (p path) children() path {
	if p.isRoot() {
		return emptyPath
	}
	return p[1:]
}

func (p path) isRoot() bool {
	return len(p) == 1
}

func (p path) validate() bool {
	isRootValid := matchFieldExpr(p.Parent())
	if p.isRoot() {
		return isRootValid
	}
	return isRootValid && p.children().validate()
}

func matchFieldExpr(value string) bool {
	return fieldRegExp.MatchString(value)
}
