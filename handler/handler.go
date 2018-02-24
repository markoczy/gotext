package handler

import (
	"regexp"
	"strings"

	"github.com/markoczy/goutil/cli"
	"github.com/markoczy/goutil/cli/parser"
	"github.com/markoczy/goutil/log"
)

var cmdParser = initParser()

var xNewLine = regexp.MustCompile("\r?\n")

const StrOkNoValue = "OK_NO_VALUE"

func Exec(args []string, input string) (*string, error) {
	log.Debugf("Command array: %v\n", args)
	log.Debugf("Input: %s\n", input)

	if len(args) < 1 {
		showHelp()
		// not set clipboard
		return nil, nil
	}

	lArgs := append(args, input)
	ifc, err := cmdParser.Exec(lArgs)
	ret := ifc.(string)
	return &ret, err
}

func showHelp() {

}

// func AddCommand(aParser parser.Parser, aName string, aPriority int,
// 	aRegex string, aArgCount int, aOperation command.Operation) error {

func initParser() parser.Parser {
	parser := cli.NewParser()

	cli.AddCommand(parser, "Uppercase", 1, "u", 1, uppercase)
	cli.AddCommand(parser, "Lowercase", 1, "l", 1, lowercase)
	cli.AddCommand(parser, "Prefix", 1, "p", 2, prefix)
	cli.AddCommand(parser, "Suffix", 1, "s", 2, suffix)
	cli.AddCommand(parser, "Trim start", 1, "ts", 2, trimStart)
	cli.AddCommand(parser, "Trim start (exclusive)", 2, "tsx", 1, trimStartX)
	cli.AddCommand(parser, "Trim end", 1, "te", 2, trimEnd)
	cli.AddCommand(parser, "Trim end (exclusive)", 2, "tex", 1, trimEndX)

	return parser
}

func uppercase(s []string) (interface{}, error) {
	return strings.ToUpper(s[1]), nil
}

func lowercase(s []string) (interface{}, error) {
	return strings.ToLower(s[1]), nil
}

func prefix(s []string) (interface{}, error) {

	// split := strings.Split(s[2], "\n")
	var ret string
	split := xNewLine.Split(s[2], -1)
	for _, e := range split {
		log.Debug("Element: " + e)
		ret += s[1] + e + "\n"
	}

	return ret, nil
}

func suffix(s []string) (interface{}, error) {
	return strings.ToLower(s[1]), nil
}

func trimStart(s []string) (interface{}, error) {
	return strings.ToLower(s[1]), nil
}

func trimStartX(s []string) (interface{}, error) {
	return strings.ToLower(s[1]), nil
}

func trimEnd(s []string) (interface{}, error) {
	return strings.ToLower(s[1]), nil
}

func trimEndX(s []string) (interface{}, error) {
	return strings.ToLower(s[1]), nil
}

// todo: replace, filter, merge, ...

// func isValidCount(s []string, req int) {
// 	if req == -1 {
// 		return true
// 	}
// 	return len(s) == req
// }
