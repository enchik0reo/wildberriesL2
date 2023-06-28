package main

import (
	"testing"
)

func TestDefaultSort(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "normal slice",
			args: args{slice: []string{"abc", "123", "def", "456"}},
			want: []string{"123", "456", "abc", "def"},
		},
		{
			name: "one element slice",
			args: args{slice: []string{"a"}},
			want: []string{"a"},
		},
		{
			name: "the same elements slise",
			args: args{slice: []string{"a", "a"}},
			want: []string{"a", "a"},
		},
		{
			name: "empty slice",
			args: args{slice: []string{}},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			answer := DefaultSort(tt.args.slice)
			for i := range answer {
				if answer[i] != tt.want[i] {
					t.Errorf("Unpack() answer = [%v], want = [%v]", answer, tt.want)
					return
				}
			}
		})
	}
}

func TestColumnSort(t *testing.T) {
	type args struct {
		slice []string
		k     int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "normal slice",
			args: args{slice: []string{"a b", "b a", "b ac"}, k: 2},
			want: []string{"b a", "b ac", "a b"},
		},
		{
			name: "normal slice with k = 1",
			args: args{slice: []string{"a b", "b a", "b ac"}, k: 1},
			want: []string{"a b", "b a", "b ac"},
		},
		{
			name: "normal slice with over range k",
			args: args{slice: []string{"b ac", "a b", "b a"}, k: 3},
			want: []string{"a b", "b a", "b ac"},
		},
		{
			name: "one element slice",
			args: args{slice: []string{"a a"}, k: 2},
			want: []string{"a a"},
		},
		{
			name: "the same elements slise",
			args: args{slice: []string{"a a", "a a"}, k: 2},
			want: []string{"a a", "a a"},
		},
		{
			name: "without two columns slice",
			args: args{slice: []string{"abc", "123", "def", "456"}, k: 2},
			want: []string{"123", "456", "abc", "def"},
		},
		{
			name: "empty slice",
			args: args{slice: []string{}, k: 2},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			answer := ColumnSort(tt.args.slice, tt.args.k)
			for i := range answer {
				if answer[i] != tt.want[i] {
					t.Errorf("Unpack() answer = [%v], want = [%v]", answer, tt.want)
					return
				}
			}
		})
	}
}

func TestNumSort(t *testing.T) {
	type args struct {
		slice []string
		k     int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "normal slice",
			args: args{slice: []string{"abc", "123", "def", "456"}},
			want: []string{"abc", "def", "123", "456"},
		},
		{
			name: "normal slice with k",
			args: args{slice: []string{"abc 10axc", "123 qww", "def 7sa", "456"}, k: 2},
			want: []string{"123 qww", "456", "def 7sa", "abc 10axc"},
		},
		{
			name: "one element slice",
			args: args{slice: []string{"123"}},
			want: []string{"123"},
		},
		{
			name: "the same elements slise",
			args: args{slice: []string{"123", "123"}},
			want: []string{"123", "123"},
		},
		{
			name: "without nums slice",
			args: args{slice: []string{"def", "abc"}},
			want: []string{"abc", "def"},
		},
		{
			name: "empty slice",
			args: args{slice: []string{}},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			answer := NumSort(tt.args.slice, tt.args.k)
			for i := range answer {
				if answer[i] != tt.want[i] {
					t.Errorf("Unpack() answer = [%v], want = [%v]", answer, tt.want)
					return
				}
			}
		})
	}
}

func TestReverseSort(t *testing.T) {
	type args struct {
		slice []string
		k     int
		n     bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "normal slice",
			args: args{slice: []string{"abc", "123", "def", "456"}},
			want: []string{"def", "abc", "456", "123"},
		},
		{
			name: "normal slice with n",
			args: args{slice: []string{"abc", "123", "def", "456"}, n: true},
			want: []string{"456", "def", "123", "abc"},
		},
		{
			name: "normal slice with k",
			args: args{slice: []string{"abc", "123", "def", "456"}, k: 2},
			want: []string{"456", "def", "123", "abc"},
		},
		{
			name: "one element slice",
			args: args{slice: []string{"123"}},
			want: []string{"123"},
		},
		{
			name: "the same elements slise",
			args: args{slice: []string{"123", "123"}},
			want: []string{"123", "123"},
		},
		{
			name: "empty slice",
			args: args{slice: []string{}},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			answer := ReverseSort(tt.args.slice, tt.args.k, tt.args.n)
			for i := range answer {
				if answer[i] != tt.want[i] {
					t.Errorf("Unpack() answer = [%v], want = [%v]", answer, tt.want)
					return
				}
			}
		})
	}
}

func TestUniqSort(t *testing.T) {
	type args struct {
		slice []string
		k     int
		n     bool
		r     bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "normal slice",
			args: args{slice: []string{"abc", "123", "def", "456"}},
			want: []string{"123", "456", "abc", "def"},
		},
		{
			name: "normal slice with dubles",
			args: args{slice: []string{"abc", "123", "abc", "def", "456", "123"}},
			want: []string{"123", "456", "abc", "def"},
		},
		{
			name: "normal slice with dubles r",
			args: args{slice: []string{"abc", "123", "abc", "def", "456", "123"}, r: true},
			want: []string{"abc", "123", "def", "456"},
		},
		{
			name: "normal slice with dubles n",
			args: args{slice: []string{"abc", "123", "abc", "def", "456", "123"}, n: true},
			want: []string{"abc", "123", "def", "456"},
		},
		{
			name: "normal slice with dubles k",
			args: args{slice: []string{"abc", "123", "abc", "def", "456", "123"}, k: 2},
			want: []string{"abc", "123", "def", "456"},
		},
		{
			name: "one element slice",
			args: args{slice: []string{"123"}},
			want: []string{"123"},
		},
		{
			name: "empty slice",
			args: args{slice: []string{}},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			answer := UniqSort(tt.args.slice, tt.args.k, tt.args.n, tt.args.r)
			for i := range answer {
				if answer[i] != tt.want[i] {
					t.Errorf("Unpack() answer = [%v], want = [%v]", answer, tt.want)
					return
				}
			}
		})
	}
}
