package krendel

import (
	"bytes"
)

const divider = 1000

type Int64 int64

func (v *Int64) String() string {
	var minus bool
	if *v < 0 {
		minus = true
		*v = 0 - *v
	}

	parts := getSliceInt64(*v)
	if minus {
		return "-" + string(bytes.Join(parts, []byte{44}))
	} else {
		return string(bytes.Join(parts, []byte{44}))
	}
}
