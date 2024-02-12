package bdkeeper

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/wurt83ow/gophkeeper-server/internal/config"
	"github.com/wurt83ow/gophkeeper-server/internal/logger"
)

// Функция для создания экземпляра BDKeeper с помощью NewBDKeeper
func newTestBDKeeper(t *testing.T, db *sql.DB) *BDKeeper {
	// Создание и инициализация нового экземпляра опций
	option := config.NewOptions()
	option.ParseFlags()

	// Получение нового логгера
	nLogger, err := logger.NewLogger(option.LogLevel())
	if err != nil {
		t.Fatalf("Error creating logger: %v", err)
	}

	// Создание экземпляра BDKeeper через конструктор
	bdk, err := NewBDKeeper(func() string { return "" }, nLogger, db)
	if err != nil {
		t.Fatalf("Error creating BDKeeper: %v", err)
	}

	return bdk
}

func TestBDKeeper_Ping(t *testing.T) {
	// Инициализация sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock database: %v", err)
	}
	defer db.Close()

	// Создание экземпляра BDKeeper через функцию newTestBDKeeper
	bdk := newTestBDKeeper(t, db)

	// Устанавливаем ожидание вызова метода Ping
	mock.ExpectPing()

	// Вызываем метод Ping
	if !bdk.Ping() {
		t.Error("Ping failed")
	}

	// Проверяем, что все ожидания выполнены
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestBDKeeper_Close(t *testing.T) {
	// Инициализация sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock database: %v", err)
	}
	defer db.Close()

	// Создание экземпляра BDKeeper через функцию newTestBDKeeper
	bdk := newTestBDKeeper(t, db)

	// Устанавливаем ожидание вызова метода Close
	mock.ExpectClose()

	// Вызываем метод Close
	if !bdk.Close() {
		t.Error("Close failed")
	}

	// Проверяем, что все ожидания выполнены
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestBDKeeper_UserExists(t *testing.T) {
	// Инициализация sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock database: %v", err)
	}
	defer db.Close()

	// Создание экземпляра BDKeeper через функцию newTestBDKeeper
	bdk := newTestBDKeeper(t, db)

	// Ожидание вызова QueryRowContext с запросом для проверки существования пользователя
	mock.ExpectQuery("SELECT COUNT(.+) FROM Users WHERE username = (.+)").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	// Проверка существования пользователя
	exists, err := bdk.UserExists(context.Background(), "testUser")
	if err != nil {
		t.Fatalf("Error checking user existence: %v", err)
	}

	// Проверка результата
	if !exists {
		t.Error("Expected user to exist")
	}

	// Проверяем, что все ожидания выполнены
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestBDKeeper_AddUser(t *testing.T) {
	// Инициализация sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock database: %v", err)
	}
	defer db.Close()

	// Создание экземпляра BDKeeper через функцию newTestBDKeeper
	bdk := newTestBDKeeper(t, db)

	// Ожидание вызова ExecContext для добавления пользователя
	mock.ExpectExec("INSERT INTO Users (.+) VALUES (.+)").WillReturnResult(sqlmock.NewResult(1, 1))

	// Добавление нового пользователя
	err = bdk.AddUser(context.Background(), "testUser", "hashedPassword")
	if err != nil {
		t.Fatalf("Error adding user: %v", err)
	}

	// Проверяем, что все ожидания выполнены
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestBDKeeper_GetPassword(t *testing.T) {
	// Инициализация sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock database: %v", err)
	}
	defer db.Close()

	// Создание экземпляра BDKeeper через функцию newTestBDKeeper
	bdk := newTestBDKeeper(t, db)

	// Ожидание вызова QueryRowContext для получения пароля пользователя
	mock.ExpectQuery("SELECT password FROM Users WHERE username = (.+)").WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow("hashedPassword"))

	// Получение пароля пользователя
	password, err := bdk.GetPassword(context.Background(), "testUser")
	if err != nil {
		t.Fatalf("Error getting password: %v", err)
	}

	// Проверка полученного пароля
	if password != "hashedPassword" {
		t.Errorf("Expected password %s, got %s", "hashedPassword", password)
	}

	// Проверяем, что все ожидания выполнены
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
func TestBDKeeper_GetUserID(t *testing.T) {
	// Инициализация sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock database: %v", err)
	}
	defer db.Close()

	// Создание экземпляра BDKeeper через функцию newTestBDKeeper
	bdk := newTestBDKeeper(t, db)

	// Ожидание вызова QueryRowContext для получения ID пользователя
	mock.ExpectQuery("SELECT id FROM Users WHERE username = (.+)").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	// Получение ID пользователя
	userID, err := bdk.GetUserID(context.Background(), "testUser")
	if err != nil {
		t.Fatalf("Error getting user ID: %v", err)
	}

	// Проверка полученного ID
	if userID != 1 {
		t.Errorf("Expected user ID %d, got %d", 1, userID)
	}

	// Проверяем, что все ожидания выполнены
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestBDKeeper_AddData(t *testing.T) {
	// Инициализация sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock database: %v", err)
	}
	defer db.Close()

	// Создание экземпляра BDKeeper через функцию newTestBDKeeper
	bdk := newTestBDKeeper(t, db)
	// Ожидание вызова Prepare
	mock.ExpectPrepare("INSERT INTO testTable(.+) VALUES(.+)")

	// Ожидание вызова ExecContext для добавления данных
	mock.ExpectExec("INSERT INTO testTable(.+) VALUES(.+)").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Добавление новых данных
	err = bdk.AddData(context.Background(), "testTable", 1, "entry_id", map[string]string{"key1": "value1", "key2": "value2"})
	if err != nil {
		t.Fatalf("Ошибка при добавлении данных: %v", err)
	}

	// Проверяем, что все ожидания выполнены
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Не выполнены ожидания: %s", err)
	}
}

