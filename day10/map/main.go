package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	testMake()
	insertMap()
	mapOper()
	stu()
	fmt.Println("-----")
	mapDel()
	fmt.Println("-------")
	mapSlice()
	fmt.Println("-------")
	mapSlice2()
	fmt.Println("-------")
	mapSlice3()
}

func testMake() {
	var a map[string]int
	fmt.Printf("a:%v\n", a)
	if a == nil {
		//a = make(map[string]int) // the cap is 0
		a = make(map[string]int, 10) // the cap is 10
		a["stu01"] = 1000
	}
	//fmt.Println(a)
	a["one"] = 10 // you must make first
	fmt.Println(a)
}

func insertMap() {
	//define and malloc
	a := make(map[string]int)
	//init
	a["one"] = 10
	a["two"] = 20
	fmt.Printf("a : %v\n", a)

	//define and init, will auto make
	aa := map[string]int{
		"one-1": 123,
		"two-2": 222,
	}
	aa["three-3"] = 333
	fmt.Printf("aa : %v\n", aa)
}

func mapOper() {
	aa := map[string]int{
		"one-1": 123,
		"two-2": 222,
	}
	aa["three-3"] = 333
	fmt.Printf("aa : %v\n", aa)

	for k, v := range aa {
		fmt.Printf("%s : %d\n", k, v)
	}

	key := "no-exist"
	v, ok := aa[key]
	if !ok {
		aa[key] = 12345
	} else {
		fmt.Printf("key[%s] value[%d]\n", key, v)
	}
	fmt.Printf("aa : %v\n", aa)
}

func stu() {
	rand.Seed(time.Now().UnixNano())
	a := make(map[string]int, 1000)

	for i := 0; i < 128; i++ {
		key := fmt.Sprintf("stu-%d", i)
		value := rand.Intn(1000)
		a[key] = value
	}

	for k, v := range a {
		fmt.Printf("%s : %d\n", k, v)
	}
}

func mapDel() {
	a := make(map[string]int, 100)
	a["stu01"] = 1
	a["stu02"] = 2
	a["stu03"] = 3
	fmt.Printf("a=%#v\n", a)
	delete(a, "stu02") // delete
	fmt.Printf("a=%#v\n", a)

	//fmt.Printf("a len=%d cap=%d\n", len(a), cap(a)) //map not cap oper
	fmt.Printf("a len=%d\n", len(a))

	b := a
	b["stu03"] = 200
	fmt.Printf("a = %#v\n", a)

	modify(a)
	fmt.Printf("a = %#v\n", a)

	sortMapShow(a)
}

func modify(a map[string]int) {
	a["noexist"] = 1000
}

func sortMapShow(a map[string]int) {
	keys := make([]string, 0, 120)
	for key, _ := range a {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%s : %d\n", key, a[key])
	}
}

func mapSlice() {
	var mapSlice []map[string]int
	mapSlice = make([]map[string]int, 5)
	fmt.Println("Before map init")
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println()

	mapSlice[0] = make(map[string]int, 10) // must init map
	mapSlice[0]["a"] = 1000
	mapSlice[0]["b"] = 2000
	mapSlice[0]["c"] = 3000
	mapSlice[0]["d"] = 4000
	mapSlice[0]["e"] = 5000
	fmt.Println("After map init")
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}

func mapSlice2() {
	rand.Seed(time.Now().UnixNano())
	var s []map[string]int
	s = make([]map[string]int, 5, 16)
	for index, value := range s {
		fmt.Printf("slice[%d]=%v\n", index, value)
	}

	s[0] = make(map[string]int, 16)
	s[0]["stu01"] = 100
	s[0]["stu02"] = 200
	s[0]["stu03"] = 300
	for index, value := range s {
		fmt.Printf("slice[%d]=%v\n", index, value)
	}
}

func mapSlice3() {
	rand.Seed(time.Now().Unix())
	var s map[string][]int
	s = make(map[string][]int, 16)
	key := "stu01"
	value, ok := s[key]
	if !ok {
		s[key] = make([]int, 0, 16)
		value = s[key]
	}
	value = append(value, 100)
	value = append(value, 200)
	value = append(value, 300)
	s[key] = value // here why must assign value to s[key]
	// because value is another slice,
	// s[key] slice just init the value slice
	// after that, the value slice content have change

	fmt.Printf("map:%v\n", s)
}
