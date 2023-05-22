package pathify

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ensureSizeOfArray(t *testing.T) {
	type args struct {
		arrayContent []any
		indexStr     string
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			name: "The size of the array mustn't be changed",
			args: args{
				arrayContent: []any{10, 20, 30},
				indexStr:     "2",
			},
			want: []any{10, 20, 30},
		},
		{
			name: "The size of the array must be changed",
			args: args{
				arrayContent: []any{10, 20, 30},
				indexStr:     "7",
			},
			want: []any{10, 20, 30, nil, nil, nil, nil, nil},
		},
		{
			name: "The index is invalid",
			args: args{
				arrayContent: []any{10, 20, 30},
				indexStr:     "A",
			},
			want: []any{10, 20, 30},
		},
		{
			name: "The index is a non-positive value",
			args: args{
				arrayContent: []any{10, 20, 30},
				indexStr:     "-A",
			},
			want: []any{10, 20, 30},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ensureSizeOfArray(tt.args.arrayContent, tt.args.indexStr), "ensureSizeOfArray(%v, %v)", tt.args.arrayContent, tt.args.indexStr)
		})
	}
}

func Test_mutator_addToBottom(t *testing.T) {
	type fields struct {
		name  string
		index string
		child *mutator
		kind  kind
		value any
	}
	type args struct {
		child *mutator
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected *mutator
	}{
		{
			name: "The root doesn't have a  child",
			fields: fields{
				name:  "root",
				child: nil,
				value: 20,
			},
			args: args{
				child: &mutator{
					name:  "child",
					value: 21,
				},
			},
			expected: &mutator{
				name: "root",
				child: &mutator{
					name:  "child",
					value: 21,
				},
				value: 20,
			},
		},
		{
			name: "The root has a child",
			fields: fields{
				name: "root",
				child: &mutator{
					name:  "child",
					value: 21,
				},
				value: 20,
			},
			args: args{
				child: &mutator{
					name:  "child",
					value: 22,
				},
			},
			expected: &mutator{
				name: "root",
				child: &mutator{
					name:  "child",
					value: 21,
					child: &mutator{
						name:  "child",
						value: 22,
					},
				},
				value: 20,
			},
		},
		{
			name: "The root has two levels child",
			fields: fields{
				name: "root",
				child: &mutator{
					name:  "child",
					value: 21,
					child: &mutator{
						name:  "child",
						value: 22,
					},
				},
				value: 20,
			},
			args: args{
				child: &mutator{
					name:  "child",
					value: 23,
				},
			},
			expected: &mutator{
				name: "root",
				child: &mutator{
					name:  "child",
					value: 21,
					child: &mutator{
						name:  "child",
						value: 22,
						child: &mutator{
							name:  "child",
							value: 23,
						},
					},
				},
				value: 20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mutator{
				name:  tt.fields.name,
				index: tt.fields.index,
				child: tt.fields.child,
				kind:  tt.fields.kind,
				value: tt.fields.value,
			}
			m.addToBottom(tt.args.child)
			a := m
			b := tt.expected
			for {
				if a == nil {
					if b != nil {
						assert.Errorf(t, errors.New("unexpected result"), "Expected %#v and obtained %#v", tt.expected, m)
					}
					return
				}
				if b == nil {
					if a != nil {
						assert.Errorf(t, errors.New("unexpected result"), "Expected %#v and obtained %#v", tt.expected, m)
					}
					return
				}
				assert.Equal(t, a.value, b.value)
				a = a.child
				b = b.child
			}
		})
	}
}

