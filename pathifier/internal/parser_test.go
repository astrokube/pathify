package internal

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parser_parse(t *testing.T) {
	type fields struct {
		regExp *regexp.Regexp
		strict bool
	}
	type args struct {
		pathExpr string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     *Mutator
		panicked bool
	}{
		{
			name: "two deep level valid expression Path",
			fields: fields{
				strict: false,
				regExp: regExpFromAttributeFormat(DefAttributeNameFormat),
			},
			args: args{
				pathExpr: "people[0].firstname",
			},
			want: &Mutator{
				name: "people",
				kind: array,
				child: &Mutator{
					index: "0",
					kind:  node,
					child: &Mutator{
						name: "firstname",
					},
				},
			},
		},
		{
			name: "An invalid expression but Strict mode is disabled",
			fields: fields{
				strict: false,
				regExp: regExpFromAttributeFormat(DefAttributeNameFormat),
			},
			args: args{
				pathExpr: "peopl\\\\e[0].firstname",
			},
			want: nil,
		},
		{
			name: "An invalid expression and Strict mode is enabled",
			fields: fields{
				strict: true,
				regExp: regExpFromAttributeFormat(DefAttributeNameFormat),
			},
			args: args{
				pathExpr: "peopl\\\\e[0].firstname",
			},
			want:     nil,
			panicked: true,
		},
		{
			name: "A simple array",
			fields: fields{
				strict: true,
				regExp: regExpFromAttributeFormat(DefAttributeNameFormat),
			},
			args: args{
				pathExpr: "[0]",
			},
			want: &Mutator{
				kind: array,
				child: &Mutator{
					index: "0",
				},
			},
		},
		{
			name: "Multiple arrays",
			fields: fields{
				strict: true,
				regExp: regExpFromAttributeFormat(DefAttributeNameFormat),
			},
			args: args{
				pathExpr: "[0][1][2].name",
			},
			want: &Mutator{
				kind: array,
				child: &Mutator{
					kind:  array,
					index: "0",
					child: &Mutator{
						kind:  array,
						index: "1",
						child: &Mutator{
							index: "2",
							kind:  array,
							child: &Mutator{
								name: "name",
							},
						},
					},
				},
			},
		},
		{
			name: "single array",
			fields: fields{
				strict: true,
				regExp: regExpFromAttributeFormat(DefAttributeNameFormat),
			},
			args: args{
				pathExpr: "[2]",
			},
			want: &Mutator{
				kind: array,
				child: &Mutator{
					index: "2",
				},
			},
		},
		{
			name: "Attributes contains dots ",
			fields: fields{
				strict: false,
				regExp: regExpFromAttributeFormat(DefAttributeNameFormat),
			},
			args: args{
				pathExpr: "annotations.\"a.b.c\"",
			},
			want: &Mutator{
				name: "annotations",
				kind: node,
				child: &Mutator{
					name: "a.b.c",
				},
			},
		},
		{
			name: "Attributes in the middle of a Path contains dots ",
			fields: fields{
				strict: false,
				regExp: regExpFromAttributeFormat(DefAttributeNameFormat),
			},
			args: args{
				pathExpr: "annotations.\"a.b.c\".Value[0].name",
			},
			want: &Mutator{
				name: "annotations",
				kind: node,
				child: &Mutator{
					name: "a.b.c",
					kind: node,
					child: &Mutator{
						name: "Value",
						kind: array,
						child: &Mutator{
							index: "0",
							kind:  node,
							child: &Mutator{
								name: "name",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				RegExp: tt.fields.regExp,
				Strict: tt.fields.strict,
			}
			println(tt.fields.regExp.String())
			println(tt.args.pathExpr)
			if tt.panicked {
				assert.Panics(t, func() { p.parse(tt.args.pathExpr) }, "The execution should end panicking")
			} else {
				res := p.parse(tt.args.pathExpr)
				assertParsedElements(t, tt.want, res)
			}
		})
	}
}

func assertParsedElements(t *testing.T, expected *Mutator, got *Mutator) {
	if expected == nil && got == nil {
		return
	}
	if (expected == nil && got != nil) || (expected != nil && got == nil) {
		t.Errorf("\nexpected= %s  , \ngot= %s", expected, got)
		return
	}
	if (expected.child == nil && got.child != nil) || (expected.child != nil && got.child == nil) {
		t.Errorf("\nexpected= %s  , \ngot= %s", expected, got)
		return
	}
	if got.child != nil {
		assertParsedElements(t, expected.child, got.child)
		return
	}
	if expected.kind != got.kind || expected.name != got.name || expected.value != got.value || expected.index != got.index {
		t.Errorf("\nexpected= %s  , got= %s", expected, got)
		return
	}
}

func Test_regExpFromAttributeFormat(t *testing.T) {
	type args struct {
		attributeFormat string
	}
	tests := []struct {
		name string
		args args
		want *regexp.Regexp
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, regExpFromAttributeFormat(tt.args.attributeFormat), "regExpFromAttributeFormat(%v)", tt.args.attributeFormat)
		})
	}
}
