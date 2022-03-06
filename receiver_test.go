package main

import (
	"testing"
)

func BenchmarkPointerRecever(b *testing.B) {
	p := makeProductInterfaceOnPointer()
	for i := 0; i < b.N; i++ {
		_ = CastToMedium(p)
	}
}

func BenchmarkValueRecever(b *testing.B) {
	p := makeProductInterfaceOnStruct()
	for i := 0; i < b.N; i++ {
		_ = CastToMedium(p)
	}
}
