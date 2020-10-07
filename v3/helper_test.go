package main

import "testing"

func Test_fnv32hash(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_alex",
			args: args{str: "alex"},
			want: 1705310533,
		},
		{
			name: "test_chi",
			args: args{str: "chi"},
			want: 1971685989,
		},
		{
			name: "test_cai",
			args: args{str: "cai"},
			want: 1769221728,
		},
		{
			name: "test_ruo",
			args: args{str: "ruo"},
			want: 734875741,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fnv32hash(tt.args.str); got != tt.want {
				t.Errorf("fnv32hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isLetter(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_number",
			args: args{b: '8'},
			want: false,
		},
		{
			name: "test_symbol",
			args: args{b: '('},
			want: false,
		},
		{
			name: "test_lowercase",
			args: args{b: 'v'},
			want: true,
		},
		{
			name: "test_uppercase",
			args: args{b: 'R'},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLetter(tt.args.b); got != tt.want {
				t.Errorf("isLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
