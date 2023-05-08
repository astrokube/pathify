package node

import (
	"reflect"
	"testing"
)

func TestRoot_Add(t *testing.T) {
	type args struct {
		path  string
		value any
	}
	tests := []struct {
		name string
		r    Root
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.Add(tt.args.path, tt.args.value)
		})
	}
}

func TestRoot_AsMap(t *testing.T) {
	tests := []struct {
		name string
		r    Root
		want map[string]any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.AsMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoot_asArray(t *testing.T) {
	tests := []struct {
		name string
		r    Root
		want []any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.asArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("asArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoot_merge(t *testing.T) {
	type args struct {
		n partial
	}
	tests := []struct {
		name string
		r    Root
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.merge(tt.args.n)
		})
	}
}
