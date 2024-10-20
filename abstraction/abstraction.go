package abstraction

import (
	"fmt"

	"github.com/senspooky/comp2007-assignment1/abstraction/queue"
	"github.com/senspooky/comp2007-assignment1/demos"
	"github.com/senspooky/comp2007-assignment1/formatter"
)

type abstraction struct{}

func AbstractionDemo() demos.Demo {
	return &abstraction{}
}

func (d *abstraction) Demo() {
	fmt.Println(formatter.F().Format("SUPPORT FOR ABSTRACTION DEMO"))

	queue := queue.New[int](5)

	for i := 1; i <= 5; i++ {
		fmt.Printf("Enqueuing %d\n", i)
		err := queue.Enqueue(i)
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
		fmt.Printf("Dequeued %d\n", v)
	}
}
