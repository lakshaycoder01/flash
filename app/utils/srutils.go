package utils

import "strconv"

func ParseInt(val string, def int) int {
	i, e := strconv.Atoi(val)

	if e != nil {
		return def
	}

	return i
}
