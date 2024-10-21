package storage

import (
	"go.uber.org/zap"
	"testing"
)

func TestStorageSetGet(t *testing.T) {
	log := zap.NewExample()
	s, _ := NewStorage(log)

	tests := []struct {
		key      string
		value    interface{}
		expected interface{}
		whatKind interface{}
	}{
		{"key1", "Goncharov", "Goncharov", "S"},
		{"key2", 22, "22", "D"},
		{"key3", true, "true", ""},
		{"key4", 3.14, "3.14", "D"}, //
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			s.Set(tt.key, tt.value)
			got := s.Get(tt.key)

			if *got != tt.expected {
				t.Errorf("for key %s %v != %v", tt.key, tt.expected, *got)
			}

			gotKind := s.GetKind(tt.key)
			if gotKind != tt.whatKind {
				t.Errorf("for key %s  kind %q != %q", tt.key, tt.whatKind, gotKind)
			}
		})

	}

}
