package timeformat

import (
	"testing"
	"time"
	"fmt"
)

func TestAll(_ *testing.T) {
	t := time.Now()

	fmt.Println("FormatTime(t)", FormatTime(t))
	fmt.Println("FormatDate(t)", FormatDate(t))
	fmt.Println("FormatMonth(t)", FormatMonth(t))

	fmt.Println("ParseTime(FormatTime(t)).String()", ParseTime(FormatTime(t)).String())
	fmt.Println("ParseDate(FormatDate(t)).String()", ParseDate(FormatDate(t)).String())
	fmt.Println("ParseMonth(FormatMonth(t)).String()", ParseMonth(FormatMonth(t)).String())

	fmt.Println("ParseTime(FormatMonth(t))", ParseTime(FormatMonth(t)))
	fmt.Println("ParseDate(FormatTime(t))", ParseDate(FormatTime(t)))
	fmt.Println("ParseTime(0)", ParseTime(0))

	fmt.Println("FormatTimeNow()", FormatTimeNow())
	fmt.Println("FormatDateNow()", FormatDateNow())
	fmt.Println("FormatMonthNow()", FormatMonthNow())
}
