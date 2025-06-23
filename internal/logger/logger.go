package logger

import (
	"log"
	"os"
)

var (
	Info *log.Logger
	Warn *log.Logger
	Error *log.Logger
)

func init() {
	flags := log.LstdFlags | log.Lshortfile

	Info = log.New(os.Stdout, "INFO: ", flags)
	Warn = log.New(os.Stdout, "WARN: ", flags)
	Error = log.New(os.Stdout, "ERROR: ", flags)
}