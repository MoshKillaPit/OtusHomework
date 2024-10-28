package main

import "testing"

func Test_worker(t *testing.T) {
	type args struct {
		countGorutines int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Usual",
			args: args{100},
			want: 102,
		},
		{
			name: "One Gorutine",
			args: args{1},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := worker(tt.args.countGorutines); got != tt.want {
				t.Errorf("worker() = %v, want %v", got, tt.want)
			}
		})
	}
}
