package node

import "testing"

func Test_partial_EqualTo(t *testing.T) {
	type fields struct {
		attributes attributes
		child      *partial
	}
	type args struct {
		other partial
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "both elements with zero value attributes",
			fields: fields{},
			args: args{
				other: partial{},
			},
			want: true,
		},
		{
			name: "One element with echild nil and the other with an empty initialized child",
			fields: fields{
				child: nil,
			},
			args: args{
				other: partial{
					child: &partial{},
				},
			},
			want: false,
		},
		{
			name: "All the attributes are filled for both and their values are the same",
			fields: fields{
				attributes: attributes{
					path:  "person.firstname",
					name:  "firstname",
					value: "John",
					kind:  attribute,
				},
			},
			args: args{
				other: partial{
					attributes: attributes{
						path:  "person.firstname",
						name:  "firstname",
						value: "John",
						kind:  attribute,
					},
				},
			},
			want: true,
		},
		{
			name: "All the attributes are filled for both and their values are the same but the value in one of them is in uppercase",
			fields: fields{
				attributes: attributes{
					path:  "person.firstname",
					name:  "firstname",
					value: "John",
					kind:  attribute,
				},
			},
			args: args{
				other: partial{
					attributes: attributes{
						path:  "person.firstname",
						name:  "firstname",
						value: "JOHN",
						kind:  attribute,
					},
				},
			},
			want: false,
		},
		{
			name: "Both objects are identical and they both  with more than one deep level of attributes",
			fields: fields{
				attributes: attributes{
					path: "person",
					name: "person",
				},
				child: &partial{
					attributes: attributes{
						path:  "person.firstname",
						name:  "firstname",
						value: "John",
						kind:  attribute,
					},
				},
			},
			args: args{
				other: partial{
					attributes: attributes{
						path: "person",
						name: "person",
					},
					child: &partial{
						attributes: attributes{
							path:  "person.firstname",
							name:  "firstname",
							value: "John",
							kind:  attribute,
						},
					},
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &partial{
				attributes: tt.fields.attributes,
				child:      tt.fields.child,
			}
			if got := n.EqualTo(tt.args.other); got != tt.want {
				t.Errorf("EqualTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partial_String(t *testing.T) {
	type fields struct {
		attributes attributes
		child      *partial
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Node with values for all the attributes except the child",
			fields: fields{
				attributes: attributes{
					path:  "person.firstname",
					name:  "firstname",
					value: "John",
					kind:  attribute,
				},
			},
			want: "Path: person.firstname, Name: firstname, Value:John",
		},
		{
			name: "Node with no value dor the attribute 'value'",
			fields: fields{
				attributes: attributes{
					path: "person.firstname",
					name: "firstname",
				},
			},
			want: "Path: person.firstname, Name: firstname",
		},
		{
			name: "Two deep level of nodes",
			fields: fields{
				attributes: attributes{
					path: "person",
					name: "person",
				},
				child: &partial{
					attributes: attributes{
						path:  "person.firstname",
						name:  "firstname",
						value: "John",
						kind:  attribute,
					},
				},
			},
			want: "Path: person, Name: person, Child:{ Path: person.firstname, Name: firstname, Value:John }",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &partial{
				attributes: tt.fields.attributes,
				child:      tt.fields.child,
			}
			if got := n.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partial_addToBottom(t *testing.T) {
	type fields struct {
		attributes attributes
		child      *partial
	}
	type args struct {
		child *partial
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   partial
	}{
		{
			name: "The child is nil",
			fields: fields{
				attributes: attributes{
					path: "person",
					name: "person",
				},
			},
			args: args{
				child: &partial{
					attributes: attributes{
						path:  "person.firstname",
						name:  "firstname",
						value: "John",
						kind:  attribute,
					},
				},
			},
			want: partial{
				attributes: attributes{
					path: "person",
					name: "person",
				},
				child: &partial{
					attributes: attributes{
						path:  "person.firstname",
						name:  "firstname",
						value: "John",
						kind:  attribute,
					},
				},
			},
		},

		{
			name: "The child is nil",
			fields: fields{
				attributes: attributes{
					path: "person",
					name: "person",
				},
				child: &partial{
					attributes: attributes{
						path: "person.job",
						name: "job",
					},
				},
			},
			args: args{
				child: &partial{
					attributes: attributes{
						path:  "person.job.role",
						name:  "role",
						value: "Software Developer",
					},
				},
			},
			want: partial{
				attributes: attributes{
					path: "person",
					name: "person",
				},
				child: &partial{
					attributes: attributes{
						path: "person.job",
						name: "job",
					},
					child: &partial{
						attributes: attributes{
							path:  "person.job.role",
							name:  "role",
							value: "Software Developer",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &partial{
				attributes: tt.fields.attributes,
				child:      tt.fields.child,
			}
			n.addToBottom(tt.args.child)
			if !n.EqualTo(tt.want) {
				t.Errorf("EqualTo() = %v, want %v", n, tt.want)
			}
		})
	}
}
