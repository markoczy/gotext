package main

import
(
	"os"
	"log"
	"github.com/atotto/clipboard"
	"github.com/markoczy/gotext/handler"
)

func main() {
	var clip, err = clipboard.ReadAll();
	if err != nil {
		log.Panicln("Something went wrong: "+err.Error())
		return
	}
	
	s, err := handler.Exec(os.Args[1:], clip)
	if err != nil {
		log.Panicln("Something went wrong: "+err.Error())
		return
	}

	if s != nil {
		log.Println("New Value found, writing")
		clipboard.WriteAll(*s) 
	}

	log.Println("Nomal Exit")
}