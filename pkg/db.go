package pkg

import (
	"database/sql"
)

const createDb string = `
  CREATE TABLE IF NOT EXISTS [tasks] (
    id INTEGER NOT NULL PRIMARY KEY,
    time DATETIME NOT NULL,
    description TEXT
  );
`

func NewDB(dbFile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbFile)

	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(createDb); err != nil {
		return nil, err
	}

	return db, nil
}
