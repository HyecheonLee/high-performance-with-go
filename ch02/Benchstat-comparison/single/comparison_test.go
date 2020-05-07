package single

import "testing"

func BenchmarkCarJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CarJSON()
	}
}
