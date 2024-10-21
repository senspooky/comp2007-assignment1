package syntax

import (
	"fmt"

	"github.com/senspooky/comp2007-assignment1/demos"
	"github.com/senspooky/comp2007-assignment1/syntax/slices"
)

func SyyntaxDemo() demos.Demo {
	return &syntaxdesign{}
}

type syntaxdesign struct{}

func (d *syntaxdesign) Demo() { main() }

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("SYNTAX DESIGN DEMO")
	fmt.Println()

	sliceOfInts := []int{2, 4, 5, 7, 1, 6, 7, 2, 91}

	sliceOfPeople := []Person{
		{
			Name: "Gordon",
			Age:  60,
		},
		{
			Name: "Nicholas",
			Age:  43,
		},
		{
			Name: "Stephen",
			Age:  39,
		},
		{
			Name: "Jessica",
			Age:  24,
		},
		{
			Name: "Madeleine",
			Age:  40,
		},
		{
			Name: "Amanda",
			Age:  22,
		},
	}

	fmt.Println("Does slice of ints have 7 in it?")
	if slices.HasComparable(sliceOfInts, 7) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	fmt.Println()

	fmt.Println("Does slice of people have somebody over the age of 70 in it?")
	if slices.HasFunc(sliceOfPeople, func(e Person) bool {
		return e.Age > 70
	}) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	fmt.Println()

	fmt.Println("Print all occurances of 2 in the slice of ints")
	result := slices.FilterComparable(sliceOfInts, 2)
	for _, e := range result {
		fmt.Println(e)
	}
	fmt.Println()

	fmt.Println("Print all occurances of people whose name is Amanda in the slice of People")
	personResult := slices.FilterFunc(sliceOfPeople, func(e Person) bool {
		return e.Name == "Amanda"
	})
	for _, e := range personResult {
		fmt.Printf("%s, %d", e.Name, e.Age)
	}
	fmt.Println()
}
