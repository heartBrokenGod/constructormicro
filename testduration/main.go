package main

import (
	"fmt"
	"time"
)

func main() {
	s := "34333h1000s2m"
	d, err := time.ParseDuration(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}
