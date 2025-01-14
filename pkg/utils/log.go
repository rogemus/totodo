package utils

import "log"

type logger struct{}

func (l logger) Success(msg string) {
	log.Printf("[SUCCESS] %s", msg)
}

func (l logger) Info(msg string) {
	log.Printf("[INFO] %s", msg)
}

func (l logger) Error(err error) {
	log.Printf("[ERROR] %s", err.Error())
}

func (l logger) Fatal(msg string) {
	log.Fatalf("[FATAL] %s", msg)
}

var Log = logger{}
