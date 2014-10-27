package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		name := "con:" + strconv.Itoa(i)
		go test(name)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("fin")
}

func test(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println(name, " hello ", i)
		time.Sleep(1 * time.Second)
	}
}
