package util

import "strconv"

func ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
