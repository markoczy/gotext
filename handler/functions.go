package handler

import (
	"github.com/markoczy/goutil/log"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
)

var xNewLine = regexp.MustCompile("\r?\n")

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
	log.Debug("Entry save")
	filepath := path.Join(os.TempDir(), "gotext")
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		os.Mkdir(filepath, 0644)
	}
	filepath = path.Join(filepath, s[1])
	err = ioutil.WriteFile(filepath, []byte(s[2]), 0666)
	if err != nil {
		return nil, err
	}
	return s[2], nil
}

func quickload(s []string) (interface{}, error) {
	log.Debug("Entry save")
	path := path.Join(os.TempDir(), "gotext", s[1])
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return string(dat), nil
}

func uppercase(s []string) (interface{}, error) {
	log.Debug("Entry uppercase")
	return strings.ToUpper(s[1]), nil
}

func lowercase(s []string) (interface{}, error) {
	log.Debug("Entry lowercase")
	return strings.ToLower(s[1]), nil
}

func filter(s []string) (interface{}, error) {
	log.Debug("Entry filter")
	strs := []string{}
	split, sep := split(s[2])
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
	split, sep := split(s[2])
	for _, e := range split {
		if strings.Index(e, s[1]) == -1 {
			strs = append(strs, e)
		}
	}
	return strings.Join(strs, sep), nil
}

func prefix(s []string) (interface{}, error) {
	log.Debug("Entry prefix")
	split, sep := split(s[2])
	strs := mapArray(split, func(e string) string {
		return s[1] + e
	})
	log.Debugf("Array now: %v\n", strs)
	return strings.Join(strs, sep), nil
}

// tt s abc
func suffix(s []string) (interface{}, error) {
	log.Debug("Entry suffix")
	split, sep := split(s[2])
	strs := mapArray(split, func(e string) string {
		return e + s[1]
	})
	return strings.Join(strs, sep), nil
}

// tt ts abc
func trimStart(s []string) (interface{}, error) {
	log.Debug("Entry trimStart")
	var size = len(s[1])
	split, sep := split(s[2])
	strs := mapArray(split, func(e string) string {
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
	split, sep := split(s[2])
	strs := mapArray(split, func(e string) string {
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
	split, sep := split(s[2])
	strs := mapArray(split, func(e string) string {
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
	split, sep := split(s[2])
	strs := mapArray(split, func(e string) string {
		// log.Debugf("idx: %v", idx)
		var idx = strings.Index(e, s[1])
		if idx > -1 {
			return e[:idx-size+1]
		} else {
			return e
		}
	})

	return strings.Join(strs, sep), nil
}

func sortFunction(s []string) (interface{}, error) {
	log.Debug("Entry sort")
	// strs := xNewLine.Split(s[1], -1)
	split, sep := split(s[1])
	sort.Strings(split)
	return strings.Join(split, sep), nil
}

func invert(s []string) (interface{}, error) {
	log.Debug("Entry invert")
	split, sep := split(s[1])
	strs := []string{}
	for _, e := range split {
		strs = append([]string{e}, strs...)
	}
	return strings.Join(strs, sep), nil
}

func removeDuplicates(s []string) (interface{}, error) {
	log.Debug("Entry remove duplicates")
	split, sep := split(s[1])
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

func clear(s []string) (interface{}, error) {
	log.Debug("Entry clear")
	return s[1], nil
}

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
	to := transformBackslashes(s[2])
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
	to := transformBackslashes(s[2])
	ret := matcher.ReplaceAllString(s[3], to)
	return ret, nil
}

func transformBackslashes(s string) string {
	// \t Insert a tab in the text at this point.
	// \b Insert a backspace in the text at this point.
	// \n Insert a newline in the text at this point.
	// \r Insert a carriage return in the text at this point.
	// \f Insert a formfeed in the text at this point.
	ret := s
	ret = strings.Replace(ret, "\\t", "\t", -1)
	ret = strings.Replace(ret, "\\b", "\b", -1)
	ret = strings.Replace(ret, "\\n", "\n", -1)
	ret = strings.Replace(ret, "\\r", "\r", -1)
	ret = strings.Replace(ret, "\\f", "\f", -1)
	return ret
}

const crlf = "\r\n"
const lf = "\n"
const cr = "\r"

func split(s string) ([]string, string) {
	split, ok := trySplit(s, crlf)
	if ok {
		return split, crlf
	}
	split, ok = trySplit(s, lf)
	if ok {
		return split, lf
	}
	split, ok = trySplit(s, cr)
	if ok {
		return split, lf
	}
	// default: lf
	return split, lf
}

func trySplit(s string, separator string) ([]string, bool) {
	split := strings.Split(s, separator)
	success := false
	if len(split) > 1 {
		success = true
	}
	return split, success
}

func mapArray(split []string, mapper func(string) string) []string {
	strs := []string{}
	for _, e := range split {
		str := mapper(e)
		strs = append(strs, str)
	}
	return strs
}
