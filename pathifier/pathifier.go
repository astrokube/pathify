package pathifier

import (
	"github.com/astrokube/pathify/pathifier/internal"
	"log"
	"reflect"
)

type Pathifier[S Type] interface {
	Set(pathValueList ...any) Pathifier[S]
	Out() S
	String(opts ...internal.OutputOpt) string
}

type Type interface {
	map[string]any | []any
}

func checkValue(value any) any {
	switch reflect.ValueOf(value).Kind() {
	case reflect.Struct:
		return nil
	default:
		return value
	}
}

type pathifier[T Type] struct {
	mutators  []internal.Mutator
	sanitizer *internal.Sanitizer
	parser    *internal.Parser
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
		attrNameFmt: internal.DefAttributeNameFormat,
	}
	for _, opt := range options {
		opt(b)
	}

	p := &pathifier[T]{
		sanitizer: &internal.Sanitizer{
			Strict: b.strictMode,
		},
		parser: &internal.Parser{
			Strict: b.strictMode,
			RegExp: internal.RegExpFromAttributeFormat(b.attrNameFmt),
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
	pathValueList := p.sanitizer.SanitizePathValueList(args...)
	for _, pathValue := range pathValueList {
		v := checkValue(pathValue.Value)
		m := p.parser.Parse(pathValue.Path)
		m.WithValue(v)
		p.mutators = append(p.mutators, *m)
	}
	return p
}

func (p *pathifier[S]) Out() S {
	var content S = p.content
	for _, m := range p.mutators {
		switch reflect.TypeOf(content).Kind() {
		case reflect.Array, reflect.Slice:
			in, ok := reflect.ValueOf(content).Interface().([]any)
			if ok {
				content, _ = reflect.ValueOf(m.Child().ToArray(in)).Interface().(S)
			}
		case reflect.Map:
			in, ok := reflect.ValueOf(content).Interface().(map[string]any)
			if ok {
				content, _ = reflect.ValueOf(m.Child().ToMap(in)).Interface().(S)
			}
		default:
			log.Fatalf("unsupporteed output type '%s'", reflect.TypeOf(content).Kind())
		}
	}
	return content
}

func (p *pathifier[S]) String(opts ...internal.OutputOpt) string {
	content := p.Out()

	return internal.NewOutput(opts...).String(content)
}
