package main
import "testing"
func getHello() string {
	return "hello"
}

func getName() string{
	return "Hello world"
}

func TestName(t *testing.T){
	name := getName()
	if name != "Hello world" {
		t.Error("no matchin string")
	}
}