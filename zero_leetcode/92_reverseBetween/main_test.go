package main

import (
	"reflect"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				head:  NewList(1, 2, 3, 4, 5),
				left:  2,
				right: 4,
			},
			want: NewList(1, 4, 3, 2, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
