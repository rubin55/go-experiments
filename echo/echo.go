package echo

import (
	"os"
	"strconv"
)

func Echo() string {
	s, sep := "", ""
	for key, val := range os.Args[1:] {
		s += sep + strconv.Itoa(key)
		sep = " "
		s += sep + val
	}

	return s
}
