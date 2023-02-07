package main

import (
	"fmt"
	"time"
)

func timeConversion(s string) string {

	layout1 := "03:04:05PM"
	layout2 := "15:04:05"
	t, err := time.Parse(layout1, s)
	if err != nil {
		panic(err)
	}
	return (t.Format(layout2))
}

func main() {

	fmt.Printf("%s", timeConversion("07:05:45PM"))
}
