package simplicityorthogonality

import (
	"fmt"

	"github.com/senspooky/comp2007-assignment1/demos"
	"github.com/senspooky/comp2007-assignment1/formatter"
	"github.com/senspooky/comp2007-assignment1/simplicityorthogonality/pages"
)

type simplicity struct{}

func SimplicityDemo() demos.Demo {
	return &simplicity{}
}

func (d *simplicity) Demo() {
	main()
}

const coordinateCount = 50

type coordinate struct {
	x, y, z int
}

func (c coordinate) String() string {
	return fmt.Sprintf("x: %d, y: %d, x: %d", c.x, c.y, c.z)
}

func main() {
	fmt.Println(formatter.F().Format("SIMPLICITY DEMO"))

	// ...
	// create an array of coordinates to test pager
	coordinates := createCoordinates()

	// create a new pager with page size 5
	pager := pages.New(5, coordinates)

	var dataPage []coordinate
	for i := 0; true; i++ {
		// get page i
		dataPage = pager.Page(i)

		// print each page
		if len(dataPage) > 0 {
			fmt.Printf("Printing Page %d\n", i+1)
			for _, c := range dataPage {
				fmt.Println(c.String())
			}
		}

		// while dataPage is not nil
		if dataPage == nil {
			break
		}
	}

	fmt.Println()
}

func createCoordinates() []coordinate {
	coordinates := make([]coordinate, 0, coordinateCount)
	for i := 0; i < coordinateCount; i++ {
		coordinates = append(coordinates, coordinate{
			x: i,
			y: i*i - (coordinateCount / 2),
			z: i*i*i + 2*i*i - 16,
		})
	}
	return coordinates
}
