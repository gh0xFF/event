package utils

import "testing"

func BenchmarkCompression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Compress(uncompressed)
	}
}
