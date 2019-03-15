package main

import "fmt"
import "time"

func main() {
	fmt.Println(time.Now())
	fmt.Println("we are running the first app function in main")
	fmt.Println(time.Now())
	constatVariables()
	fmt.Println(time.Now())
	dataType()
}

func constatVariables() {
	// declarations of variables
	var a = 23
	// declarations of type variables
	var b string = "String One"
	var c int = 2
	// short hand declarations of variables
	p := 3
	fmt.Println("The value a: ", a, " the value b: ", b, " the value of c: ", c, " the p variable: ", p)
}

func dataType() {
	var a int = 3
	var b int8 = 23
	var c int16 = 3243
	var d int32 = 434
	var e int64 = 435434

	var f float32 = 4345443
	var g float64 = 435443

	var h string = "The string data stype"

	var i bool = true
	j := 93
	var k int = 23
	var m *int = &j
	// * get data from the memory location and
	// & get the memory location of the store value for the variable
	fmt.Println(a, b, c, d, e, f, g, h, i, j, " Memory addres of k: ", &k, " M pointing from J", *m)

}
func loopsInGo(){
	
}