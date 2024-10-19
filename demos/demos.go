package demos

type Demo interface {
	Demo()
}

type DemoBuilder func() Demo
