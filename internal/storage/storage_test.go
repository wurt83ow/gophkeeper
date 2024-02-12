package storage

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

type mockKeeper struct{}

func (m *mockKeeper) UserExists(ctx context.Context, username string) (bool, error) {
	return true, nil
}

func (m *mockKeeper) AddUser(ctx context.Context, username string, hashedPassword string) error {
	return nil
}

func (m *mockKeeper) GetPassword(ctx context.Context, username string) (string, error) {
	return "hashedPassword", nil
}

func (m *mockKeeper) GetUserID(ctx context.Context, username string) (int, error) {
	return 123, nil
}

func (m *mockKeeper) AddData(ctx context.Context, table string, user_id int, entry_id string, data map[string]string) error {
	return nil
}

func (m *mockKeeper) UpdateData(ctx context.Context, table string, user_id int, entry_id string, data map[string]string) error {
	return nil
}

func (m *mockKeeper) DeleteData(ctx context.Context, table string, user_id int, entry_id string) error {
	return nil
}

func (m *mockKeeper) GetAllData(ctx context.Context, table string, user_id int, last_sync time.Time, incl_del bool) ([]map[string]string, error) {
	return nil, nil
}

type mockLogger struct{}

func (m *mockLogger) Info(string, ...zapcore.Field) {}

func TestMemoryStorage_UserExists(t *testing.T) {
	storage := NewMemoryStorage(&mockKeeper{}, &mockLogger{})
	exists, err := storage.UserExists(context.Background(), "test")
	assert.NoError(t, err)
	assert.True(t, exists)
}

func TestMemoryStorage_AddUser(t *testing.T) {
	storage := NewMemoryStorage(&mockKeeper{}, &mockLogger{})
	err := storage.AddUser(context.Background(), "test", "hashedPassword")
	assert.NoError(t, err)
}

func TestMemoryStorage_GetPassword(t *testing.T) {
	storage := NewMemoryStorage(&mockKeeper{}, &mockLogger{})
	password, err := storage.GetPassword(context.Background(), "test")
	assert.NoError(t, err)
	assert.Equal(t, "hashedPassword", password)
}

func TestMemoryStorage_GetUserID(t *testing.T) {
	storage := NewMemoryStorage(&mockKeeper{}, &mockLogger{})
	userID, err := storage.GetUserID(context.Background(), "test")
	assert.NoError(t, err)
	assert.Equal(t, 123, userID)
}

func TestMemoryStorage_AddData(t *testing.T) {
	storage := NewMemoryStorage(&mockKeeper{}, &mockLogger{})
	err := storage.AddData(context.Background(), "table", 123, "entry", map[string]string{"key": "value"})
	assert.NoError(t, err)
}

func TestMemoryStorage_UpdateData(t *testing.T) {
	storage := NewMemoryStorage(&mockKeeper{}, &mockLogger{})
	err := storage.UpdateData(context.Background(), "table", 123, "entry", map[string]string{"key": "value"})
	assert.NoError(t, err)
}

func TestMemoryStorage_DeleteData(t *testing.T) {
	storage := NewMemoryStorage(&mockKeeper{}, &mockLogger{})
	err := storage.DeleteData(context.Background(), "table", 123, "entry")
	assert.NoError(t, err)
}

func TestMemoryStorage_GetAllData(t *testing.T) {
	storage := NewMemoryStorage(&mockKeeper{}, &mockLogger{})
	data, err := storage.GetAllData(context.Background(), "table", 123, time.Now(), false)
	assert.NoError(t, err)
	assert.Nil(t, data)
}
