package internal

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

const (
	DefAttributeNameFormat = `("[A-Za-z_]+[A-Za-z0-9_./-]*"|[A-Za-z_]+[A-Za-z0-9_/-]*)`
	arrayIndexExprStr      = `([0-9]+|\*)`
)

type Parser struct {
	RegExp *regexp.Regexp
	Strict bool
}

func RegExpFromAttributeFormat(attributeFormat string) *regexp.Regexp {
	regExpStr := fmt.Sprintf(`^(?P<parent>(((\.)?%s|\[%s\]))*)((\.)(?P<attribute>%s)|(\[(?P<index>%s)\]))$`,
		attributeFormat, arrayIndexExprStr, attributeFormat, arrayIndexExprStr)
	return regexp.MustCompile(regExpStr)
}

func (p *Parser) Parse(pathExpr string) *Mutator {
	match := p.RegExp.FindStringSubmatch(pathExpr)
	if match == nil {
		if p.Strict {
			log.Panicf("invalid Path  '%v'. Path doesn't meet defined format", pathExpr)
		} else {
			return nil
		}
	}
	subMatchMap := map[string]string{}
	for i, name := range p.RegExp.SubexpNames() {
		if i != 0 {
			subMatchMap[name] = match[i]
		}
	}
	attr := subMatchMap["attribute"]
	if strings.HasPrefix(attr, ".") {
		attr = attr[1:]
	}
	parentExpr := subMatchMap["parent"]
	if strings.HasSuffix(parentExpr, ".") {
		parentExpr = parentExpr[:len(parentExpr)-1]
	}
	arrayIndex := subMatchMap["index"]
	m := &Mutator{
		name: attr,
	}
	if arrayIndex != "" {
		m.index = arrayIndex
		parent := &Mutator{}
		if parentExpr != "" {
			parent = p.Parse(parentExpr)
			if parent == nil {
				parent = &Mutator{
					name: parentExpr,
				}
			}
		}
		parent.kind = array
		parent.addToBottom(m)
		return parent
	}
	if parentExpr != "" {
		if attr != "" {
			if attr[0] == '"' && attr[len(attr)-1] == '"' {
				m.name = attr[1 : len(attr)-1]
			}
		}
		parent := p.Parse(parentExpr)
		if parent == nil {
			parent = &Mutator{
				name: parentExpr,
			}
		}
		parent.addToBottom(m)
		return parent
	}
	return m
}
