package pkg

import (
	"database/sql"
)

const createDb string = `
  CREATE TABLE IF NOT EXISTS tasks (
    id            INTEGER NOT NULL PRIMARY KEY,
    created       DATETIME DEFAULT CURRENT_TIMESTAMP,
    description   TEXT NOT NULL,
    status        TEXT NOT NULL
  );

  INSERT INTO tasks (id, description, status) VALUES 
    (1, "K2-11: component", "todo"),
    (2, "K2-47: comp v2", "todo"),
    (3, "K2-64: bug comp", "done"),
    (4, "K2-03: big bug", "todo"),
    (5, "K2-73: error bug", "done"),
    (6, "K2-11: bug", "active");
`

func NewDB(dbFile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbFile)

	if err != nil {
		return nil, err
	}

	// if _, err := db.Exec(createDb); err != nil {
	// 	return nil, err
	// }

	return db, nil
}
