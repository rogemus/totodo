package pkg

import (
	"database/sql"
)

const createDb string = `
  CREATE TABLE IF NOT EXISTS [tasks] (
    id            INTEGER NOT NULL PRIMARY KEY,
    created       DATETIME DEFAULT CURRENT_TIMESTAMP,
    description   TEXT NOT NULL
  );

  INSERT INTO "tasks" (description) VALUES 
    ("K2-11: component"),
    ("K2-47: comp v2"),
    ("K2-4703: bug");
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
