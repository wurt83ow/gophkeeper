// Package storage provides an in-memory storage implementation with CRUD operations for URL and user data.
// It includes interfaces and a MemoryStorage type implementing these interfaces.
package storage

import (
	"errors"
	"sync"

	"go.uber.org/zap/zapcore"
)

// ErrConflict indicates a data conflict in the store.
var ErrConflict = errors.New("data conflict")

// Log is an interface representing a logger with Info method.
type Log interface {
	Info(string, ...zapcore.Field)
}

// MemoryStorage is an in-memory storage implementation with CRUD operations for URL and user data.
type MemoryStorage struct {
	keeper Keeper
	log    Log
	dmx    sync.RWMutex
	umx    sync.RWMutex
}

// NewMemoryStorage creates a new MemoryStorage instance with the provided Keeper and logger.
func NewMemoryStorage(keeper Keeper, log Log) *MemoryStorage {

	return &MemoryStorage{

		keeper: keeper,
		log:    log,
	}
}
