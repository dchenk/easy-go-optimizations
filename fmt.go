package easy_go_optimizations

import "fmt"

var strsLiteral = []string{"\"%s\"", "abc \"%s\" \"%s\"", "abcd \"%s\" ab \"%s\" \"%s\"", "abcde \"%s\" \"%s\" abcde \"%s\"abcde\"%s\" abcde"}

var strsQ = []string{"%q", "abc %q %q", "abcd %q ab %q %q", "abcde %q %q abcde %qabcde%q abcde"}

func PrintfQuote(format string, toQuote []interface{}) string {
	return fmt.Sprintf(format, toQuote...)
}
