package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello")
	fmt.Println("The Current Time is: ", time.Now())
	one()
	two()
	three()
	a := []int{3, 4, 3, 5}
	a = append(a, 9)
	fmt.Println(a)

}

func one() {
	time.Sleep(100 * 100 * time.Millisecond)
	fmt.Println("One: ", time.Now())
}
func two() {
	time.Sleep(200 * 200 * time.Millisecond)
	fmt.Println("two: ", time.Now())
}
func three() {
	time.Sleep(300 * 300 * time.Millisecond)
	fmt.Println("three: ", time.Now())
}
