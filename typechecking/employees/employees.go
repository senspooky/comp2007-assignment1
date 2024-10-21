package employees

import "fmt"

const fullTimeEquivalentHours = 38

type Employee struct {
	Name                     string
	HoursPerWeek             float32
	FullTimeEquivalentSalary int32
}

func (e Employee) GetFTE() float64 {
	var fte float64 = float64(e.HoursPerWeek) / fullTimeEquivalentHours
	return fte
}

func (e Employee) GetSalary() float64 {
	return e.GetFTE() * float64(e.FullTimeEquivalentSalary)
}

func (e Employee) String() string {
	return fmt.Sprintf(`Employee: %s
Hours per week: %f
Full Time Equivallent Salary: %d`, e.Name, e.HoursPerWeek, e.FullTimeEquivalentSalary)
}
