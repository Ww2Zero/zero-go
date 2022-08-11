package _22_coinChange

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				coins:  []int{1, 2, 5},
				amount: 11,
			},
			want: 3,
		},
		{
			name: "test1",
			args: args{
				coins:  []int{2},
				amount: 3,
			},
			want: -1,
		},
		{
			name: "test2",
			args: args{
				coins:  []int{1, 2, 5},
				amount: 63,
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChangeBrute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_coinChangeBrute(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				coins:  []int{1, 2, 5},
				amount: 11,
			},
			want: 3,
		},
		{
			name: "test1",
			args: args{
				coins:  []int{2},
				amount: 3,
			},
			want: -1,
		},
		{
			name: "test2",
			args: args{
				coins:  []int{1, 2, 5},
				amount: 63,
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := coinChangeBrute(tt.args.coins, tt.args.amount); gotRes != tt.want {
				t.Errorf("coinChangeBrute() = %v, want %v", gotRes, tt.want)
			}
		})
	}
}
