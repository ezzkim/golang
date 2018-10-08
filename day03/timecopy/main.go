package main

import (
	"fmt"
	"time"
)

func testTime() {
	now := time.Now()
	fmt.Printf("current time : %v\n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	timeStamp := now.Unix()
	fmt.Printf("timeStamp : %v\n", timeStamp)
}

func testTimeStamp(timeStamp int64) {
	timeObj := time.Unix(timeStamp, 0)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()

	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

}

func testTime2() {
	//ticker := time.Tick(time.Second)
	ticker := time.Tick(time.Second * 3)
	for i := range ticker {
		fmt.Printf("%v\n", i)
		processTask()
	}
}

func processTask() {
	fmt.Println("do task")
}

func testConst() {
	fmt.Printf("nm : %d\n", time.Nanosecond)
	fmt.Printf("mic s : %d\n", time.Microsecond)
	fmt.Printf("mili : %d\n", time.Millisecond)
	fmt.Printf("s : %d\n", time.Second)
}

func main() {
	testTime()
	timeStamp := time.Now().Unix()
	testTimeStamp(timeStamp)

	testConst()
	testFormat()
	testCost()
	testTime2()
}

func testCost() {
	start := time.Now().UnixNano()
	for i:=0;i<10; i++ {
		time.Sleep(time.Millisecond)
	}
	end := time.Now().UnixNano()
	cost := (end - start)/1000
	fmt.Printf("cost : %d us\n", cost)
}

func testFormat() {
	now := time.Now()
	//timeStr := now.Format("2006/01/02 : 15:04:05") 
	timeStr := now.Format("2006-01-02 : 15:04:05") // must write like this, the go birth-time
	fmt.Printf("time : %s\n", timeStr)
}
