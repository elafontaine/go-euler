package main

import "testing"

func Test_fibonacci(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1 = 1", args{1}, 1},
		{"2 = 2", args{1}, 1},
		{"3 = 3", args{1}, 1},
		{"4 = 5", args{1}, 1},
		{"5 = 8", args{1}, 1},
		{"6 = 13", args{1}, 1},
		{"7 = 21", args{1}, 1},
		{"8 = 34", args{1}, 1},
		{"9 = 55", args{1}, 1},
		{"10 = 89", args{1}, 1},
		{"6 = 13", args{1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacci(tt.args.x); got != tt.want {
				t.Errorf("fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
