package store

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

// InitDB initializes SQLite database and creates necessary tables
func InitDB(dbPath string) (*DB, error) {
	if dbPath == "" {
		dbPath = ":memory:"
	}

	sqlDB, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	db := &DB{sqlDB}

	if err := db.createTables(); err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	log.Println("Database initialized successfully")
	return db, nil
}

func (db *DB) createTables() error {
	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		email TEXT UNIQUE NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS devices (
		id TEXT PRIMARY KEY,
		user_id TEXT NOT NULL,
		device_code TEXT UNIQUE NOT NULL,
		device_name TEXT,
		paired BOOLEAN DEFAULT FALSE,
		code_expires_at DATETIME,
		last_sync DATETIME,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);

	CREATE TABLE IF NOT EXISTS sync_log (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		device_id TEXT NOT NULL,
		file_path TEXT NOT NULL,
		action TEXT NOT NULL,
		content_hash TEXT,
		version_id TEXT UNIQUE,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
		synced BOOLEAN DEFAULT FALSE,
		FOREIGN KEY (device_id) REFERENCES devices(id)
	);

	CREATE TABLE IF NOT EXISTS markdown_metadata (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		file_path TEXT UNIQUE NOT NULL,
		s3_key TEXT,
		content_hash TEXT,
		file_size INTEGER,
		last_modified DATETIME,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS conflicts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		file_path TEXT NOT NULL,
		device_id TEXT NOT NULL,
		version_a TEXT,
		version_b TEXT,
		resolved BOOLEAN DEFAULT FALSE,
		resolution TEXT,
		resolved_at DATETIME,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (device_id) REFERENCES devices(id)
	);
	`

	_, err := db.Exec(schema)
	return err
}

func (db *DB) GetUserByID(userID string) (map[string]interface{}, error) {
	row := db.QueryRow("SELECT id, email, created_at FROM users WHERE id = ?", userID)
	user := make(map[string]interface{})
	err := row.Scan(&user["id"], &user["email"], &user["created_at"])
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}

func (db *DB) CreateUser(userID, email string) error {
	_, err := db.Exec("INSERT INTO users (id, email) VALUES (?, ?)", userID, email)
	return err
}

func (db *DB) GetDeviceByCode(deviceCode string) (map[string]interface{}, error) {
	row := db.QueryRow(`
		SELECT id, user_id, device_code, device_name, paired, code_expires_at, last_sync, created_at
		FROM devices
		WHERE device_code = ?
	`, deviceCode)
	device := make(map[string]interface{})
	err := row.Scan(
		&device["id"], &device["user_id"], &device["device_code"],
		&device["device_name"], &device["paired"], &device["code_expires_at"],
		&device["last_sync"], &device["created_at"],
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return device, err
}

func (db *DB) CreateDevice(deviceID, userID, deviceCode, deviceName string, expiresAt string) error {
	_, err := db.Exec(`
		INSERT INTO devices (id, user_id, device_code, device_name, code_expires_at)
		VALUES (?, ?, ?, ?, ?)
	`, deviceID, userID, deviceCode, deviceName, expiresAt)
	return err
}
