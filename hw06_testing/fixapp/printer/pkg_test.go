package printer

import (
	"testing"

	"github.com/MoshKillaPit/OtusHomework/hw06_testing/fixapp/types"
)

func TestPrintStaff(t *testing.T) {
	type args struct {
		staff []types.Employee
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test",
			args: args{staff: []types.Employee{{
				UserID:       2,
				Age:          12,
				Name:         "Jon",
				DepartmentID: 1,
			}}},
		},
		{
			name: "Empty",
			args: args{[]types.Employee{}},
		},
		{
			name: "nil",
			args: args{nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			PrintStaff(tt.args.staff)
		})
	}
}
