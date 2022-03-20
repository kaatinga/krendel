package krendel

import (
	"github.com/dustin/go-humanize"
	"testing"
)

func BenchmarkComma(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		humanize.Comma(123456)
		humanize.Comma(1234567891234)
		humanize.Comma(0)
	}
}

var testInt1 Int64 = 123456
var testInt2 Int64 = 1234567891234
var testInt3 Int64 = 0

func BenchmarkInt64_String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testInt1.String()
		testInt2.String()
		testInt3.String()
	}
}
