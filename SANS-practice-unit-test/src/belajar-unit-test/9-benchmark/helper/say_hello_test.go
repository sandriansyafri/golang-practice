package helper

import (
	"belajar-unit-test/belajar-unit-test/9-benchmark/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkSayHello(t *testing.B) {
	for i := 0; i < t.N; i++ {
		SayHello("R")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Benchmark-1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("AAAAAAAA")
		}
	})
	b.Run("Benchmark-2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("AA")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	benchs := []*entity.Bench{
		{Name: "B-1", Request: "A"},
		{Name: "B-2", Request: "AA"},
		{Name: "B-3", Request: "AAA"},
	}

	for _, bench := range benchs {
		b.Run(bench.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SayHello(bench.Request)
			}
		})
	}

}

func TestTable(t *testing.T) {
	tests := []*entity.Test{
		{Name: "Rian", Expect: "Hello Rian", Request: "Rian"},
		{Name: "Hafid", Expect: "Hello Hafid", Request: "Hafid"},
		{Name: "Yogi", Expect: "Hello Yogi", Request: "Yogi"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := SayHello(test.Request)
			assert.Equal(t, test.Expect, result)
		})
	}

}
