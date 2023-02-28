package std_time

import (
	"fmt"
	"time"
)

func TestTime() {
	testNow()
	testTimeStamp()
}

func testNow() {
	now := time.Now()
	fmt.Printf("now: %T\n", now)
	fmt.Printf("now: %v\n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	fmt.Printf("%T %T %T %T %T %T\n", year, month, day, hour, minute, second)
}

func testTimeStamp() {
	now := time.Now()
	fmt.Printf("now.Unix(): %v\n", now.Unix())
}
