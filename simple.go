package main

import (
  "fmt"
  "time"
)

func main() {
  test()
  time.Sleep(3 * time.Second)
  fmt.Println("fin")
}

func test() {
  for i := 0; i < 5; i++ {
    fmt.Println("hello", i)
    time.Sleep(1 * time.Second)
  }

}