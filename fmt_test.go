package easy_go_optimizations

import "testing"

var replace = [][]interface{}{
	{"string to quote"},
	{"string to quote", "string to quote"},
	{"string to quote", "string to quote", "string to quote"},
	{"string to quote", "string to quote", "string to quote", "string to quote"},
}

func TestQuoteLiteral(t *testing.T) {
	for i := range strsLiteral {
		got1 := PrintfQuote(strsLiteral[i], replace[i])
		got2 := PrintfQuote(strsQ[i], replace[i])
		if got1 != got2 {
			t.Errorf("(index %d) got %q and %q", i, got1, got2)
		}
	}
}

var result string

func BenchmarkPrintfLiteral_Small(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result = PrintfQuote(strsLiteral[0], replace[0])
	}
}

func BenchmarkPrintfLiteral_Small2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result = PrintfQuote(strsLiteral[1], replace[1])
	}
}

func BenchmarkPrintfLiteral_Medium(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result = PrintfQuote(strsLiteral[2], replace[2])
	}
}

func BenchmarkPrintfLiteral_Big(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result = PrintfQuote(strsLiteral[3], replace[3])
	}
}

func BenchmarkPrintfQ_Small(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result = PrintfQuote(strsQ[0], replace[0])
	}
}

func BenchmarkPrintfQ_Small2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result = PrintfQuote(strsQ[1], replace[1])
	}
}

func BenchmarkPrintfQ_Medium(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result = PrintfQuote(strsQ[2], replace[2])
	}
}

func BenchmarkPrintfQ_Big(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result = PrintfQuote(strsQ[3], replace[3])
	}
}

// Results April 15, 2018; go version 1.10 darwin/amd64 -- The clear performance winner is using \"s\" but you don't
// get escaping of non-printable characters.
//  BenchmarkPrintfLiteral_Small-8          20000000                72.5 ns/op            32 B/op          1 allocs/op
//  BenchmarkPrintfLiteral_Small2-8         20000000               100 ns/op              48 B/op          1 allocs/op
//  BenchmarkPrintfLiteral_Medium-8         10000000               125 ns/op              64 B/op          1 allocs/op
//  BenchmarkPrintfLiteral_Big-8            10000000               159 ns/op              96 B/op          1 allocs/op
//  BenchmarkPrintfQ_Small-8                10000000               224 ns/op              32 B/op          1 allocs/op
//  BenchmarkPrintfQ_Small2-8                3000000               429 ns/op              48 B/op          1 allocs/op
//  BenchmarkPrintfQ_Medium-8                2000000               619 ns/op              64 B/op          1 allocs/op
//  BenchmarkPrintfQ_Big-8                   2000000               817 ns/op              96 B/op          1 allocs/op
