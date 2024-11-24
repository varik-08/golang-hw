package main

import (
	"fmt"

	"github.com/fixme_my_friend/hw02_fix_app/printer"
	"github.com/fixme_my_friend/hw02_fix_app/reader"
	"github.com/fixme_my_friend/hw02_fix_app/types"
)

func main() {
	var path string
	var staff []types.Employee

	fmt.Printf("Enter data file path: ")

	_, err := fmt.Scanln(&path)
	if err != nil {
		if len(path) == 0 {
			path = "data.json"
		}
	}

	staff, err = reader.ReadJSON(path)
	if err != nil {
		fmt.Printf("Error ReadJSON: %v", err)
		return
	}

	printer.PrintStaff(staff)
}
