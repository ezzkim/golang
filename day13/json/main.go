package main

import (
	"encoding/json"
	"fmt"
)

var jsonStr = `{
	"students":[
		{
			"name": "stu01",
			"sex": "man"
		},
		{
			"name": "stu02",
			"sex": "woman"
		}
	]
}`

type Students struct {
	Students []Student `json:"students"`
}

type Student struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
}

func main() {
	testMarshal()
	fmt.Println("-----------")
	testUnmarshal()
}

func testMarshal() {
	var sts Students = Students{}
	st1 := Student{
		Name: "stu-01",
		Sex:  "man",
	}
	st2 := Student{
		Name: "stu-02",
		Sex:  "woman",
	}
	sts.Students = make([]Student, 0)
	sts.Students = append(sts.Students, st1)
	sts.Students = append(sts.Students, st2)

	fmt.Printf("sts = %v\n", sts)
	contents, err := json.Marshal(sts)
	if err != nil {
		fmt.Printf("marshal failure : err = %s\n", err.Error())
		return
	}
	fmt.Printf("%s\n", string(contents))
}

func testUnmarshal() {
	var sts Students = Students{}
	json.Unmarshal([]byte(jsonStr), &sts)
	fmt.Printf("%v\n", sts)
}
