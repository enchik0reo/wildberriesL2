package main

import "testing"

func TestAnagrams(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want map[string][]string
	}{
		{
			name: "normal",
			args: args{words: []string{"слиток", "пятак", "столик", "пятка", "тяпка", "листок", "зебра"}},
			want: map[string][]string{"пятак": {"пятак", "пятка", "тяпка"}, "слиток": {"листок", "слиток", "столик"}},
		},
		{
			name: "no anagram",
			args: args{words: []string{"слиток", "пятак", "зебра"}},
			want: map[string][]string{},
		},
		{
			name: "one element",
			args: args{words: []string{"слиток"}},
			want: map[string][]string{},
		},
		{
			name: "empty slice",
			args: args{words: []string{}},
			want: map[string][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Anagrams(tt.args.words)
			for key, anagram := range res {
				for i := range anagram {
					if anagram[i] != tt.want[key][i] {
						t.Errorf("Unpack() answer = [%v], want = [%v]", res, tt.want)
						return
					}
				}
			}
		})
	}
}
