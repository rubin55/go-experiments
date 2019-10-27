package unique

import (
	"bufio"
	"fmt"
	"os"
)

type File = os.File

func Unique(i *File) string {
	s := ""
	counts := make(map[string]int)
	input :=bufio.NewScanner(i)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 0 {
			s += fmt.Sprintf("%d\t%s\n", n, line)
		}
	}

	return s
}