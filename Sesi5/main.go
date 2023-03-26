package main

import (
	"fmt"
	"time"
)

func main() {
	go taskA()
	taskB()
	time.Sleep(2 * time.Second)
}
func taskA() {
	fmt.Println("Task A")
}
func taskB() {
	fmt.Println("Task B")
}
