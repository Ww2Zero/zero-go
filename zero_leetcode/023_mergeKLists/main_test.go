package _23_mergeKLists

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test",
			args: args{
				lists: []*ListNode{
					NewList(1, 4, 5),
					NewList(1, 3, 4),
					NewList(2, 6)},
			},
			want: NewList(1, 1, 2, 3, 4, 4, 5, 6),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
