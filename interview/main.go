package main

type FooAdapter interface {
	Read()
	Write()
	Close()
}

type Foo struct {
	adapter FooAdapter
}

func NewFoo(a FooAdapter) *Foo {
	return &Foo{adapter: a}
}

type StubAdapter struct {
	FooAdapter // inherits Write() and Close()
}

func (s *StubAdapter) Read() {
	// simulate a read
}

func main() {
	f := NewFoo(&StubAdapter{})
	// do something with foo
	f.adapter.Read()
	f.adapter.Write()
}
