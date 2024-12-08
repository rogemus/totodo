package pkg

import (
	"database/sql"
)

const createDb string = `
  CREATE TABLE IF NOT EXISTS tasks (
    id            INTEGER NOT NULL PRIMARY KEY,
    created       DATETIME DEFAULT CURRENT_TIMESTAMP,
    description   TEXT NOT NULL
  );

  CREATE TABLE IF NOT EXISTS activeTask (
    id                        INTEGER NOT NULL PRIMARY KEY,
    taskid                    INTEGER,
    trackingid                INTEGER,
    created                   DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(taskid)       REFERENCES tasks(id),
    FOREIGN KEY(trackingid)   REFERENCES tracking(id)
  );

  CREATE TABLE IF NOT EXISTS tracking (
    id                    INTEGER NOT NULL PRIMARY KEY,
    start_time            DATETIME,
    end_time              DATETIME,
    taskid                INTEGER,
    created               DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(taskid)   REFERENCES tasks(id)
  );

  INSERT INTO tasks (id, description) VALUES 
    (1, "K2-11: component"),
    (2, "K2-47: comp v2"),
    (3, "K2-4703: bug");
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