func Test_mutator_toArray(t *testing.T) {
	type fields struct {
		name  string
		index string
		child *mutator
		kind  kind
		value any
	}
	type args struct {
		content []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []any
	}{
		{
			name: "Add a new entry into the array",
			fields: fields{
				index: "1",
				child: &mutator{
					name:  "firstname",
					value: "Mary",
				},
				kind: node,
			},
			args: args{
				content: []any{
					map[string]any{
						"firstname": "Jane",
					},
				},
			},
			want: []any{
				map[string]any{
					"firstname": "Jane",
				},
				map[string]any{
					"firstname": "Mary",
				},
			},
		},
		{
			name: "Modify the value of an existing item in the array",
			fields: fields{
				index: "1",
				child: &mutator{
					name:  "firstname",
					value: "Mary",
				},
				kind: node,
			},
			args: args{
				content: []any{
					map[string]any{
						"firstname": "Jane",
					},
				},
			},
			want: []any{
				map[string]any{
					"firstname": "Jane",
				},
				map[string]any{
					"firstname": "Mary",
				},
			},
		},
		{
			name: "The initial array is empty",
			fields: fields{
				index: "1",
				child: &mutator{
					name:  "firstname",
					value: "Mary",
				},
				kind: node,
			},
			args: args{
				content: []any{},
			},
			want: []any{
				nil,
				map[string]any{
					"firstname": "Mary",
				},
			},
		},
		{
			name: "Add a new entry into the array that contains sub arrays",
			fields: fields{
				index: "1",
				child: &mutator{
					index: "3",
					value: "hello",
				},
				kind: array,
			},
			args: args{
				content: []any{
					[]any{
						"my friend",
					},
				},
			},
			want: []any{
				[]any{
					"my friend",
				},
				[]any{
					nil, nil, nil, "hello",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mutator{
				name:  tt.fields.name,
				index: tt.fields.index,
				child: tt.fields.child,
				kind:  tt.fields.kind,
				value: tt.fields.value,
			}

			assert.Equalf(t, tt.want, m.toArray(tt.args.content), "toArray(%v)", tt.args.content)
		})
	}
}

func Test_mutator_toMap(t *testing.T) {
	type fields struct {
		parentExpr string
		name       string
		index      string
		path       string
		child      *mutator
		kind       kind
		value      any
	}
	type args struct {
		content map[string]any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]any
	}{
		{
			name: "Add a new entry into the array",
			fields: fields{
				name: "item2",
				child: &mutator{
					name:  "firstname",
					value: "Mary",
				},
				kind: node,
			},
			args: args{
				content: map[string]any{
					"item1": map[string]any{
						"firstname": "Jane",
					},
				},
			},
			want: map[string]any{
				"item1": map[string]any{
					"firstname": "Jane",
				},
				"item2": map[string]any{
					"firstname": "Mary",
				},
			},
		},
		{
			name: "Modify the value of an existing item in the map",
			fields: fields{
				name: "item1",
				child: &mutator{
					name:  "firstname",
					value: "Mary",
				},
				kind: node,
			},
			args: args{
				content: map[string]any{
					"item1": map[string]any{
						"firstname": "Jane",
					},
				},
			},
			want: map[string]any{
				"item1": map[string]any{
					"firstname": "Mary",
				},
			},
		},
		{
			name: "The initial array is empty",
			fields: fields{
				name: "item",
				child: &mutator{
					name:  "firstname",
					value: "Mary",
				},
				kind: node,
			},
			args: args{
				content: map[string]any{},
			},
			want: map[string]any{
				"item": map[string]any{
					"firstname": "Mary",
				},
			},
		},
		{
			name: "Add a new entry into the map that contains  arrays",
			fields: fields{
				name: "item2",
				child: &mutator{
					index: "3",
					value: "hello",
				},
				kind: array,
			},
			args: args{
				content: map[string]any{
					"item1": []any{
						"my friend",
					},
				},
			},
			want: map[string]any{
				"item1": []any{
					"my friend",
				},
				"item2": []any{
					nil, nil, nil, "hello",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mutator{
				name:  tt.fields.name,
				index: tt.fields.index,
				child: tt.fields.child,
				kind:  tt.fields.kind,
				value: tt.fields.value,
			}
			assert.Equalf(t, tt.want, m.toMap(tt.args.content), "toMap(%v)", tt.args.content)
		})
	}
}

func Test_mutator_withValue(t *testing.T) {
	type fields struct {
		parentExpr string
		name       string
		index      string
		path       string
		child      *mutator
		kind       kind
		value      any
	}
	type args struct {
		value any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "AAdd a value to the mutator which sdoesn't contain any value yet",
			fields: fields{
				value: nil,
			},
			args: args{
				value: 20,
			},
		},
		{
			name: "AAdd a value to the mutator and overwrite its value",
			fields: fields{
				value: 21,
			},
			args: args{
				value: 20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mutator{
				name:  tt.fields.name,
				index: tt.fields.index,
				child: tt.fields.child,
				kind:  tt.fields.kind,
				value: tt.fields.value,
			}
			m.withValue(tt.args.value)
			assert.Equal(t, tt.args.value, m.value)
		})
	}
}
