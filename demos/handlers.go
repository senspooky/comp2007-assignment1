package demos

type Demo interface {
	Run()
}

type DemoBuilder func() Demo
