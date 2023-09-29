package database

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
)

// define errors
var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row does not exist")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

// SQLiteRepository for interacting with the DB

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

// migrate: create if not exists
// integer unixTimestamp represents time in unix format
func (r *SQLiteRepository) Migrate() error {
	query := `
    CREATE TABLE IF NOT EXISTS files(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        filename TEXT UNIQUE NOT NULL,
		mime TEXT,
		description TEXT,
		uploader TEXT,
		unixTimestamp INTEGER
    );
    `

	_, err := r.db.Exec(query)
	return err
}

// create: insert row from Metadata
func (r *SQLiteRepository) Create(file Metadata) (*Metadata, error) {
	res, err := r.db.Exec("INSERT INTO files(filename, mime, description, uploader, unixTimestamp) values(?,?,?,?,?)", file.Filename, file.Mime, file.Description, file.Uploader, file.UnixTimestamp)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	file.ID = id

	return &file, nil
}

// all: get all rows
func (r *SQLiteRepository) All() ([]Metadata, error) {
	rows, err := r.db.Query("SELECT * FROM files")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []Metadata
	for rows.Next() {
		var file Metadata
		if err := rows.Scan(&file.ID, &file.Filename, &file.Mime, &file.Description, &file.Uploader, &file.UnixTimestamp); err != nil {
			return nil, err
		}
		all = append(all, file)
	}
	return all, nil
}

// get the row of selected (unique) id
func (r *SQLiteRepository) GetById(id int64) (*Metadata, error) {
	row := r.db.QueryRow("SELECT * FROM files WHERE id = ?", id)

	var file Metadata
	if err := row.Scan(&file.ID, &file.Filename, &file.Mime, &file.Description, &file.Uploader, &file.UnixTimestamp); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &file, nil
}

// delete row of (unique) id
func (r *SQLiteRepository) Delete(id int64) error {
	res, err := r.db.Exec("DELETE FROM files WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrDeleteFailed
	}

	return err
}
