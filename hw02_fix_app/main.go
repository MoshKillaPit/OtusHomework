package main

import (
	"fmt"
	"github.com/fixme_my_friend/hw06_testing/fix_app/types"
	"github.com/fixme_my_friend/hw06_testing/fix_app/types/printer"
	"github.com/fixme_my_friend/hw06_testing/fix_app/types/reader"
)

func main() {
	path := "data.json"
	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path)
	if err != nil {
		panic(err)
	}

	printer.PrintStaff(staff)
}
