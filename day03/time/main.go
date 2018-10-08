package main

import (
	"fmt"
	"time"
)

func main() {
	//testTime()
	timeStamp := time.Now().Unix()
	testTimeStamp(timeStamp)
	//testTicker()
	testConst()
	testFormat()
}

func testTime() {
	now := time.Now()
	fmt.Printf("now time : %v\n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	timeStamp := now.Unix()
	fmt.Printf("timeStamp : %d\n", timeStamp) // from 1970

	//timeStamp convert to time
}

func testTimeStamp(timeStamp int64) {
	timeobj := time.Unix(timeStamp, 0)

	year := timeobj.Year()
	month := timeobj.Month()
	day := timeobj.Day()
	hour := timeobj.Hour()
	minute := timeobj.Minute()
	second := timeobj.Second()

	fmt.Printf("timeStmap = %d\n", timeStamp)
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

//定时器
func testTicker() {
	ticker := time.Tick(time.Second) // create a tunnal and send now time to the ticker's channel
	for i := range ticker {
		fmt.Printf("now : %v\n", i)
		timeOutHandler()
	}
}

func timeOutHandler() {
	fmt.Println("call timer handler")
}

func testConst() {
	fmt.Printf("nano second:%d\n", time.Nanosecond)
	fmt.Printf("micro second:%d\n", time.Microsecond)
	fmt.Printf("mill second:%d\n", time.Millisecond)
	fmt.Printf("second:%d\n", time.Second)
}

func testFormat() {
	now := time.Now()
	//timeStr := now.Format("2006/01/02 : 15:04:05")
	timeStr := now.Format("2006-01-02 15:04:05") // must write like this, the go birth-time
	fmt.Printf("time : %s\n", timeStr)
}
