package main

import (
	"fmt"
	"go-misis/internal/pkg/storage"
	"go.uber.org/zap"
)

func main() {
	log := zap.NewExample()
	s, _ := storage.NewStorage(log)

	s.Set("key1", "Goncharov")
	s.Set("key2", "22")

	fmt.Println(*s.Get("key1"))
	fmt.Println(*s.Get("key2"))

	fmt.Println(s.GetKind("key1"))
	fmt.Println(s.GetKind("key2"))
}
