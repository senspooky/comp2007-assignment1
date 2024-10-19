package aliasing

import (
	"fmt"

	"github.com/senspooky/comp2007-assignment1/demos"
	"github.com/senspooky/comp2007-assignment1/formatter"
)

type aliasing struct{}

func AliasingDemo() demos.Demo {
	return &aliasing{}
}

func (d *aliasing) Demo() {
	fmt.Println(formatter.F().Format("RESTRICTED ALIASING DEMO"))

	aliasingDemo()

	fmt.Println()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("My name is %s and I'm %d", p.Name, p.Age)
}

func aliasingDemo() {
	// foo and bar are pointers to a Person struct
	foo := &Person{
		Name: "Foo",
		Age:  30,
	}
	bar := &Person{
		Name: "bar",
		Age:  32,
	}

	// set bar to value of foo
	bar = foo

	// print both foo and bar
	fmt.Println("Value of foo:")
	fmt.Printf("%s\n", foo.String())
	fmt.Println("Value of bar:")
	fmt.Printf("%s\n", bar.String())

	// compare the pair
	if foo == bar {
		fmt.Println("foo and bar point to the same address in memory")
	}
}
