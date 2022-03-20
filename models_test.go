package krendel

import (
	"math"
	"testing"
)

func TestInt64_String(t *testing.T) {
	tests := []struct {
		name string
		v    Int64
		want string
	}{
		{"-123", -123, "-123"},
		{"-1234", -1234, "-1,234"},
		{"-123456", -123456, "-123,456"},
		{"123056", 123056, "123,056"},
		{"123456", 123456, "123,456"},
		{"123006", 123006, "123,006"},
		{"0", 0, "0"},
		{"minimum-1", math.MinInt64 + 1, "-9,223,372,036,854,775,807"},
		{"minimum", math.MinInt64, "-9,223,372,036,854,775,808"},
		{"maximum", math.MaxInt64, "9,223,372,036,854,775,807"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
