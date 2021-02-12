package main

import (
	"reflect"
	"testing"
)

func Test_zero1(t *testing.T) {
	type args struct {
		a []int
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test (1,2)", args: args{
			a: []int{1,2},
			n: 1,
		}, want: []int{2,1}},
		{name: "test (1,3,5,7)", args: args{
			a: []int{1,3,5,7},
			n: 2,
		}, want:[]int{5,7,1,3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zero(tt.args.a, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zero() = %v, want %v", got, tt.want)
			}
		})
	}
}