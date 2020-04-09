package main

import (
	"reflect"
	"testing"
)

func Test_digits_generator(t *testing.T) {
	tests := []struct {
		name string
		want <-chan int
	}{

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digitsGenerator(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("digits_generator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_forXDigits(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want <-chan int
	}{
		{"channel contains array of numbers",args{number: 0},}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := forXDigits(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("forXDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}