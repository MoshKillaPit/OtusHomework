package main

import "testing"

func Test_sort(t *testing.T) {
	type args struct {
		pages *[]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test",
			args: args{pages: &[]int{2, 1, 3, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			sort(tt.args.pages)
		})
	}
}

func Test_binarSearch(t *testing.T) {
	type args struct {
		needle int
		pages  []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test",
			args: args{
				needle: 10,
				pages:  []int{2, 1, 3, 5},
			},
			want: false,
		},
		{
			name: "Nil",
			args: args{needle: -1, pages: nil},
			want: true,
		},
		{
			name: "Norm",
			args: args{needle: 10, pages: []int{2, 1, 3, 10}},
			want: false,
		},
		{
			name: "Empty",
			args: args{
				needle: 2,
				pages:  []int{2, 2, 2, 2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarSearch(tt.args.needle, tt.args.pages); got != tt.want {
				t.Errorf("binarSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
