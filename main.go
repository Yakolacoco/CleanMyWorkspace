package main

import (
	"CleanMyWorkspace/Clean"
	"fmt"

	"github.com/Mentor-Paris/CleanMyWorkspace"
)

func main() {
	workspace := CleanMyWorkspace.GenererateWorkSpace()

	fmt.Println("===== Espace de travail sale =====")
	for _, row := range *workspace {
		for _, item := range row {
			if item == nil {
				fmt.Print("|nil")
			} else {
				fmt.Print("|" + *item)
			}
		}
		fmt.Println("|")
	}

	cleaned := Clean.CleanWorkSpace(workspace)

	fmt.Println("\n===== Espace de travail nettoyé =====")
	for _, row := range *cleaned {
		for _, item := range row {
			if item == nil {
				fmt.Print("|nil")
			} else {
				fmt.Print("|" + *item)
			}
		}
		fmt.Println("|")
	}
}
