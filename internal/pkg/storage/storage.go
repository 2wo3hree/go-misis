package storage

import (
	"go.uber.org/zap"
	"strconv"
)

type Storage struct {
	inner  map[string]string
	logger *zap.Logger
}

func NewStorage(logger *zap.Logger) (Storage, error) {
	logger.Info("created new storage")
	return Storage{
		inner:  make(map[string]string),
		logger: logger,
	}, nil
}

func (r Storage) Set(key, value string) {
	r.inner[key] = value
	r.logger.Info("key set", zap.String("key", key), zap.String("value", value))
}

func (r Storage) Get(key string) *string {
	res, ok := r.inner[key]
	if !ok {
		return nil
	}
	return &res
}

func (r Storage) GetKind(key string) string {
	value, ok := r.inner[key]
	if !ok {
		return ""
	}
	if _, err := strconv.Atoi(value); err == nil {
		return "D"
	}
	return "S"
}
