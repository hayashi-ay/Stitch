package main_test

import (
	main "Stitch"
	"testing"
)

func Test_ToFileBasename(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal",
			args: args{"https://hayapenguin.com"},
			want: "hayapenguin.com",
		},
		{
			name: "trailing slash is removed",
			args: args{"https://hayapenguin.com/"},
			want: "hayapenguin.com",
		},
		{
			name: "filepath is given",
			args: args{"https://hayapenguin.com/preseed.cfg"},
			want: "hayapenguin.com/preseed.cfg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := main.ToFileBasename(tt.args.url); got != tt.want {
				t.Errorf("ToFileBasename() = %v, want %v", got, tt.want)
			}
		})
	}
}
