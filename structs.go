package main

import "fmt"

type Shape struct {
	radius float64
}

func HandleError(){

	// recover holds the error that's thrown from panic
	r := recover()
	fmt.Println(r)
}

func (s Shape) getArea() float64 {

	// defer runs the following function after the enclosing function finishes its execution
	defer HandleError()
	panic("Function not implemented")
}

type Circle interface {
	getArea() float64
}


func main() {
	var circle Circle
	circle = Shape{radius: 2}
	circle.getArea()
}
