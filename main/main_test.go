package main

import "testing"

func TestGreeter(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "proper greet message should be returned",
			args: args{
				msg: "welcome",
			},
			want: "Hello welcome",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := greeter(tt.args.msg); got != tt.want {
				t.Errorf("greeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
