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
	reportCmd := cmd.NewReportCmd(tasksRepo)
	showCmd := cmd.NewShowCmd(tasksRepo)
	startCmd := cmd.NewStartCmd(tasksRepo)
	editCmd := cmd.NewEditCmd(tasksRepo)
	stopCmd := cmd.NewStopCmd(tasksRepo)
	helpCmd := cmd.NewHelpCmd([]pkg.Cmd{
		addCmd,
		editCmd,
		reportCmd,
		showCmd,
		startCmd,
		stopCmd,
	})

	switch command {
	case addCmd.Cmd:
		addCmd.Run(args)
	case reportCmd.Cmd:
		reportCmd.Run(args)
	case showCmd.Cmd:
		showCmd.Run(args)
	case startCmd.Cmd:
		startCmd.Run(args)
	case stopCmd.Cmd:
		stopCmd.Run(args)
	case helpCmd.Cmd:
		helpCmd.Run()
	default:
		// TODO display available cmds
		fmt.Printf("invalid comments")
	}
}
