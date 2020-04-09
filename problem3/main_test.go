package main

import (
	"reflect"
	"testing"
)

func Test_find_out_prime(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test", args{13195}, []int{5, 7, 13, 29}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find_out_prime(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("find_out_prime() = %v, want %v", got, tt.want)
			}
		})
	}
}
