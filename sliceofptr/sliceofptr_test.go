package sliceofptr

import (
	"testing"
)

func TestSliceOfPtr(t *testing.T) {
	persons := []*Person{
		{Name: "Tony", Age: 33},
		{Name: "Jenny", Age: 18},
	}
	ageInc(persons)
	for _, p := range persons {
		t.Log(p.Age)
	}
}

func ageInc(persons []*Person) {
	for _, p := range persons {
		p.Age = p.Age + 1
	}
}
