package main

import (
	"fmt"
)

type Student struct {
	Username string
	Score    float32
	Grade    string
	Sex      int
}

func InputStudent() *Student {
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

	stu := newStudent(username, sex, score, grade)

	return stu
}

func newStudent(username string, sex int, score float32, grade string) *Student {
	student := &Student{
		Username: username,
		Sex:      sex,
		Score:    score,
		Grade:    grade,
	}

	return student
}
