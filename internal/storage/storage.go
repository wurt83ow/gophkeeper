// Package storage provides implementations of storage interfaces.
package storage

import (
	"context"
	"errors"
	"time"

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
}

// Keeper represents the storage keeper interface.
type Keeper interface {
	// UserExists checks if a user exists.
	UserExists(ctx context.Context, username string) (bool, error)
	// AddUser adds a new user to the storage.
	AddUser(ctx context.Context, username string, hashedPassword string) error
	// GetPassword retrieves the password for the given username.
	GetPassword(ctx context.Context, username string) (string, error)
	// GetUserID retrieves the user ID for the given username.
	GetUserID(ctx context.Context, username string) (int, error)
	// AddData adds data to the storage.
	AddData(ctx context.Context, table string, user_id int, entry_id string, data map[string]string) error
	// UpdateData updates existing data in the storage.
	UpdateData(ctx context.Context, table string, user_id int, entry_id string, data map[string]string) error
	// DeleteData deletes data from the storage.
	DeleteData(ctx context.Context, table string, user_id int, entry_id string) error
	// GetAllData retrieves all data from the storage.
	GetAllData(ctx context.Context, table string, user_id int, last_sync time.Time, incl_del bool) ([]map[string]string, error)
	// ClearData clears data from the storage.
	ClearData(ctx context.Context, table string, userID int) error
}

// NewMemoryStorage creates a new MemoryStorage instance with the provided Keeper and logger.
func NewMemoryStorage(keeper Keeper, log Log) *MemoryStorage {
	return &MemoryStorage{
		keeper: keeper,
		log:    log,
	}
}

// UserExists checks if a user exists.
func (ms *MemoryStorage) UserExists(ctx context.Context, username string) (bool, error) {
	return ms.keeper.UserExists(ctx, username)
}

// AddUser adds a new user to the storage.
func (ms *MemoryStorage) AddUser(ctx context.Context, username string, hashedPassword string) error {
	return ms.keeper.AddUser(ctx, username, hashedPassword)
}

// GetPassword retrieves the password for the given username.
func (ms *MemoryStorage) GetPassword(ctx context.Context, username string) (string, error) {
	return ms.keeper.GetPassword(ctx, username)
}

// GetUserID retrieves the user ID for the given username.
func (ms *MemoryStorage) GetUserID(ctx context.Context, username string) (int, error) {
	return ms.keeper.GetUserID(ctx, username)
}

// AddData adds data to the storage.
func (ms *MemoryStorage) AddData(ctx context.Context, table string, user_id int, entry_id string, data map[string]string) error {
	return ms.keeper.AddData(ctx, table, user_id, entry_id, data)
}

// UpdateData updates existing data in the storage.
func (ms *MemoryStorage) UpdateData(ctx context.Context, table string, user_id int, entry_id string, data map[string]string) error {
	return ms.keeper.UpdateData(ctx, table, user_id, entry_id, data)
}

// DeleteData deletes data from the storage.
func (ms *MemoryStorage) DeleteData(ctx context.Context, table string, user_id int, entry_id string) error {
	return ms.keeper.DeleteData(ctx, table, user_id, entry_id)
}

// GetAllData retrieves all data from the storage.
func (ms *MemoryStorage) GetAllData(ctx context.Context, table string, user_id int, last_sync time.Time, incl_del bool) ([]map[string]string, error) {
	return ms.keeper.GetAllData(ctx, table, user_id, last_sync, incl_del)
}

// ClearData clears data from the storage.
func (ms *MemoryStorage) ClearData(ctx context.Context, userID int, table string) error {
	return ms.keeper.ClearData(ctx, table, userID)
}
