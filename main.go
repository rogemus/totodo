package main

import (
	"fmt"
	"log"
	"os"
	"totodo/cmd"
	"totodo/pkg"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var err error
	command := os.Args[1]
	args := os.Args[2:]

	// Read .env
	if err = godotenv.Load(); err != nil {
		log.Fatalf("[FATAL] %s", err.Error())
	}

	// Init DB (SQL Lite)
	dbConnectionStr := fmt.Sprintf("%s?parseTime=true", os.Getenv("DB_FILE"))
	db, err := pkg.NewDB(dbConnectionStr)

	if err != nil {
		log.Fatalf("[FATAL] %s", err.Error())
	}

	// Init Repository
	tasksRepo := pkg.NewTasksRepository(db)

	// Init commands
	addCmd := cmd.NewAddCmd(tasksRepo)

	switch command {
	case addCmd.Cmd:
		addCmd.Run(args)
	}
}
