package storage

import (
	"go.uber.org/zap"
	"fmt"
	"strconv"

)

type Value struct {
	s string
	d int
	a any
	b bool
}

type Storage struct {
	inner  map[string]Value
	logger *zap.Logger
}

func NewStorage(logger *zap.Logger) (Storage, error) {
	logger.Info("created new storage")
	return Storage{
		inner:  make(map[string]Value),
		logger: logger,
	}, nil
}

func (r Storage) Set(key string, value interface{}) {
	r.inner[key] = Value{s: fmt.Sprintf("%v", value)}
	r.logger.Info("key set", zap.String("key", key), zap.String("value", fmt.Sprintf("%v", value)))
}

func (r Storage) Get(key string) *string {
	res, ok := r.inner[key]
	if !ok {
		return nil
	}
	return &res.s
}

func (r Storage) GetKind(key string) string {
	value, ok := r.inner[key]
	if !ok {
		return ""
	}
	if _, err := strconv.Atoi(value.s); err == nil {
		return "D"
	}
	return "S"
}