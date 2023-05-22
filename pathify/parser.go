package pathify

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

const (
	defAttributeNameFormat = `(\"[A-Za-z_]+[A-Za-z0-9_./-]*\"|[A-Za-z_]+[A-Za-z0-9_/-]*)`
	arrayIndexExprStr      = `([0-9]+|\*)`
)

type parser struct {
	regExp *regexp.Regexp
	strict bool
}

func regExpFromAttributeFormat(attributeFormat string) *regexp.Regexp {
	regExpStr := fmt.Sprintf(`^(?P<parent>((%s|\[%s\]))*)((\.)(?P<attribute>%s)|(\[(?P<index>%s)\]))$`, attributeFormat, arrayIndexExprStr, attributeFormat, arrayIndexExprStr)
	return regexp.MustCompile(regExpStr)
}

func (p *parser) parse(pathExpr string) *mutator {
	match := p.regExp.FindStringSubmatch(pathExpr)
	if match == nil {
		if p.strict {
			log.Panicf("invalid path  '%v'. Path doesn't meet defined format", pathExpr)
		} else {
			return nil
		}
	}
	subMatchMap := map[string]string{}
	for i, name := range p.regExp.SubexpNames() {
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
	m := &mutator{
		name: attr,
	}
	if arrayIndex != "" {
		m.index = arrayIndex
		var parent = &mutator{}
		if parentExpr != "" {
			parent = p.parse(parentExpr)
			if parent == nil {
				parent = &mutator{
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
		parent := p.parse(parentExpr)
		parent.addToBottom(m)
		return parent
	}
	return m
}
