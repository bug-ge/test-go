package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Once

	go m.Do(func1)
	go m.Do(func1)
	time.Sleep(time.Second)
}

func func1() {
	fmt.Println("hello world")
}
