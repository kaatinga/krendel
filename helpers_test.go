package krendel

import (
	"math"
	"reflect"
	"strconv"
	"testing"
)

func Test_getSliceInt64(t *testing.T) {
	tests := []struct {
		num  int64
		want [][]byte
	}{
		{1234, [][]byte{{0x31}, {0x32, 0x33, 0x34}}},
		{1034, [][]byte{{0x31}, {0x30, 0x33, 0x34}}},
		{123456, [][]byte{{0x31, 0x32, 0x33}, {0x34, 0x35, 0x36}}},
		{123456789, [][]byte{{0x31, 0x32, 0x33}, {0x34, 0x35, 0x36}, {0x37, 0x38, 0x39}}},
		{0, [][]byte{{0x30}}},
		{1, [][]byte{{0x31}}},
		{math.MaxInt64, [][]byte{{57}, {50, 50, 51}, {51, 55, 50}, {48, 51, 54}, {56, 53, 52}, {55, 55, 53}, {56, 48, 55}}},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(int(tt.num)), func(t *testing.T) {
			if got := getSliceInt64(tt.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSliceInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
