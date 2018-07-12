package handler

import (
	"fmt"
	"github.com/markoczy/goutil/cli"
	"github.com/markoczy/goutil/cli/parser"
	"github.com/markoczy/goutil/log"
)

var cmdParser = initParser()

// Exec ...
func Exec(args []string, input *string) (*string, error) {
	log.Debugf("Command array: %v\n", args)
	log.Debugf("Input: %s\n", input)

	if len(args) < 1 {
		showHelp()
		// not set clipboard
		return nil, nil
	}

	var lArgs []string
	if input != nil {
		lArgs = append(args, *input)
	}
	ifc, err := cmdParser.Exec(lArgs)
	if err != nil {
		return nil, err
	}
	ret := ifc.(string)

	return &ret, err
}

// func AddCommand(aParser parser.Parser, aName string, aPriority int,
// 	aRegex string, aArgCount int, aOperation command.Operation) error {

func initParser() parser.Parser {
	parser := cli.NewParser()

	cli.AddCommand(parser, "Save", 1, "^((sv)|(save))$", 2, save)
	cli.AddCommand(parser, "Quicksave", 1, "^((qs)|(quicksave))$", 2, quicksave)
	cli.AddCommand(parser, "Load", 1, "^((ld)|(load))$", 2, load)
	cli.AddCommand(parser, "Quickload", 1, "^((ql)|(quickload))$", 2, quickload)
	cli.AddCommand(parser, "Uppercase", 1, "^((u)|(uc)|(upper)|(uppercase))$", 1, uppercase)
	cli.AddCommand(parser, "Lowercase", 1, "^((l)|(lc)|(lower)|(lowercase))$", 1, lowercase)
	cli.AddCommand(parser, "Prefix", 1, "^((p)|(pr)|(pre)|(prefix))$", 2, prefix)
	cli.AddCommand(parser, "Suffix", 1, "^((s)|(po)|(post)|(suffix))$", 2, suffix)
	cli.AddCommand(parser, "Trim start", 1, "^((ts)|(tstart)|(trimstart))$", 2, trimStart)
	cli.AddCommand(parser, "Trim start (exclusive)", 1, "^((tsx)|(tstartx)|(trimstartx))$", 1, trimStartX)
	cli.AddCommand(parser, "Trim end", 1, "^((te)|(tend)|(trimend))$", 2, trimEnd)
	cli.AddCommand(parser, "Trim end (exclusive)", 2, "^((tex)|(tendx)|(trimendx))$", 1, trimEndX)
	cli.AddCommand(parser, "Sort", 1, "^((o)|(sort)|(order))$", 1, sortFunction)
	cli.AddCommand(parser, "Invert", 1, "^((i)|(inv)|(invert))$", 1, invert)
	cli.AddCommand(parser, "Remove Duplicates", 1, "^((rd)|(remdup)|(nodup))$", 1, removeDuplicates)
	cli.AddCommand(parser, "Clear formtting", 1, "^((c)|(cl)|(clear))$", 1, clear)
	cli.AddCommand(parser, "Filter", 1, "^((f)|(filter))$", 2, filter)
	cli.AddCommand(parser, "Filter exclusive", 1, "^((fx)|(filterx))$", 2, filterExclusive)
	cli.AddCommand(parser, "Replace", 1, "^((r)|(replace))$", 3, replace)
	cli.AddCommand(parser, "Replace regex", 1, "^((rx)|(replacex))$", 3, replaceX)
	cli.AddCommand(parser, "Replace Transform", 1, "^((rt)|(replacet))$", 3, replaceT)
	cli.AddCommand(parser, "Replace Regex Transform", 1, "^((rxt)|(rtx)|(replacetx)|(replacext))$", 3, replaceXT)
	cli.AddCommand(parser, "Show Help", 99, ".*", 3, showHelpFunction)

	return parser
}

func showHelp() {
	fmt.Println()
	fmt.Println("************************************************************")
	fmt.Println("* TextTools GO by A. Markoczy")
	fmt.Println("************************************************************")
	fmt.Println("*")
	fmt.Println("**** Help section ****")
	fmt.Println("*")
	fmt.Println("* Single Commands:")
	fmt.Println("*")
	fmt.Println("* u                : UPPERCASE")
	fmt.Println("* l                : lowercase")
	fmt.Println("* c                : Clear formatting")
	fmt.Println("* i                : Invert line order")
	fmt.Println("* sort             : Sort (alt: 'o')")
	fmt.Println("* rdup             : remove all duplicates (alt: 'rd')")
	fmt.Println("* filter [txt]     : Select lines with [txt] (alt: 'f')")
	fmt.Println("* filterx [txt]    : Exclude lines with [txt] (alt: 'fx')")
	fmt.Println("* pre [txt]        : prefix [txt] by line (alt: 'p')")
	fmt.Println("* post [txt]       : suffix [txt] by line (alt: 's')")
	fmt.Println("* ts [txt]         : trim start to end of [txt] by line")
	fmt.Println("* tsx [txt]        : trim start to start of [txt] by line")
	fmt.Println("* te [txt]         : trim end to start of [txt] by line")
	fmt.Println("* tex [txt]        : trim end to end of [txt] by line")
	fmt.Println("* save [path]      : Clipboard to file (alt 'sv')")
	fmt.Println("* load [path]      : File to clipbooard (alt 'ld')")
	fmt.Println("* qs [name]        : Quicksave (save to temp file)")
	fmt.Println("* ql [name]        : Quickload (load from temp file)")
	fmt.Println("* r [in] [out]     : replace all [in] with [out]")
	fmt.Println("* rx [in] [out]    : replace regex mode")
	fmt.Println("* rt [in] [out]    : replace transform backslashes")
	fmt.Println("* rxt [in] [out]   : replace regex transform backslashes")
	fmt.Println("*")
	fmt.Println("* > All operations are applied to the Clipboard")
	fmt.Println("************************************************************")
}

func showHelpFunction(s []string) (interface{}, error) {
	showHelp()
	return s[1], nil
}

// todo: replace, filter, merge, ...

// func isValidCount(s []string, req int) {
// 	if req == -1 {
// 		return true
// 	}
// 	return len(s) == req
// }
