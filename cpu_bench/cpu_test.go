package cpu_bench

import "testing"

func BenchmarkSync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		seq()
	}
}

func BenchmarkParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parallel()
	}
}

func BenchmarkLimitParallelTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		limitParallel(2)
	}
}

func BenchmarkLimitParallelFour(b *testing.B) {
	for i := 0; i < b.N; i++ {
		limitParallel(4)
	}
}

func BenchmarkLimitParallelEight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		limitParallel(8)
	}
}
