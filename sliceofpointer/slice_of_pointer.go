package sliceofpointer

type foo struct {
	a int
}

type bar struct{
	b int
}

// Dummy -
type Dummy struct {
	foo *foo
	bar *bar
}
