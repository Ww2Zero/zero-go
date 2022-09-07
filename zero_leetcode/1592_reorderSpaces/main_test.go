package _592_reorderSpaces

import "testing"

func Test_reorderSpaces(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{text: "  this   is  a sentence "},
			want: "this   is   a   sentence",
		},
		{
			name: "test2",
			args: args{
				text: "  hello",
			},
			want: "hello  ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorderSpaces(tt.args.text); got != tt.want {
				t.Errorf("reorderSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
