package demos

import "fmt"

type exceptionHandling struct{}

func ExceptionHandlingDemo() Demo {
	return &exceptionHandling{}
}

func (d *exceptionHandling) Run() {
	fmt.Print("not implemented")
}
