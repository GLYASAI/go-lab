package sliceofpointer

import "testing"

func BenchmarkSlicePointers(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		slice := make([]*Dummy, 0, 100)
		for j := 0; j < 100; j++ {
			slice = append(slice, &Dummy{foo: &foo{}, bar: &bar{}})
		}
	}
}

func BenchmarkSliceNoPointers(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		slice := make([]Dummy, 0, 100)
		for j := 0; j < 100; j++ {
			slice = append(slice, Dummy{foo: &foo{}, bar: &bar{}})
		}
	}
}