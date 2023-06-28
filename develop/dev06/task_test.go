package main

import (
	"testing"
)

func TestFields(t *testing.T) {
	type args struct {
		strs []string
		f    int
		d    string
		s    bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "normal with f",
			args: args{strs: []string{"a a", "a/a", "a	a", "a:a"}, f: 2, d: "	"},
			want: []string{"a a", "a/a", "a", "a:a"},
		},
		{
			name: "normal with f, d",
			args: args{strs: []string{"a a", "a/a", "a	a", "a:a"}, f: 2, d: ":"},
			want: []string{"a a", "a/a", "a	a", "a"},
		},
		{
			name: "normal with f, d, s",
			args: args{strs: []string{"a a", "a/a", "a	a", "a:a"}, f: 2, d: "/", s: true},
			want: []string{"a"},
		},
		{
			name: "no matches with f, d, s",
			args: args{strs: []string{"a a", "a/a", "a	a", "a:a"}, f: 2, d: "::", s: true},
			want: []string{},
		},
		{
			name: "over f",
			args: args{strs: []string{"a a", "a/a", "a	a", "a:a"}, f: 3, d: "	", s: true},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Fields(tt.args.strs, tt.args.f, tt.args.d, tt.args.s)
			for i := range res {
				if res[i] != tt.want[i] {
					t.Errorf("Unpack() answer = [%v], want = [%v]", res, tt.want)
					return
				}
			}
		})
	}
}
