package handler

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"syscall"

	"github.com/markoczy/gotext/common"
	"github.com/markoczy/goutil/log"
	"golang.org/x/crypto/ssh/terminal"
)

//=================================================
//
// Command Array specification:
//
// s[0]     = Command String
// s[1:n-1] = Command Params (optional)
// s[n]     = Clipboard
//
// => n is the number of params (1 terminated)
//    For a function using only Clipboard as
//    parameter n is equal 1.
//
//=================================================

//*** Single Var Commands (Clipboard = s[1]) ***//

func uppercase(s []string) (interface{}, error) {
	log.Debug("Entry uppercase")
	return strings.ToUpper(s[1]), nil
}

func lowercase(s []string) (interface{}, error) {
	log.Debug("Entry lowercase")
	return strings.ToLower(s[1]), nil
}

func clear(s []string) (interface{}, error) {
	log.Debug("Entry clear")
	return s[1], nil
}

func invert(s []string) (interface{}, error) {
	log.Debug("Entry invert")
	split, sep := common.Split(s[1])
	strs := []string{}
	for _, e := range split {
		strs = append([]string{e}, strs...)
	}
	return strings.Join(strs, sep), nil
}

func paste(s []string) (interface{}, error) {
	fmt.Print(s[1])
	return s[1], nil
}

func sortFunction(s []string) (interface{}, error) {
	log.Debug("Entry sort")
	split, sep := common.Split(s[1])
	sort.Strings(split)
	return strings.Join(split, sep), nil
}

func removeDuplicates(s []string) (interface{}, error) {
	log.Debug("Entry remove duplicates")
	split, sep := common.Split(s[1])
	strs := []string{}
	for _, e := range split {
		isDup := false
		for _, str := range strs {
			if e == str {
				isDup = true
			}
		}
		if !isDup {
			strs = append(strs, e)
		}
	}
	return strings.Join(strs, sep), nil
}

func rot13(s []string) (interface{}, error) {
	ret := ""
	for _, v := range s[1] {
		chr := v
		if chr > 64 && chr < 91 {
			chr = 65 + ((chr - 52) % 26)
		} else if chr > 96 && chr < 123 {
			chr = 97 + ((chr - 84) % 26)
		}
		ret += string(chr)
	}
	return ret, nil
}

func purge(s []string) (interface{}, error) {
	log.Debug(fmt.Sprintf("Entry purge"))
	askConfirm := !(len(s) > 2 && strings.ToUpper(s[1]) == "-Y")

	doPurge := false
	if askConfirm {
		fmt.Print("Really purge all quicksaves? [y/n]: ")
		var input string
		fmt.Scanf("%s", &input)
		if strings.ToUpper(input) == "Y" {
			doPurge = true
		}
	}

	if !askConfirm || doPurge {
		folder, err := common.InitQuickSaveDir()
		if err != nil {
			return nil, err
		}

		_, err = os.Stat(folder)
		if !os.IsNotExist(err) {
			err = os.RemoveAll(folder)
			if err != nil {
				return nil, err
			}
		}
	}

	return strings.ToLower(s[1]), nil
}

func login(s []string) (interface{}, error) {
	log.Debug(fmt.Sprintf("Entry login"))
	var pw string
	if len(s) > 2 {
		pw = s[1]
	} else {
		fmt.Print("Enter password: ")
		tmp, err := terminal.ReadPassword(int(syscall.Stdin))
		fmt.Print("\n")
		if err != nil {
			return nil, err
		}
		pw = string(tmp)
	}

	fmt.Println(pw)

	return strings.ToLower(s[1]), nil
}

//*** Double Var Commands (Clipboard = s[2]) ***//

func filter(s []string) (interface{}, error) {
	log.Debug("Entry filter")
	strs := []string{}
	split, sep := common.Split(s[2])
	for _, e := range split {
		if strings.Index(e, s[1]) > -1 {
			strs = append(strs, e)
		}
	}
	return strings.Join(strs, sep), nil
}

func filterExclusive(s []string) (interface{}, error) {
	log.Debug("Entry filter")
	strs := []string{}
	split, sep := common.Split(s[2])
	for _, e := range split {
		if strings.Index(e, s[1]) == -1 {
			strs = append(strs, e)
		}
	}
	return strings.Join(strs, sep), nil
}

func prefix(s []string) (interface{}, error) {
	log.Debug("Entry prefix")
	split, sep := common.Split(s[2])
	strs := common.MapArray(split, func(e string) string {
		return s[1] + e
	})
	log.Debugf("Array now: %v\n", strs)
	return strings.Join(strs, sep), nil
}

// tt s abc
func suffix(s []string) (interface{}, error) {
	log.Debug("Entry suffix")
	split, sep := common.Split(s[2])
	strs := common.MapArray(split, func(e string) string {
		return e + s[1]
	})
	return strings.Join(strs, sep), nil
}

