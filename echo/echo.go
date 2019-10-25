package echo

import (
	"os"
	"strings"
)

func Echo() string {
	return strings.Join(os.Args[1:], " ")
}
