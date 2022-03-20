package krendel

import (
	"strconv"
	"testing"
)

func TestDigits(t *testing.T) {
	tests := []struct {
		number Int64
		want   byte
	}{
		{123, 3},
		{1234, 4},
		{12345, 5},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(int(tt.number)), func(t *testing.T) {
			if got := Digits(tt.number); got != tt.want {
				t.Errorf("Digits() = %v, want %v", got, tt.want)
			}
		})
	}
}
