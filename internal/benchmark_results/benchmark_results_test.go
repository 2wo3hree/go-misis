package benchmark_results

import (
	"go-misis/internal/pkg/storage"
	"go.uber.org/zap"
	"testing"
)

func BenchmarkGet(b *testing.B) {
	log := zap.NewExample()
	s, _ := storage.NewStorage(log)
	key := "key1"
	value := "Goncharov"
	s.Set(key, value)
	
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Get(key)
	}
}

func BenchmarkSet(b *testing.B) {
	log := zap.NewExample()
	s, _ := storage.NewStorage(log)
	key := "key1"
	value := "Goncharov"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Set(key, value)
	}
}

func BenchmarkSetGet(b *testing.B) {
	log := zap.NewExample()
	s, _ := storage.NewStorage(log)
	key := "key1"
	value := "Goncharov"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Set(key, value)
		s.Get(key)
	}
}
