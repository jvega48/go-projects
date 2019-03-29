package main

import "fmt"
import "time"
import "io/ioutil"
import "net/http"

func main() {
	getNetWorkAddress()
	fmt.Println(time.Now())
	fmt.Println("we are running the first app function in main")
	fmt.Println(time.Now())
	constatVariables()
	fmt.Println(time.Now())
	dataType()
	fmt.Println(time.Now())
	loopsInGo()
	var arrayOne[3] int = sort()
	fmt.Println(arrayOne)
	// array initiation
	var strArray[3] string 
	strArray[0] = "String"
	strArray[1] = "Two" 
	strArray[2] = "Three"
	// short hand array initiation
	srtArr := [1]int{3}
	fmt.Println(srtArr)
	// short hand array declarations
	numbers :=[...]int {2,3,2,3}
	// make a copy and slice and array
	p := make([]int, len(strArray))
	z := buildingStruct(3,2)
	fmt.Println(p)
	fmt.Println(z)
	fmt.Println(numbers)

	fmt.Println(*pointers())
	// lambda

	// myLambdaFunction := func() bool { return 10 > 1000 }
	// fmt.Println("Lambda functions: ", myLambdaFunction)

	fmt.Println(subtractions(99))
}
func buildingStruct(a int,b int) int{
	return a * b
}

// name return values 
// specify the value names in the signature
func subtractions(sum int) (x int, y int){
	x = sum * 3
	y = sum - y
	return
}

func pointers() (value *int){
	g := 23424
	return &g
}

func sort() [3]int{
	return [...]int{2,3,4}
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
	i := 1
	for i < 10{
		fmt.Println(i)
		i++
	}
	for ii := 1; ii <= 10; ii++{
		fmt.Println(ii)
	}

	for a := 1; a <= 10; a++{
		fmt.Println("The number: ", a)
		if a%2 == 0 {
			fmt.Println("even: ", a)
		} else {
			fmt.Println("odd: ", a)
		}
	}

	// array of strings 
	name := []string{"jose", "pancho","jack"}
	// range for loop
	for a, val := range name{
		fmt.Printf("At position %d, the character %s \n", a, val)
	}
}

func getNetWorkAddress() {
	resp, _ := http.Get("https://development-double-door.herokuapp.com")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
	resp.Body.Close()
}

