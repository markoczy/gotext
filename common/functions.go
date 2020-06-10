package common

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

const (
	crlf = "\r\n"
	lf   = "\n"
	cr   = "\r"
)

func initRootDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	rootDir := path.Join(usr.HomeDir, ".gotext")
	err = CreateFolderIfNotExists(rootDir)
	return rootDir, err
}

func initChildDir(child string) (string, error) {
	rootDir, err := initRootDir()
	if err != nil {
		return "", err
	}
	childDir := path.Join(rootDir, child)
	err = CreateFolderIfNotExists(childDir)
	return childDir, err
}

func InitQuickSaveDir() (string, error) {
	return initChildDir("data")
}

func InitKeyDir() (string, error) {
	return initChildDir("key")
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func CreateFolderIfNotExists(path string) error {
	var err error
	if !FileExists(path) {
		err = os.Mkdir(path, 0644)
	}
	return err
}

func TransformBackslashes(s string) string {
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

func MapArray(split []string, mapper func(string) string) []string {
	strs := []string{}
	for _, e := range split {
		str := mapper(e)
		strs = append(strs, str)
	}
	return strs
}

func GetPassword() (string, error) {
	fmt.Print("Enter password: ")
	tmp, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Print("\n")
	if err != nil {
		return "", err
	}
	return string(tmp), nil
}

// split splits a single String into a slice of strings. Supports any common end-line sequence and returns the sequence that was found as second return value.
func Split(s string) ([]string, string) {
	// TODO performance could be optimized by first searching for any of the line endings and not using trySplit.
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
