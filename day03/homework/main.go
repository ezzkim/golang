package main

import (
	"fmt"
	"time"
)

func main() {
	testFormat()
	testFormat2()
	testUseTime()
}

func testFormat() {
	now := time.Now()
	timeStr := now.Format("2006/01/02 15:04:05")
	fmt.Printf("time:%s\n", timeStr)
}

func testFormat2() {
	now := time.Now()

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func testUseTime() {
	start := time.Now().UnixNano()

	for i := 0; i < 5; i++ {
		fmt.Printf("sleep %d second\n", i)
		time.Sleep(time.Second)
	}

	end := time.Now().UnixNano()
	microCost := (end - start) / 1000
	millCost := (end - start) / 1000 / 1000

	fmt.Printf("cost times : %d micro seconds\n", microCost)
	fmt.Printf("cost times : %d mill seconds\n", millCost)
}
