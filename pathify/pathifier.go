package pathify

import (
	"log"
	"reflect"
)

type Pathifier[S Type] interface {
	Set(pathValueList ...any) Pathifier[S]
	Out() S
}

type Type interface {
	map[string]any | []any
}

type pathifier[T Type] struct {
	mutators  []mutator
	sanitizer *sanitizer
	parser    *parser
	content   T
}

type PathifyOpt func(sanitizer *builder)

type builder struct {
	strictMode  bool
	attrNameFmt string
}

func New(options ...PathifyOpt) Pathifier[map[string]any] {
	content := make(map[string]any)
	return Load[map[string]any](content, options...)
}

func Load[T Type](content T, options ...PathifyOpt) Pathifier[T] {
	b := &builder{
		strictMode:  false,
		attrNameFmt: defAttributeNameFormat,
	}
	for _, opt := range options {
		opt(b)
	}

	p := &pathifier[T]{
		sanitizer: &sanitizer{
			strict: b.strictMode,
		},
		parser: &parser{
			strict: b.strictMode,
			regExp: regExpFromAttributeFormat(b.attrNameFmt),
		},
		content: content,
	}

	return p
}

func WithStrictMode(strict bool) func(builder *builder) {
	return func(builder *builder) {
		builder.strictMode = strict
	}
}

func WithAttributeNameFormat(attrNameFmt string) func(builder *builder) {
	return func(opts *builder) {
		opts.attrNameFmt = attrNameFmt
	}
}

func (p *pathifier[S]) Set(args ...any) Pathifier[S] {
	pathValueList := p.sanitizer.sanitize(args...)
	for _, pathValue := range pathValueList {
		m := p.parser.parse(pathValue.path)
		m.withValue(pathValue.value)
		p.mutators = append(p.mutators, *m)
	}
	return p
}

func (p *pathifier[S]) Out() S {
	var content S = p.content
	for _, m := range p.mutators {
		switch reflect.TypeOf(content).Kind() {
		case reflect.Array, reflect.Slice:
			in := reflect.ValueOf(content).Interface().([]any)
			content = reflect.ValueOf(m.child.toArray(in)).Interface().(S)
		case reflect.Map:
			in := reflect.ValueOf(content).Interface().(map[string]any)
			content = reflect.ValueOf(m.child.toMap(in)).Interface().(S)
		default:
			log.Fatalf("unsupporteed output type '%s'", reflect.TypeOf(content).Kind())
		}
	}
	return content
}
