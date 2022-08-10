package main

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test2",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "test3",
			args: args{n: 3},
			want: 2,
		},
		{
			name: "test4",
			args: args{n: 4},
			want: 3,
		},
		{
			name: "test100",
			args: args{n: 100},
			want: 3736710778780434371,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibBase(tt.args.n); got != tt.want {
				t.Errorf("fibBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fibDp(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test2",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "test3",
			args: args{n: 3},
			want: 2,
		},
		{
			name: "test4",
			args: args{n: 4},
			want: 3,
		},
		{
			name: "test100",
			args: args{n: 100},
			want: 3736710778780434371,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibMemo(tt.args.n); got != tt.want {
				t.Errorf("fibMemo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fibDP(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test2",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "test3",
			args: args{n: 3},
			want: 2,
		},
		{
			name: "test4",
			args: args{n: 4},
			want: 3,
		},
		{
			name: "test100",
			args: args{n: 100},
			want: 3736710778780434371,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibDP(tt.args.n); got != tt.want {
				t.Errorf("fibDP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fib1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test2",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "test3",
			args: args{n: 3},
			want: 2,
		},
		{
			name: "test4",
			args: args{n: 4},
			want: 3,
		},
		{
			name: "test100",
			args: args{n: 100},
			want: 3736710778780434371,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