// tt ts abc
func trimStart(s []string) (interface{}, error) {
	log.Debug("Entry trimStart")
	var size = len(s[1])
	split, sep := common.Split(s[2])
	strs := common.MapArray(split, func(e string) string {
		var idx = strings.Index(e, s[1])
		if idx > -1 {
			return e[idx+size:]
		} else {
			return e
		}
	})

	return strings.Join(strs, sep), nil
}

// tt tsx abc
func trimStartX(s []string) (interface{}, error) {
	log.Debug("Entry trimStartX")
	split, sep := common.Split(s[2])
	strs := common.MapArray(split, func(e string) string {
		var idx = strings.Index(e, s[1])
		if idx > -1 {
			return e[idx:]
		} else {
			return e
		}
	})

	return strings.Join(strs, sep), nil

}

func trimEnd(s []string) (interface{}, error) {
	log.Debug("Entry trimEnd")
	split, sep := common.Split(s[2])
	strs := common.MapArray(split, func(e string) string {
		var idx = strings.Index(e, s[1])
		if idx > -1 {
			return e[:idx]
		} else {
			return e
		}
	})

	return strings.Join(strs, sep), nil
}

func trimEndX(s []string) (interface{}, error) {
	log.Debug("Entry trimEndX")
	var size = len(s[1])
	split, sep := common.Split(s[2])
	strs := common.MapArray(split, func(e string) string {
		// log.Debugf("idx: %v", idx)
		var idx = strings.Index(e, s[1])
		if idx > -1 {
			return e[:idx+size]
		} else {
			return e
		}
	})

	return strings.Join(strs, sep), nil
}

func save(s []string) (interface{}, error) {
	log.Debug("Entry save")
	err := ioutil.WriteFile(s[1], []byte(s[2]), 0666)
	if err != nil {
		return nil, err
	}
	return s[2], nil
}

func load(s []string) (interface{}, error) {
	log.Debug("Entry load")
	dat, err := ioutil.ReadFile(s[1])
	if err != nil {
		return nil, err
	}
	return string(dat), nil
}

func quicksave(s []string) (interface{}, error) {
	log.Debug("Entry quicksave")
	folder, err := common.InitQuickSaveDir()
	if err != nil {
		return nil, err
	}

	path := path.Join(folder, s[1])
	err = ioutil.WriteFile(path, []byte(s[2]), 0666)
	if err != nil {
		return nil, err
	}
	return s[2], nil
}

func quickload(s []string) (interface{}, error) {
	log.Debug("Entry quickload")
	folder, err := common.InitQuickSaveDir()
	if err != nil {
		return nil, err
	}
	path := path.Join(folder, s[1])
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return string(dat), nil
}

func skipBegin(s []string) (interface{}, error) {
	log.Debug("Entry skipBegin")
	n, err := strconv.Atoi(s[1])
	if err != nil {
		return nil, err
	}
	split, sep := common.Split(s[2])

	return strings.Join(split[n:], sep), nil
}

func skipEnd(s []string) (interface{}, error) {
	log.Debug("Entry skipBegin")
	n, err := strconv.Atoi(s[1])
	if err != nil {
		return nil, err
	}
	split, sep := common.Split(s[2])

	return strings.Join(split[:len(split)-n], sep), nil
}

//*** Triple Var Commands (Clipboard = s[3]) ***//

func replace(s []string) (interface{}, error) {
	// 1 from, 2 to, 3 clipboard
	log.Debug("Entry replace")
	ret := strings.Replace(s[3], s[1], s[2], -1)
	return ret, nil
}

func replaceX(s []string) (interface{}, error) {
	// 1 from, 2 to, 3 clipboard
	log.Debug("Entry replace regex")
	matcher, err := regexp.Compile(s[1])
	if err != nil {
		return nil, err
	}
	ret := matcher.ReplaceAllString(s[3], s[2])
	return ret, nil
}

func replaceT(s []string) (interface{}, error) {
	// 1 from, 2 to, 3 clipboard
	log.Debug("Entry replace transform")
	to := common.TransformBackslashes(s[2])
	log.Debugf("To: %s", to)
	ret := strings.Replace(s[3], s[1], to, -1)
	return ret, nil
}

func replaceXT(s []string) (interface{}, error) {
	// 1 from, 2 to, 3 clipboard
	log.Debug("Entry replace regex translate")
	matcher, err := regexp.Compile(s[1])
	if err != nil {
		return nil, err
	}
	to := common.TransformBackslashes(s[2])
	ret := matcher.ReplaceAllString(s[3], to)
	return ret, nil
}

//*** Reusable Low-Level Functions ***//
