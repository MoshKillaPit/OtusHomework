package chessboarddz

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

func Test_paint1(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test",
			args: args{0},
			want: "Ошибка значения: размер должен быть больше нуля",
		},
		{
			name: "Norm",
			args: args{1},
			want: " \n",
		},
		{
			name: "Test",
			args: args{4},
			want: " # #\n# # \n # #\n# # \n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := paint(tt.args.size); got != tt.want {
				t.Errorf("paint() = %v, want %v", got, tt.want)
			}
		})
	}
}
