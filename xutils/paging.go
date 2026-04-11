package xutils

import "github.com/jindasoft/jinda-platforms/xconst"

func SetDefaultPageSize(offset, limit int64) (int64, int64) {
	o := max(offset, 0)

	var l int64
	if limit > 0 {
		l = min(limit, xconst.DefaultMaxPageLimit)
	} else {
		l = xconst.DefaultMinPageLimit
	}

	return o, l
}

func SetDefaultPageSizeInt(offset, limit int) (int, int) {
	o := max(offset, 0)

	var l int
	if limit > 0 {
		l = min(limit, int(xconst.DefaultMaxPageLimit))
	} else {
		l = int(xconst.DefaultMinPageLimit)
	}

	return o, l
}
