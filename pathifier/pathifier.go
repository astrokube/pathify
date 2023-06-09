package pathifier

import (
	"log"
	"reflect"

	"github.com/astrokube/pathify/pathifier/internal"
)

type Pathifier[S Type] interface {
	Set(pathValueList ...any) Pathifier[S]
	With(opts ...SetterOpt) func(pathValueList ...any) Pathifier[S]
	Out() S
	YAML() string
	JSON() string

	String(opts ...internal.OutputOpt) string
}

type Type interface {
	map[string]any | []any | any
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
	setter    *internal.Setter
	content   T
}

type PathifyOpt func(sanitizer *builder)

type builder struct {
	strictMode  bool
	attrNameFmt string
}

func New[S Type](options ...PathifyOpt) Pathifier[S] {
	var content S
	return Load[S](content, options...)
}

func Load[T Type](content T, options ...PathifyOpt) Pathifier[T] {
	b := &builder{
		strictMode:  false,
		attrNameFmt: internal.DefAttributeNameFormat,
	}
	for _, opt := range options {
		opt(b)
	}
	pathRegExp, attrRegExpr := internal.RegExpsFromAttributeFormat(b.attrNameFmt)
	p := &pathifier[T]{
		sanitizer: &internal.Sanitizer{
			Strict: b.strictMode,
		},
		parser: &internal.Parser{
			Strict:          b.strictMode,
			RegExp:          pathRegExp,
			AttributeRegExp: attrRegExpr,
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

func (p *pathifier[S]) With(opts ...SetterOpt) func(args ...any) Pathifier[S] {
	setter := internal.NewSetter(opts...)
	return func(args ...any) Pathifier[S] {
		pathValueList := p.sanitizer.SanitizePathValueList(args...)
		p.mutators = append(p.mutators, setter.Set(p.parser, pathValueList)...)
		return p
	}
}

func (p *pathifier[S]) Set(args ...any) Pathifier[S] {
	pathValueList := p.sanitizer.SanitizePathValueList(args...)
	p.mutators = append(p.mutators, internal.NewSetter().Set(p.parser, pathValueList)...)
	return p
}

func (p *pathifier[S]) Out() S {
	var content S = p.content
	for _, m := range p.mutators {
		switch reflect.ValueOf(content).Kind() {
		case reflect.Slice, reflect.Array:
			/**
			if len(content) == 0 {
				in := make([]any, 0)
				content, _ = reflect.ValueOf(m.Child().ToArray(in)).Interface().(S)
				continue
			}**/
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

func (p *pathifier[S]) YAML() string {
	return internal.NewOutput().YAML(p.Out())
}

func (p *pathifier[S]) JSON() string {
	return internal.NewOutput().JSON(p.Out())
}

type setter[S Type] struct {
	s         *internal.Setter
	pathifier *pathifier[S]
}

func (s *setter[S]) Set(args ...any) Pathifier[S] {
	pathValueList := s.pathifier.sanitizer.SanitizePathValueList(args...)
	s.pathifier.mutators = append(s.pathifier.mutators, internal.NewSetter().Set(s.pathifier.parser, pathValueList)...)
	return s.pathifier
}
