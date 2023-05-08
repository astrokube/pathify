package node

import (
	"testing"
)

func Test_normalizeAttributeName(t *testing.T) {
	type args struct {
		value string
		index string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "a basic attribute with no  array index",
			args: args{
				value: "firstname",
			},
			want: "firstname",
		},
		{
			name: "a basic attribute with  array index",
			args: args{
				value: "firstname",
				index: "4",
			},
			want: "4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalizeAttributeName(tt.args.value, tt.args.index); got != tt.want {
				t.Errorf("normalizeAttributeName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pathToNode(t *testing.T) {
	type args struct {
		exp   string
		value any
	}
	tests := []struct {
		name string
		args args
		want partial
	}{
		{
			name: "single attribute",
			args: args{
				exp:   "firstname",
				value: "Jane",
			},
			want: partial{
				attributes: attributes{
					path:  "firstname",
					name:  "firstname",
					value: "Jane",
					kind:  attribute,
				},
			},
		},
		{
			name: "A complex attribute (3 deep levels)",
			args: args{
				exp:   "person.job.role",
				value: "Technical writer",
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
							value: "Technical writer",
						},
					},
				},
			},
		},
		{
			name: "A complex with array at the end",
			args: args{
				exp:   "person.job.roles[0]",
				value: "Technical writer",
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
							path: "person.job.roles",
							name: "roles",
							kind: array,
						},
						child: &partial{
							attributes: attributes{
								path:  "person.job.roles[0]",
								name:  "0",
								value: "Technical writer",
							},
						},
					},
				},
			},
		},
		{
			name: "A complex with array at the beginning and at the end",
			args: args{
				exp:   "people[0].job.roles[0]",
				value: "Technical writer",
			},
			want: partial{
				attributes: attributes{
					path: "people",
					name: "people",
					kind: array,
				},
				child: &partial{
					attributes: attributes{
						path: "people[0]",
						name: "0",
					},
					child: &partial{
						attributes: attributes{
							path: "people[0].job",
							name: "job",
						},
						child: &partial{
							attributes: attributes{
								path: "people[0].job.roles",
								name: "roles",
								kind: array,
							},
							child: &partial{
								attributes: attributes{
									path:  "people[0].job.roles[0]",
									name:  "0",
									value: "Technical writer",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "It contains escaped character (.)",
			args: args{
				exp:   "annotation.\"aws/ingress.loadBalancer\"",
				value: true,
			},
			want: partial{
				attributes: attributes{
					path: "annotation",
					name: "annotation",
				},
				child: &partial{
					attributes: attributes{
						path:  "annotation.\"aws/ingress.loadBalancer\"",
						name:  "aws/ingress.loadBalancer",
						value: true,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathToNode(tt.args.exp, tt.args.value); !got.EqualTo(tt.want) {
				t.Errorf("\n got = %s\nwant = %s", got.String(), tt.want.String())
			}
		})
	}
}