func TestBDKeeper_UpdateData(t *testing.T) {
	// Инициализация sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock database: %v", err)
	}
	defer db.Close()

	// Создание экземпляра BDKeeper через функцию newTestBDKeeper
	bdk := newTestBDKeeper(t, db)

	// Ожидание вызова Prepare
	mock.ExpectPrepare("UPDATE testTable SET(.+) WHERE user_id = (.+) AND id = (.+)")

	// Ожидание вызова ExecContext для обновления данных
	mock.ExpectExec("UPDATE testTable SET(.+) WHERE user_id = (.+) AND id = (.+)").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Обновление данных
	err = bdk.UpdateData(context.Background(), "testTable", 1, "entryID", map[string]string{"key1": "value1", "key2": "value2"})
	if err != nil {
		t.Fatalf("Ошибка при обновлении данных: %v", err)
	}

	// Проверяем, что все ожидания выполнены
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Не выполнены ожидания: %s", err)
	}
}

func TestBDKeeper_DeleteData(t *testing.T) {
	// Инициализация sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock database: %v", err)
	}
	defer db.Close()

	// Создание экземпляра BDKeeper через функцию newTestBDKeeper
	bdk := newTestBDKeeper(t, db)

	// Не ожидаем вызов Query
	mock.ExpectQuery("SELECT COUNT(.+) FROM testTable WHERE user_id = (.+) AND id = (.+)").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	// Ожидание вызова ExecContext для удаления данных
	mock.ExpectExec("DELETE FROM testTable WHERE user_id = (.+) AND id = (.+)").
		WithArgs(1, "entryID").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Удаление данных
	err = bdk.DeleteData(context.Background(), "testTable", 1, "entryID")
	if err != nil {
		t.Fatalf("Ошибка при удалении данных: %v", err)
	}

	// Проверяем, что все ожидания выполнены
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Не выполнены ожидания: %s", err)
	}
}
