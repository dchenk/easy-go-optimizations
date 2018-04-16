package easy_go_optimizations

import "strings"

func init() {

	strings.Index("abc", "b")

	someVar := "abc"
	strings.Index(someVar, "b")

	strings.IndexByte("abc", 'b')
	strings.IndexByte(someVar, 'b')

}
