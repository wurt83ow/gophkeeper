package bdkeeper

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // registers a migrate driver.
	_ "github.com/jackc/pgx/v5/stdlib"                   // registers a pgx driver.
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Log interface {
	Info(string, ...zapcore.Field)
}

type BDKeeper struct {
	conn *sql.DB
	log  Log
}

func NewBDKeeper(dsn func() string, log Log) *BDKeeper {
	addr := dsn()
	if addr == "" {
		log.Info("database dsn is empty")

		return nil
	}

	conn, err := sql.Open("pgx", dsn())
	if err != nil {
		log.Info("Unable to connection to database: ", zap.Error(err))

		return nil
	}

	driver, err := postgres.WithInstance(conn, new(postgres.Config))
	if err != nil {
		log.Info("error getting driver: ", zap.Error(err))

		return nil
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Info("error getting getwd: ", zap.Error(err))
	}

	// fix error test path
	mp := dir + "/migrations"

	var path string
	if _, err := os.Stat(mp); err != nil {
		path = "../../"
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%smigrations", path),
		"postgres",
		driver)
	if err != nil {
		log.Info("Error creating migration instance : ", zap.Error(err))
	}

	err = m.Up()
	if err != nil {
		log.Info("Error while performing migration: ", zap.Error(err))
	}

	log.Info("Connected!")

	return &BDKeeper{
		conn: conn,
		log:  log,
	}
}

// Ping checks the connectivity to the PostgreSQL database and returns true if successful, otherwise false.
func (bdk *BDKeeper) Ping() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := bdk.conn.PingContext(ctx); err != nil {
		return false
	}

	return true
}

// Close closes the connection to the PostgreSQL database and returns true if successful, otherwise false.
func (bdk *BDKeeper) Close() bool {
	bdk.log.Info("Stop database")
	bdk.conn.Close()
	bdk.log.Info("All sql queries are completed")
	return true
}

func (bdk *BDKeeper) UserExists(ctx context.Context, username string) (bool, error) {
	// Запрос для проверки наличия пользователя в базе данных
	query := `SELECT COUNT(*) FROM Users WHERE username = $1;`

	// Выполнение запроса
	row := bdk.conn.QueryRowContext(ctx, query, username)

	// Получение результата
	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	// Если количество записей больше 0, значит пользователь существует
	return count > 0, nil
}

func (bdk *BDKeeper) AddUser(ctx context.Context, username string, hashedPassword string) error {
	// Запрос для добавления нового пользователя в базу данных
	query := `INSERT INTO Users (username, password) VALUES ($1, $2);`

	// Выполнение запроса
	_, err := bdk.conn.ExecContext(ctx, query, username, hashedPassword)
	return err
}

func (bdk *BDKeeper) GetPassword(ctx context.Context, username string) (string, error) {
	// Запрос для получения хешированного пароля пользователя из базы данных
	query := `SELECT password FROM Users WHERE username = $1;`

	// Выполнение запроса
	row := bdk.conn.QueryRowContext(ctx, query, username)

	// Получение результата
	var password string
	err := row.Scan(&password)
	if err != nil {
		return "", err
	}

	// Возвращаем хешированный пароль
	return password, nil
}

func (bdk *BDKeeper) GetUserID(ctx context.Context, username string) (int, error) {
	// Запрос для получения идентификатора пользователя из базы данных
	query := `SELECT id FROM Users WHERE username = $1;`

	// Выполнение запроса
	row := bdk.conn.QueryRowContext(ctx, query, username)

	// Получение результата
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	// Возвращаем идентификатор пользователя
	return id, nil
}

func (bdk *BDKeeper) AddData(ctx context.Context, table string, user_id int, entry_id string, data map[string]string) error {

	keys := make([]string, 0, len(data)+2)        // +2 для user_id и entry_id
	values := make([]interface{}, 0, len(data)+2) // +2 для user_id и entry_id

	// Добавьте user_id и entry_id в начало списков ключей и значений
	keys = append(keys, "user_id", "id")
	values = append(values, user_id, entry_id)

	for key, value := range data {
		keys = append(keys, key)
		values = append(values, value)
	}

	// Создаем подстановочные знаки для значений
	placeholders := make([]string, len(values))
	for i := range values {
		placeholders[i] = "$" + strconv.Itoa(i+1)
	}

	stmt, err := bdk.conn.Prepare(fmt.Sprintf("INSERT INTO %s(%s) values(%s)", table, strings.Join(keys, ","), strings.Join(placeholders, ",")))
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, values...)

	return err
}

func (bdk *BDKeeper) UpdateData(ctx context.Context, table string, user_id int, entry_id string, data map[string]string) error {
	setClauses := make([]string, 0, len(data))
	values := make([]interface{}, 0, len(data)+2) // +2 для user_id и id

	i := 1
	for key, value := range data {
		setClauses = append(setClauses, key+" = $"+strconv.Itoa(i))
		values = append(values, value)
		i++
	}

	// Добавьте user_id и id в конец списка значений
	values = append(values, user_id, entry_id)

	stmt, err := bdk.conn.Prepare(fmt.Sprintf("UPDATE %s SET %s WHERE user_id = $%d AND id = $%d", table, strings.Join(setClauses, ","), i, i+1))
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, values...)
	return err
}

func (bdk *BDKeeper) DeleteData(ctx context.Context, table string, user_id int, entry_id string) error {
	// Check user_id and table
	if user_id == 0 || table == "" {
		return errors.New("user_id and table must be specified")
	}

	// Check id
	if entry_id == "" {
		return errors.New("id must be specified")
	}

	// Prepare the query
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE user_id = $1 AND id = $2", table)
	args := []interface{}{user_id, entry_id}

	// Execute the query
	row := bdk.conn.QueryRowContext(ctx, query, args...)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return err
	}

	// Check the number of records found
	if count > 1 {
		return errors.New("More than one record found")
	} else if count == 0 {
		return errors.New("No records found")
	}

	// Delete the record
	query = strings.Replace(query, "SELECT COUNT(*)", "DELETE", 1)
	_, err = bdk.conn.ExecContext(ctx, query, args...)
	return err
}

func (bdk *BDKeeper) GetAllData(ctx context.Context, table string, user_id int) ([]map[string]string, error) {
	// Получаем все колонки таблицы
	rows, err := bdk.conn.QueryContext(ctx, fmt.Sprintf(`SELECT column_name FROM information_schema.columns WHERE table_name = '%s'`, strings.ToLower(table)))
	if err != nil {
		return nil, fmt.Errorf("failed to get columns: %w", err)
	}

	var cols []string
	for rows.Next() {
		var col string
		err := rows.Scan(&col)
		if err != nil {
			return nil, fmt.Errorf("failed to scan column: %w", err)
		}

		cols = append(cols, col)
	}

	// Запрашиваем все данные из таблицы для данного user_id
	rows, err = bdk.conn.QueryContext(ctx, fmt.Sprintf("SELECT %s FROM %s WHERE user_id = %d", strings.Join(cols, ","), table, user_id))
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	values := make([]interface{}, len(cols))

	for i := range values {
		values[i] = new(sql.NullString)
	}

	var data []map[string]string
	for rows.Next() {
		err := rows.Scan(values...)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		row := make(map[string]string)
		for i, column := range cols {
			if ns, ok := values[i].(*sql.NullString); ok {
				row[column] = ns.String

			}
		}

		data = append(data, row)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows encountered an error: %w", err)
	}

	return data, nil
}

func (bdk *BDKeeper) ClearData(ctx context.Context, table string, userID int) error {
	stmt, err := bdk.conn.Prepare(fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", table))
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, userID)
	return err
}
