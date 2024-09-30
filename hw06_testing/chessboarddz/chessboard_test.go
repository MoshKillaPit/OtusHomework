package main

import "testing"

func Test_checksize(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Test",
			args:    args{2},
			wantErr: false,
		},
		{
			name:    "Zero",
			args:    args{0},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checksize(tt.args.size); (err != nil) != tt.wantErr {
				t.Errorf("checksize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_paint(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test",
			args: args{2},
		},
		{
			name: "Zero",
			args: args{0},
		},
		{
			name: "",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			paint(tt.args.size)
		})
	}
}
