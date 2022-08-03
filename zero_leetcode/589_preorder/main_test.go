package main

import (
	"math"
	"reflect"
	"testing"
)

func Test_preorder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		// TODO: Add test cases.
		{
			name:    "test01",
			args:    args{root: nil},
			wantAns: []int{},
		},
		{
			name:    "test02",
			args:    args{root: NewTree([]int{1, math.MinInt, 3, 2, 4, math.MinInt, 5, 6})},
			wantAns: []int{1, 3, 2, 4, 5, 6},
		},
		{
			name:    "test03",
			args:    args{root: NewTree([]int{1, math.MinInt, 2, 3, 4, 5, math.MinInt, math.MinInt, 6, 7, math.MinInt, 8, math.MinInt, 9, 10, math.MinInt, math.MinInt, 11, math.MinInt, 12, math.MinInt, 13, math.MinInt, math.MinInt, 14})},
			wantAns: []int{1, 2, 3, 6, 7, 11, 14, 4, 8, 12, 5, 9, 13, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := preorder(tt.args.root); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("preorder() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
