package main

import (
	"testing"
)

func BenchmarkPointerRecever(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := makeProductInterfaceOnPointer()
		p.SetId(i)
		_ = CastToMedium(p)
	}
}

func BenchmarkValueRecever(b *testing.B) {
	p := makeProductInterfaceOnStruct()
	for i := 0; i < b.N; i++ {
		p.SetId(i)
		_ = CastToMedium(p)
	}
}
