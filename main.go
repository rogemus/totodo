package main

import (
	"fmt"
	"log"
	"os"
	"totodo/cmd"
	"totodo/pkg"
	"totodo/pkg/model"
	repo "totodo/pkg/repository"

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

	// Init Repositories
	tasksRepo := repo.NewTasksRepository(db)

	// Init commands
	addCmd := cmd.NewAddCmd(tasksRepo)
	reportCmd := cmd.NewReportCmd(tasksRepo)
	showCmd := cmd.NewShowCmd(tasksRepo)
	editCmd := cmd.NewEditCmd(tasksRepo)
	helpCmd := cmd.NewHelpCmd([]model.Cmd{
		addCmd,
		editCmd,
		reportCmd,
		showCmd,
	})

	switch command {
	case addCmd.Cmd:
		addCmd.Run(args)
	case editCmd.Cmd:
		editCmd.Run(args)
	case reportCmd.Cmd:
		reportCmd.Run(args)
	case showCmd.Cmd:
		showCmd.Run(args)
	case helpCmd.Cmd:
		helpCmd.Run()
	default:
		// TODO display available cmds
		fmt.Println("invalid comments")
	}
}
