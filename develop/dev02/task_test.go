package main

import "testing"

func TestUnpack(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "error",
			args:    args{str: "45"},
			want:    "45",
			wantErr: true,
		},
		{
			name:    "standard string",
			args:    args{str: "abcd"},
			want:    "abcd",
			wantErr: false,
		},
		{
			name:    "string with nums",
			args:    args{str: "a4bc2d5e"},
			want:    "aaaabccddddde",
			wantErr: false,
		},
		{
			name:    "empty string",
			args:    args{str: ""},
			want:    "",
			wantErr: false,
		},
		{
			name:    `string with \ №1`,
			args:    args{str: `qwe\4\5`},
			want:    "qwe45",
			wantErr: false,
		},
		{
			name:    `string with \ №2`,
			args:    args{str: `qwe\45`},
			want:    "qwe44444",
			wantErr: false,
		},
		{
			name:    `string with \ №3`,
			args:    args{str: `qwe\\5`},
			want:    `qwe\\\\\`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := Unpack(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unpack() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if ans != tt.want {
				t.Errorf("Unpack() answer = [%v], want = [%v]", ans, tt.want)
				return
			}
		})
	}
}
