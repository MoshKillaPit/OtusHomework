package reader

import (
	"github.com/MoshKillaPit/OtusHomework/hw06_testing/fixapp/types"
	"reflect"
	"testing"
)

func TestReadJSON(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    []types.Employee
		wantErr bool
	}{
		{
			name: "Test",
			args: args{"../data.json"},
			want: []types.Employee{{
				UserID:       10,
				Age:          25,
				Name:         "Rob",
				DepartmentID: 3,
			},
				{
					UserID:       11,
					Age:          30,
					Name:         "George",
					DepartmentID: 2,
				},
			},
			wantErr: false,
		},
		{
			name:    "NotTrue",
			args:    args{"../not_found.json"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadJSON(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}
