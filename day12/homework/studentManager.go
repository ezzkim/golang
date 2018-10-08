package main

import (
	"fmt"
	"os"
	//"flag"
)

type Student struct {
	Username string
	Score    float32
	Grade    string
	Sex      int
}

var Allstudents = make([]*Student, 0)

func showMenu() {
	fmt.Println("1. add student")
	fmt.Println("2. modify student")
	fmt.Println("3. show all student")
	fmt.Println("4. add exited\n\n")
}

func AddStudent() {
	stu := inputStudent()

	for index, v := range Allstudents {
		if v.Username == stu.Username {
			Allstudents[index] = stu // update
			return
		}
	}
	Allstudents = append(Allstudents, stu)
	fmt.Printf("insert usr %s success\n", stu.Username)
}

func ModifyStudent() {
	stu := inputStudent()
	
	for index, v := range Allstudents {
		if v.Username == stu.Username {
			Allstudents[index] = stu // update
			fmt.Printf("modify usr %s success\n", v.Username)
			return
		}
	}
	fmt.Printf("user %s is not found\n", stu.Username)
}

func ShowAllStudent() {
	for _, v := range Allstudents {
		fmt.Printf("user:%s info:%#v\n", v.Username, v)
	}
}

func NewStudent(username string, sex int, score float32, grade string) *Student {
	student := &Student{
		Username: username,
		Sex:      sex,
		Score:    score,
		Grade:    grade,
	}

	return student
}

func inputStudent() *Student {
	var (
		username string
		sex      int
		grade    string
		score    float32
	)
	
	fmt.Println("Please inpuit username:")
	fmt.Scanf("%s\n", &username)
	fmt.Println("Please inpuit sex[0|1]:")
	fmt.Scanf("%d\n", &sex)
	fmt.Println("Please inpuit grade[0-6]:")
	fmt.Scanf("%s\n", &grade)
	fmt.Println("Please inpuit score[0-100]:")
	fmt.Scanf("%f\n", &score)

	stu := NewStudent(username, sex, score, grade)

	return stu
}

func main() {
	for {
		showMenu()

		var sel int
		fmt.Scanf("%d\n", &sel)
		switch sel {
		case 1:
			AddStudent()
		case 2:
			ModifyStudent()
		case 3:
			ShowAllStudent()
		case 4:
			os.Exit(0)
			//break
		}
	}
}
