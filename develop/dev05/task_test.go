package main

import "testing"

func TestAfter(t *testing.T) {
	type args struct {
		search string
		strs   []string
		a      int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "normal",
			args: args{search: "ap", strs: []string{"apple", "apricot", "banana", "lemon"}, a: 1},
			want: []string{"apple", "apricot", "banana"},
		},
		{
			name: "a out of range",
			args: args{search: "ap", strs: []string{"apple", "apricot", "banana", "lemon"}, a: 4},
			want: []string{"apple", "apricot", "banana", "lemon"},
		},
		{
			name: "no one want",
			args: args{search: "coc", strs: []string{"apple", "apricot", "banana", "lemon"}, a: 1},
			want: []string{},
		},
		{
			name: "a is 0",
			args: args{search: "ban", strs: []string{"apple", "apricot", "banana", "lemon"}, a: 0},
			want: []string{"banana"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := After(tt.args.search, tt.args.strs, tt.args.a)
			for i := range res {
				if res[i] != tt.want[i] {
					t.Errorf("Unpack() answer = [%v], want = [%v]", res, tt.want)
					return
				}
			}
		})
	}
}

func TestBefore(t *testing.T) {
	type args struct {
		search string
		strs   []string
		b      int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "normal",
			args: args{search: "ban", strs: []string{"apple", "apricot", "banana", "lemon"}, b: 1},
			want: []string{"apricot", "banana"},
		},
		{
			name: "b out of range",
			args: args{search: "ban", strs: []string{"apple", "apricot", "banana", "lemon"}, b: 4},
			want: []string{"apple", "apricot", "banana"},
		},
		{
			name: "no one want",
			args: args{search: "coc", strs: []string{"apple", "apricot", "banana", "lemon"}, b: 1},
			want: []string{},
		},
		{
			name: "b is 0",
			args: args{search: "ban", strs: []string{"apple", "apricot", "banana", "lemon"}, b: 0},
			want: []string{"banana"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Before(tt.args.search, tt.args.strs, tt.args.b)
			for i := range res {
				if res[i] != tt.want[i] {
					t.Errorf("Unpack() answer = [%v], want = [%v]", res, tt.want)
					return
				}
			}
		})
	}
}

func TestIgnoreCase(t *testing.T) {
	type args struct {
		search string
		strs   []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "normal",
			args: args{search: "Ap", strs: []string{"apple", "apricot", "banana", "lemon"}},
			want: []string{"apple", "apricot"},
		},
		{
			name: "no one want",
			args: args{search: "coc", strs: []string{"apple", "apricot", "banana", "lemon"}},
			want: []string{},
		},
		{
			name: "empty",
			args: args{search: "coc", strs: []string{}},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := IgnoreCase(tt.args.search, tt.args.strs)
			for i := range res {
				if res[i] != tt.want[i] {
					t.Errorf("Unpack() answer = [%v], want = [%v]", res, tt.want)
					return
				}
			}
		})
	}
}

func TestInvert(t *testing.T) {
	type args struct {
		search string
		strs   []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "normal",
			args: args{search: "ap", strs: []string{"apple", "apricot", "banana", "lemon"}},
			want: []string{"banana", "lemon"},
		},
		{
			name: "no one want",
			args: args{search: "ap", strs: []string{"apple", "apricot"}},
			want: []string{},
		},
		{
			name: "empty",
			args: args{search: "coc", strs: []string{}},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Invert(tt.args.search, tt.args.strs)
			for i := range res {
				if res[i] != tt.want[i] {
					t.Errorf("Unpack() answer = [%v], want = [%v]", res, tt.want)
					return
				}
			}
		})
	}
}

func TestFixed(t *testing.T) {
	type args struct {
		search string
		strs   []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "normal",
			args: args{search: "apple", strs: []string{"apple", "apricot", "banana", "lemon"}},
			want: []string{"apple"},
		},
		{
			name: "no one want",
			args: args{search: "ap", strs: []string{"apple", "apricot", "banana", "lemon"}},
			want: []string{},
		},
		{
			name: "empty",
			args: args{search: "coc", strs: []string{}},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Fixed(tt.args.search, tt.args.strs)
			for i := range res {
				if res[i] != tt.want[i] {
					t.Errorf("Unpack() answer = [%v], want = [%v]", res, tt.want)
					return
				}
			}
		})
	}
}

func TestLineNum(t *testing.T) {
	type args struct {
		search string
		strs   []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "normal",
			args: args{search: "ap", strs: []string{"apple", "apricot", "banana", "lemon"}},
			want: []string{"\u001b[32m1\u001b[0m\u001b[36m:\u001b[0mapple", "\u001b[32m2\u001b[0m\u001b[36m:\u001b[0mapricot"},
		},
		{
			name: "no one want",
			args: args{search: "coc", strs: []string{"apple", "apricot", "banana", "lemon"}},
			want: []string{},
		},
		{
			name: "empty",
			args: args{search: "coc", strs: []string{}},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := LineNum(tt.args.search, tt.args.strs)
			for i := range res {
				if res[i] != tt.want[i] {
					t.Errorf("Unpack() answer = [%v], want = [%v]", res, tt.want)
					return
				}
			}
		})
	}
}

func TestDef(t *testing.T) {
	type args struct {
		search string
		strs   []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "normal",
			args: args{search: "a", strs: []string{"apple", "apricot", "banana", "lemon"}},
			want: []string{"apple", "apricot", "banana"},
		},
		{
			name: "no one want",
			args: args{search: "coc", strs: []string{"apple", "apricot", "banana", "lemon"}},
			want: []string{},
		},
		{
			name: "empty",
			args: args{search: "ap", strs: []string{}},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Def(tt.args.search, tt.args.strs)
			for i := range res {
				if res[i] != tt.want[i] {
					t.Errorf("Unpack() answer = [%v], want = [%v]", res, tt.want)
					return
				}
			}
		})
	}
}
