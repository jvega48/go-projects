package main

import "fmt"
type Person struct{
	name string
	age int
	bday string
}


type numbers struct {
	A int
	B int
}

func Multiply(a int , b int) int {
	return a * b
}

func home(x *Person) (int, string, string) {
	return x.age, x.name, x.bday
}
func main(){ 
	person1 := Person{name: "Jose", age: 0000, bday: "32/232/32"}
	clonen,clonea,cloneb := home(&person1)
	fmt.Println(clonen, clonea,cloneb)
	num := numbers{4,3}
	sum := Multiply(num.A, num.B)
	fmt.Println(sum)
}
