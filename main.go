package main

import (
	"github.com/atotto/clipboard"
	"github.com/markoczy/gotext/handler"
	"github.com/markoczy/goutil/log"
	"github.com/markoczy/goutil/log/logconfig"
	"os"
)

func main() {
	logconfig.SetDefaultLogLevel(logconfig.ERROR)

	var clip *string
	read, err := clipboard.ReadAll()
	if err != nil {
		log.Debug("Clipboard not available: " + err.Error())
	} else {
		clip = &read
	}

	s, err := handler.Exec(os.Args[1:], clip)
	if err != nil {
		log.Error("Something went wrong: " + err.Error())
		return
	}

	if s != nil {
		log.Debug("New Value found, writing")
		clipboard.WriteAll(*s)
	}

	log.Debug("Normal Exit")
}
