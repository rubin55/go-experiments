package unique

import (
	"bufio"
	"fmt"
	"os"
)

type File = os.File

func Unique(f *File) string {
	s := ""
	counts := make(map[string]int)
	countLines(f, counts)
	f.Close()

	for line, n := range counts {
		if n > 0 {
			s += fmt.Sprintf("%d\t%s\n", n, line)
		}
	}

	return s
}

func NewUnique(files []string) string {
	s := ""
	counts := make(map[string]int)

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unique: %v\n", err)
			continue
		}
		countLines(f, counts)
		f.Close()
	}

	for line, n := range counts {
		if n > 1 {
			s += fmt.Sprintf("%d\t%s\n", n, line)
		}
	}
	return s
}

func countLines(f *File, counts map[string]int) {
	i := bufio.NewScanner(f)
	for i.Scan() {
		counts[i.Text()]++
	}
}