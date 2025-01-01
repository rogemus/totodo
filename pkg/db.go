package pkg

import (
	"database/sql"
)

const createDb string = `
  CREATE TABLE IF NOT EXISTS tasks (
    id            INTEGER NOT NULL PRIMARY KEY,
    created       DATETIME DEFAULT CURRENT_TIMESTAMP,
    description   TEXT NOT NULL,
    status        TEXT NOT NULL,
    listId        INTEGER NOT NULL,
    FOREIGN KEY(listId) REFERENCES lists(listId)
  );

  CREATE TABLE IF NOT EXISTS lists (
    id        INTEGER NOT NULL PRIMARY KEY,
    created   DATETIME DEFAULT CURRENT_TIMESTAMP,
    name      TEXT NOT NULL
  );

  INSERT INTO lists (id, name) VALUES 
    (1, "tasks"),
    (2, "coding");

  INSERT INTO tasks (id, description, status, listId) VALUES 
    (1, "K2-11: component", "todo", 2),
    (2, "K2-47: comp v2", "todo", 2),
    (3, "K2-64: bug comp", "done", 2),
    (4, "K2-03: big bug", "todo", 2),
    (5, "K2-73: error bug", "done", 1),
    (6, "K2-11: bug", "active", 1);
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
