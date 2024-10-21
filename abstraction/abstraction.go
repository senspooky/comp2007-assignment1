package abstraction

import (
	"fmt"
	"math"

	"github.com/senspooky/comp2007-assignment1/abstraction/queue"
	"github.com/senspooky/comp2007-assignment1/demos"
	"github.com/senspooky/comp2007-assignment1/formatter"
)

type abstraction struct{}

func AbstractionDemo() demos.Demo {
	return &abstraction{}
}

type Float float64

func (m Float) Pow(pow Float) Float {
	return Float(math.Pow(float64(m), float64(pow)))
}

func (d *abstraction) Demo() {
	main()
}

func main() {
	fmt.Println(formatter.F().Format("SUPPORT FOR ABSTRACTION DEMO"))

	queue := queue.New[Float](5)

	for i := 1; i <= 5; i++ {
		fmt.Printf("Enqueuing %f\n", Float(i))
		err := queue.Enqueue(Float(i))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	for i := 1; i <= 5; i++ {
		fmt.Println("Dequeing...")
		v, err := queue.Dequeue()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Dequeued %f\n", v)
		pow := v.Pow(2)
		fmt.Printf("Result of dequeued float to the power of 2: %f\n", pow)
	}
}
