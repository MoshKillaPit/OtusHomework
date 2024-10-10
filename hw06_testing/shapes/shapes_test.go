package shapes

import "testing"

func Test_calculateArea(t *testing.T) {
	type args struct {
		s any
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:    "MinusArea",
			args:    args{rectangle{2, -2}},
			want:    0,
			wantErr: true,
		},
		{
			name:    "Shape",
			args:    args{-1},
			want:    0,
			wantErr: true,
		},
		{
			name:    "TestZero",
			args:    args{0},
			want:    0,
			wantErr: true,
		},
		{
			name:    "Circle",
			args:    args{circle{2}},
			want:    6.28,
			wantErr: false,
		},
		{
			name: "Rectangle",
			args: args{rectangle{
				height: 2,
				width:  2,
			}},
			want:    4,
			wantErr: false,
		},
		{
			name: "Triangle",
			args: args{triangle{
				height: 2,
				width:  2,
			}},
			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateArea(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculateArea() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calculateArea() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rectangle_Calculation(t *testing.T) {
	type fields struct {
		height float64
		width  float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "LogickTest",
			fields: fields{2, 2},
			want:   4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := rectangle{
				height: tt.fields.height,
				width:  tt.fields.width,
			}
			if got := r.Calculation(); got != tt.want {
				t.Errorf("Calculation() = %v, want %v", got, tt.want)
			}
		})
	}
}
