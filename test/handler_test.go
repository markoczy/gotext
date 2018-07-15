package test

import (
	"github.com/markoczy/gotext/handler"
	"github.com/markoczy/goutil/log"
	"github.com/markoczy/goutil/log/logconfig"
	"testing"
)

var strnil = "<nil>"

func initTest(name string) {
	logconfig.SetDefaultLogLevel(logconfig.DEBUG)
	log.Debugf("Starting test %s", name)
}

func ptr(s string) *string {
	return &s
}

func getOrNil(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return strnil
}

type Testtable struct {
	args   []string
	clip   *string
	result *string
}

func execTests(tables []Testtable, t *testing.T) {
	for _, table := range tables {
		out, err := handler.Exec(table.args, table.clip)
		if err != nil {
			t.Errorf("Error occured: %s", err.Error())
		}
		str := getOrNil(out)
		res := getOrNil(table.result)
		if str != res {
			t.Errorf("Result of (%v, %s) was incorrect, got: %s, want: %s.", table.args, *table.clip, str, res)
		}
	}
}

// TestHelp help function
func TestHelp(t *testing.T) {
	initTest("help function")
	tables := []Testtable{
		{[]string{}, ptr("abc"), nil},
		{[]string{"invalidcommand"}, ptr("abc"), nil},
		{[]string{"u", "42"}, ptr("abc"), nil},
	}
	execTests(tables, t)
}

// TestUppercase uppercase function
func TestUppercase(t *testing.T) {
	initTest("uppercase")
	tables := []Testtable{
		{[]string{"u"}, ptr("test"), ptr("TEST")},
		{[]string{"uc"}, ptr("abc123"), ptr("ABC123")},
	}
	execTests(tables, t)
}

// TestLowercase lowercase function
func TestLowercase(t *testing.T) {
	initTest("lowercase")
	tables := []Testtable{
		{[]string{"l"}, ptr("TEST"), ptr("test")},
		{[]string{"lc"}, ptr("ABC123"), ptr("abc123")},
	}
	execTests(tables, t)
}

// TestPrefix prefix function
func TestPrefix(t *testing.T) {
	initTest("prefix")
	tables := []Testtable{
		{[]string{"p", "abc"}, ptr("xyz"), ptr("abcxyz")},
		{[]string{"pre", "<<<"}, ptr("abc"), ptr("<<<abc")},
	}
	execTests(tables, t)
}

// TestSuffix suffix function
func TestSuffix(t *testing.T) {
	initTest("suffix")
	tables := []Testtable{
		{[]string{"s", "xyz"}, ptr("abc"), ptr("abcxyz")},
		{[]string{"post", ">>>"}, ptr("abc"), ptr("abc>>>")},
	}
	execTests(tables, t)
}

// TestTrimStrart trim start function
func TestTrimStart(t *testing.T) {
	initTest("trim start")
	tables := []Testtable{
		{[]string{"ts", "abc"}, ptr("kkkkkabcxyz"), ptr("xyz")},
	}
	execTests(tables, t)
}

// TestTrimStrart trim start function
func TestTrimEnd(t *testing.T) {
	initTest("trim end")
	tables := []Testtable{
		{[]string{"te", "abc"}, ptr("xyzabcrrrrrrrr"), ptr("xyz")},
	}
	execTests(tables, t)
}

// TestTrimStartX trim start function
func TestTrimStartX(t *testing.T) {
	initTest("trim start exclusive")
	tables := []Testtable{
		{[]string{"tsx", "abc"}, ptr("kkkkkabcxyz"), ptr("abcxyz")},
	}
	execTests(tables, t)
}

// TestTrimEndX trim end function
func TestTrimEndX(t *testing.T) {
	initTest("trim end exclusive")
	tables := []Testtable{
		{[]string{"tex", "xyz"}, ptr("abcxyzrrrrrrr"), ptr("abcxyz")},
		{[]string{"tex", "r"}, ptr("abcrxyz"), ptr("abcr")},
	}
	execTests(tables, t)
}

// TestSort sort function
func TestSort(t *testing.T) {
	initTest("sort")
	tables := []Testtable{
		{[]string{"o"}, ptr("d\nc\nb\na"), ptr("a\nb\nc\nd")},
		{[]string{"sort"}, ptr("d\r\nc\r\nb\r\na"), ptr("a\r\nb\r\nc\r\nd")},
	}
	execTests(tables, t)
}

// TestTrimStrart trim start function
func TestInvert(t *testing.T) {
	initTest("invert")
	tables := []Testtable{
		{[]string{"i"}, ptr("d\nc\nb\na"), ptr("a\nb\nc\nd")},
		{[]string{"invert"}, ptr("d\r\nc\r\nb\r\na"), ptr("a\r\nb\r\nc\r\nd")},
	}
	execTests(tables, t)
}

// TestRemoveDuplicates
func TestRemoveDuplicates(t *testing.T) {
	initTest("remove duplicates")
	tables := []Testtable{
		{[]string{"rd"}, ptr("a\na\nb\nb"), ptr("a\nb")},
		{[]string{"remdup"}, ptr("a\r\na\r\nb\r\nb"), ptr("a\r\nb")},
	}
	execTests(tables, t)
}

// TestFilter
func TestFilter(t *testing.T) {
	initTest("filter")
	tables := []Testtable{
		{[]string{"f", "a"}, ptr("a\na\nb\nc"), ptr("a\na")},
		{[]string{"filter", "a"}, ptr("a\r\na\r\nb\r\nc"), ptr("a\r\na")},
	}
	execTests(tables, t)
}

// TestFilterX
func TestFilterX(t *testing.T) {
	initTest("filter exclusive")
	tables := []Testtable{
		{[]string{"fx", "a"}, ptr("a\na\nb\nc"), ptr("b\nc")},
		{[]string{"filterx", "a"}, ptr("a\r\na\r\nb\r\nc"), ptr("b\r\nc")},
	}
	execTests(tables, t)
}

// TestReplace ...
func TestReplace(t *testing.T) {
	initTest("replace")
	tables := []Testtable{
		{[]string{"r", "a", "b"}, ptr("abc\nxyz"), ptr("bbc\nxyz")},
		{[]string{"replace", "a", "b"}, ptr("abc\r\nxyz"), ptr("bbc\r\nxyz")},
		{[]string{"replace", "a", "\\n"}, ptr("abc"), ptr("\\nbc")},
	}
	execTests(tables, t)
}

// TestReplaceX replace regeX
func TestReplaceX(t *testing.T) {
	initTest("replacex")
	tables := []Testtable{
		{[]string{"rx", "\r?\n", ";"}, ptr("abc\nxyz"), ptr("abc;xyz")},
		{[]string{"rx", "\r?\n", ";"}, ptr("abc\r\nxyz"), ptr("abc;xyz")},
	}
	execTests(tables, t)
}

// TestReplaceX replace regeX + Transform backslashes
func TestReplaceXT(t *testing.T) {
	initTest("replacex")
	tables := []Testtable{
		{[]string{"rxt", ";", "\\n"}, ptr("abc;xyz"), ptr("abc\nxyz")},
		{[]string{"rxt", ";", "\\r\\n"}, ptr("abc;xyz"), ptr("abc\r\nxyz")},
	}
	execTests(tables, t)
}
