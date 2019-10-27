package echo

import (
	"strconv"
)

func Echo(i []string) string {
	s, sep := "", ""
	for key, val := range i[1:] {
		s += sep + strconv.Itoa(key)
		sep = " "
		s += sep + val
	}

	return s
}
