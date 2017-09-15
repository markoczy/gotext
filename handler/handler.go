package handler

import 
(
	"log"
	"strings"
	"github.com/markoczy/goutil/cli"
	"github.com/markoczy/goutil/cli/parser"
)

var cmdParser = initParser()

const OK_NO_VALUE = "OK_NO_VALUE"

func Exec(args []string, input string) (*string,error) {
	log.Printf("Command array: %v\n",args)
	log.Printf("Input: %s\n",input)
	
	if len(args)<1 {
		showHelp()
		// not set clipboard
		return nil, nil
	}

	lArgs := append(args,input)
	ifc, err := parser.Exec(&cmdParser, lArgs)
	ret := ifc.(string)
	return &ret, err
}

func showHelp() {

}

func initParser() parser.Parser {
	parser := cli.NewParser();

	cli.AddCommand(parser,"Uppercase","u",1,uppercase)
	
	return *parser
}

func uppercase(s []string) (interface{}, error) {
	return strings.ToUpper(s[1]), nil
}