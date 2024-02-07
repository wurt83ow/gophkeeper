package storage

import (
	"context"
	"errors"

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

type Keeper interface {
	UserExists(ctx context.Context, username string) (bool, error)
	AddUser(ctx context.Context, username string, hashedPassword string) error
	GetPassword(ctx context.Context, username string) (string, error)
	GetUserID(ctx context.Context, username string) (int, error)
	AddData(ctx context.Context, table string, user_id int, entry_id string, data map[string]string) error
	UpdateData(ctx context.Context, table string, user_id int, entry_id string, data map[string]string) error
	DeleteData(ctx context.Context, table string, user_id int, entry_id string) error
	GetData(ctx context.Context, table string, user_id int, entry_id string) (map[string]string, error)
	GetAllData(ctx context.Context, table string, columns ...string) ([]map[string]string, error)
	ClearData(ctx context.Context, table string, userID int) error
}

// NewMemoryStorage creates a new MemoryStorage instance with the provided Keeper and logger.
func NewMemoryStorage(keeper Keeper, log Log) *MemoryStorage {
	return &MemoryStorage{
		keeper: keeper,
		log:    log,
	}
}

func (ms *MemoryStorage) UserExists(ctx context.Context, username string) (bool, error) {
	return ms.keeper.UserExists(ctx, username)
}

func (ms *MemoryStorage) AddUser(ctx context.Context, username string, hashedPassword string) error {
	return ms.keeper.AddUser(ctx, username, hashedPassword)
}

func (ms *MemoryStorage) GetPassword(ctx context.Context, username string) (string, error) {
	return ms.keeper.GetPassword(ctx, username)
}

func (ms *MemoryStorage) GetUserID(ctx context.Context, username string) (int, error) {
	return ms.keeper.GetUserID(ctx, username)
}

func (ms *MemoryStorage) AddData(ctx context.Context, table string, user_id int, entry_id string, data map[string]string) error {
	return ms.keeper.AddData(ctx, table, user_id, entry_id, data)
}

func (ms *MemoryStorage) UpdateData(ctx context.Context, table string, user_id int, entry_id string, data map[string]string) error {
	return ms.keeper.UpdateData(ctx, table, user_id, entry_id, data)
}

func (ms *MemoryStorage) DeleteData(ctx context.Context, table string, user_id int, entry_id string) error {
	return ms.keeper.DeleteData(ctx, table, user_id, entry_id)
}

func (ms *MemoryStorage) GetData(ctx context.Context, table string, user_id int, entry_id string) (map[string]string, error) {
	return ms.keeper.GetData(ctx, table, user_id, entry_id)
}

func (ms *MemoryStorage) GetAllData(ctx context.Context, table string, columns ...string) ([]map[string]string, error) {
	return ms.keeper.GetAllData(ctx, table, columns...)
}

func (ms *MemoryStorage) ClearData(ctx context.Context, userID int, table string) error {
	return ms.keeper.ClearData(ctx, table, userID)
}
