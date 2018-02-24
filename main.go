package main

import
(
	"os"
	"github.com/atotto/clipboard"
	"github.com/markoczy/goutil/log"
	"github.com/markoczy/gotext/handler"
)

func main() {
	var clip, err = clipboard.ReadAll();
	if err != nil {
		log.Error("Something went wrong: "+err.Error())
		return
	}
	
	s, err := handler.Exec(os.Args[1:], clip)
	if err != nil {
		log.Error("Something went wrong: "+err.Error())
		return
	}

	if s != nil {
		log.Debug("New Value found, writing")
		clipboard.WriteAll(*s) 
	}

	log.Debug("Nomal Exit")
}