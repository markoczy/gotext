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

	cli.AddCommand(parser, "Uppercase", 1, "^((u)|(uc)|(upper)|(uppercase))$", 1, uppercase)
	cli.AddCommand(parser, "Lowercase", 1, "^((l)|(lc)|(lower)|(lowercase))$", 1, lowercase)
	cli.AddCommand(parser, "Prefix", 1, "^((p)|(pr)|(pre)|(prefix))$", 2, prefix)
	cli.AddCommand(parser, "Suffix", 1, "^((s)|(po)|(post)|(suffix))$", 2, suffix)
	cli.AddCommand(parser, "Trim start", 1, "^((ts)|(tstart)|(trimstart))$", 2, trimStart)
	cli.AddCommand(parser, "Trim start (exclusive)", 2, "^((tsx)|(tstartx)|(trimstartx))$", 1, trimStartX)
	cli.AddCommand(parser, "Trim end", 1, "^((te)|(tend)|(trimend))$", 2, trimEnd)
	cli.AddCommand(parser, "Trim end (exclusive)", 2, "^((tex)|(tendx)|(trimendx))$", 1, trimEndX)

	return parser
}

func uppercase(s []string) (interface{}, error) {
	log.Debug("Entry uppercase")
	return strings.ToUpper(s[1]), nil
}

func lowercase(s []string) (interface{}, error) {
	log.Debug("Entry lowercase")
	return strings.ToLower(s[1]), nil
}

func prefix(s []string) (interface{}, error) {
	log.Debug("Entry prefix")
	var ret string
	split := xNewLine.Split(s[2], -1)
	for _, e := range split {
		ret += s[1] + e + "\n"
	}

	return ret, nil
}

// tt s abc
func suffix(s []string) (interface{}, error) {
	log.Debug("Entry suffix")
	var ret string
	split := xNewLine.Split(s[2], -1)
	for _, e := range split {
		ret += e + s[1] + "\n"
	}

	return ret, nil
}

// tt ts abc
func trimStart(s []string) (interface{}, error) {
	log.Debug("Entry trimStart")
	var size = len(s[1])
	var ret string
	split := xNewLine.Split(s[2], -1)
	for _, e := range split {
		// log.Debugf("idx: %v", idx)
		var idx = strings.Index(e, s[1])
		if idx > -1 {
			ret += e[idx+size:] + "\n"
		} else {
			ret += e + "\n"
		}
	}

	return ret, nil
}

// tt tsx abc
func trimStartX(s []string) (interface{}, error) {
	log.Debug("Entry trimStartX")
	var ret string
	split := xNewLine.Split(s[2], -1)
	for _, e := range split {
		// log.Debugf("idx: %v", idx)
		var idx = strings.Index(e, s[1])
		if idx > -1 {
			ret += e[idx:] + "\n"
		} else {
			ret += e + "\n"
		}
	}

	return ret, nil

}

func trimEnd(s []string) (interface{}, error) {
	log.Debug("Entry trimEnd")
	var ret string
	split := xNewLine.Split(s[2], -1)
	for _, e := range split {
		// log.Debugf("idx: %v", idx)
		var idx = strings.Index(e, s[1])
		if idx > -1 {
			ret += e[:idx] + "\n"
		} else {
			ret += e + "\n"
		}
	}

	return ret, nil
}

func trimEndX(s []string) (interface{}, error) {
	log.Debug("Entry trimEndX")
	var size = len(s[1])
	var ret string
	split := xNewLine.Split(s[2], -1)
	for _, e := range split {
		// log.Debugf("idx: %v", idx)
		var idx = strings.Index(e, s[1])
		if idx > -1 {
			ret += e[:idx-size+1] + "\n"
		} else {
			ret += e + "\n"
		}
	}

	return ret, nil
}

// todo: replace, filter, merge, ...

// func isValidCount(s []string, req int) {
// 	if req == -1 {
// 		return true
// 	}
// 	return len(s) == req
// }
