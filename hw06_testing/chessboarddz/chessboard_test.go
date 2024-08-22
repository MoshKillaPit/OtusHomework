package chessboarddz

import "testing"

func Test_value(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		{
			name:    "TestSize",
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := value()
			if (err != nil) != tt.wantErr {
				t.Errorf("value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("value() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_chessboard(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "Test",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chessboard(); got != tt.want {
				t.Errorf("chessboard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_chessboarddz(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chessboarddz()
		})
	}
}
