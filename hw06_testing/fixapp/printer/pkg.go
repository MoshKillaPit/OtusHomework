package printer

import (
	"fmt"

	"github.com/MoshKillaPit/OtusHomework/hw06_testing/fixapp/types"
)

func PrintStaff(staff []types.Employee) {
	for _, v := range staff {
		fmt.Println(v.String())
	}
}
