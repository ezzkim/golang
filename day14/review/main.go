package main

import (
	"fmt"
	"os"
)

func main() {
	studentManager := GetStudentMgrInstance()

	for {
		showMenu()

		var sel int
		fmt.Scanf("%d\n", &sel)
		switch sel {
		case 1:
			err := studentManager.AddStudent()
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			}
		case 2:
			err := studentManager.ModifyStudent()
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			}
		case 3:
			studentManager.ShowAllStudent()
		case 4:
			os.Exit(0)
		}
	}
}

func showMenu() {
	fmt.Println("1. add student")
	fmt.Println("2. modify student")
	fmt.Println("3. show all student")
	fmt.Println("4. add exited\n\n")
}
