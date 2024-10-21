package typechecking

import (
	"fmt"

	"github.com/senspooky/comp2007-assignment1/demos"
	"github.com/senspooky/comp2007-assignment1/typechecking/employees"
)

func TypeCheckingDemo() demos.Demo {
	return &typechecking{}
}

type typechecking struct{}

func (d *typechecking) Demo() {
	main()
}

func main() {
	fmt.Println("TYPE CHECKING DEMO")

	foo := employees.Employee{
		Name:                     "Foo",
		HoursPerWeek:             38,
		FullTimeEquivalentSalary: 79000,
	}
	bar := employees.Employee{
		Name:                     "Bar",
		HoursPerWeek:             19,
		FullTimeEquivalentSalary: 112000,
	}

	fmt.Println("Employees:")
	fmt.Println()
	fmt.Println(foo.String())
	fmt.Println(bar.String())
	fmt.Println()

	for _, e := range []employees.Employee{foo, bar} {
		fmt.Printf("%s's Full Time Equivalent Ratio: %f\n", e.Name, e.GetFTE())
		fmt.Printf("%s's Actual Salary: %f\n", e.Name, e.GetSalary())
		fmt.Println()
	}
}
