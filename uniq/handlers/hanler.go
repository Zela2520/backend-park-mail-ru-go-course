package handler

import (
	"bufio"
	"fmt"
	"io"
)

func CountUniq(input io.Reader, output io.Writer) error {
	counts := make(map[string]int)
	in := bufio.NewScanner(input)
	var prev string

	for in.Scan() {
		txt := in.Text()
		counts[txt]++

		if prev != txt && prev != "" {
			for line, n := range counts {
				fmt.Fprintln(output, n, line)
			}

			for k := range counts {
				delete(counts, k)
			}
		}

		if txt == prev {
			continue
		}

		prev = txt
	}

	return nil
}

func Uniq(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)
	var prev string

	for in.Scan() {
		txt := in.Text()
		if txt == prev {
			continue
		}
		if txt == io.EOF.Error() {
			break
		}

		prev = txt
		fmt.Fprintln(output, txt)
	}

	return nil
}
