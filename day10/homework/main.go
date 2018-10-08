package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	subStrNumber("how do you do")
	fmt.Println("----------------")
	studentInfo()
}

func subStrNumber(str string) {
	substrs := strings.Split(str, " ")
	fmt.Printf("substrs : %v\n", substrs)
	strNumbers := make(map[string]int, 0)
	for _, item := range substrs {
		//fmt.Printf("item[%s]\n", item)
		value, ok := strNumbers[item]
		if !ok {
			strNumbers[item] = 1
		} else {
			value++
			strNumbers[item] = value
		}
	}

	for k, v := range strNumbers {
		fmt.Printf("%s : %d\n", k, v)
	}
}

func studentInfo() {
	//testInterface()

	//first key is id, second key is diff item
	var stuMap map[int]map[string]interface{}
	//insert student(id=1,name=stu01,score=78.5, age=18)
	stuMap = make(map[int]map[string]interface{}, 16)
	/*
		var id int = 1
		var name string = "stu01"
		var score float32 = 78.5
		var age int = 18

		value, ok := stuMap[id]
		if !ok {
			value = make(map[string]interface{}, 8)
		}
		value["name"] = name
		value["id"] = id
		value["score"] = score
		value["age"] = age

		stuMap[id] = value

		fmt.Printf("stumap = %v\n", stuMap)
	*/
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		value, ok := stuMap[i]
		if !ok {
			value = make(map[string]interface{}, 8)
		}
		value["name"] = fmt.Sprintf("stu%d", i)
		value["id"] = i
		value["score"] = rand.Float32() * 100
		value["age"] = rand.Intn(100)

		stuMap[i] = value
	}

	for k, v := range stuMap {
		fmt.Printf("id=%d stu info=%v\n", k, v)
	}
}

func testInterface() {
	//an empty interface can save any type data
	var a interface{}
	var b int = 100
	var c float32 = 12.1
	var d string = "hello"

	a = b
	fmt.Printf("a = %v\n", a)
	a = c
	fmt.Printf("a = %v\n", a)
	a = d
	fmt.Printf("a = %v\n", a)
}
