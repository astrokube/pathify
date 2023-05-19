package pathify

import (
	"fmt"
	"log"
	"regexp"
)

const (
	defAttributeNameFormat = `[A-Za-z_]+[A-Za-z0-9_./-]*`
	arrayIndexExprStr      = `([0-9]+|\*)`
)

type parser struct {
	regExp *regexp.Regexp
	strict bool
}

func regExpFromAttributeFormat(attributeFormat string) *regexp.Regexp {
	attributeRegExpStr := fmt.Sprintf(`((%s|\"%s\"))`, attributeFormat, attributeFormat)
	regExpStr := fmt.Sprintf(`^(?P<parent>((%s(\[%s\])*|\[%s\])\.)*)((?P<attribute>%s)|(\[(?P<index>%s)\]))*$`,
		attributeRegExpStr, arrayIndexExprStr, arrayIndexExprStr, attributeRegExpStr, arrayIndexExprStr)
	return regexp.MustCompile(regExpStr)
}

func (p *parser) parse(pathExpr string) *mutator {
	match := p.regExp.FindStringSubmatch(pathExpr)
	if match == nil && p.strict {
		log.Fatalf("invalid path  '%v'. Path doesn't meet defined format", pathExpr)
	}
	subMatchMap := map[string]string{}
	for i, name := range p.regExp.SubexpNames() {
		if i != 0 {
			subMatchMap[name] = match[i]
		}
	}
	attr := subMatchMap["attribute"]
	parentExpr := subMatchMap["parent"]
	arrayIndex := subMatchMap["index"]
	m := &mutator{
		name: attr,
	}
	if arrayIndex != "" {

		m.child = &mutator{
			index: arrayIndex,
		}
		if parentExpr == "" {
			m.kind = array
			return m
		}
		parent := p.parse(parentExpr)
		parent.kind = array
		parent.child = m
		return parent
	}
	if parentExpr != "" {
		if attr != "" {
			if attr[0] == '"' && attr[len(attr)-1] == '"' {
				m.name = attr[1 : len(attr)-1]
			}
		}
		parent := p.parse(parentExpr[:len(parentExpr)-1])
		// parent.kind = node
		parent.addToBottom(m)
		return parent
	}

	return m
}
