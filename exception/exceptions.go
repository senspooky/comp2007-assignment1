package exceptions

import (
	"errors"
	"fmt"

	"github.com/senspooky/comp2007-assignment1/demos"
)

type exceptionHandling struct{}

func ExceptionDemo() demos.Demo {
	return &exceptionHandling{}
}

func (d *exceptionHandling) Demo() {
	main()
}

func main() {
	fmt.Println("EXCEPTION HANDLING DEMO")
	fmt.Println("ERRORS")
	errorDemo()
	fmt.Println()
	fmt.Println("PANICS")
	panicDemo()
	fmt.Println()
}

func errorDemo() {
	// Lets try divide by zero
	fmt.Print("Dividing 10 by 2...\n")
	result2, err := lessAnxiousDivisionFunction(10.0, 0)
	if err != nil { // check the error
		fmt.Printf("error encountered while dividing: %v\n", err)
		return // end execution
	}
	fmt.Printf("Result: %f\n", result2)
}

func panicDemo() {
	defer func() {
		v := recover()
		if v != nil {
			fmt.Printf("recovered from a panic: %v\n", v)
		}
	}()

	// Lets try divide by zero
	fmt.Print("Dividing 10 by 2\n")
	result2 := anxiousDivisionFunction(10.0, 0)
	fmt.Printf("Result: %f\n", result2)
}

func anxiousDivisionFunction(numerator, denominator float64) float64 {
	if denominator == 0 {
		// panic; you can't divide by zero!
		panic("Uh-oh! You divided by zero. Hope you call recover()!")
	}
	result := numerator / denominator

	return result
}

func lessAnxiousDivisionFunction(numerator, denominator float64) (float64, error) {
	if denominator == 0 {
		// return an error; you can't divide by zero!
		return 0, errors.New("Uh-oh! You divided by zero. Better check the returned error!")
	}
	result := numerator / denominator

	return result, nil // return a nil error, because no error occurred
}
