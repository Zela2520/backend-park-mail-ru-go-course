package handler

import (
	"bufio"
	"fmt"
	"io"
)

func Uniq(input io.Reader, output io.Writer, val interface{}) error {
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

func CountUniq(input io.Reader, output io.Writer, val interface{}) error {
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

func GetRepeatedLines(input io.Reader, output io.Writer, val interface{}) error {
	fmt.Println("GetRepeatedLines has been called")
	return nil
}

func GetNotRepeatedLines(input io.Reader, output io.Writer, val interface{}) error {
	fmt.Println("GetNotRepeatedLines has been called")
	return nil
}

func GetLinesCompareNWord(input io.Reader, output io.Writer, val interface{}) error {
	fmt.Println("GetLinesCompareNWord has been called")
	return nil
}

func GetLinesCompareNChar(input io.Reader, output io.Writer, val interface{}) error {
	fmt.Println("GetLinesCompareNChar has been called")
	return nil
}

func GetLinesWithoutRegister(input io.Reader, output io.Writer, val interface{}) error {
	fmt.Println("GetLinesWithoutRegister has been called")
	return nil
}

type Handler struct {
	HandleMap map[string]func(input io.Reader, output io.Writer, val interface{}) error
}

func NewHandler() *Handler {
	newMap := make(map[string]func(input io.Reader, output io.Writer, val interface{}) error)
	newMap["c"] = CountUniq
	newMap["d"] = GetRepeatedLines
	newMap["u"] = GetNotRepeatedLines
	newMap["f"] = GetLinesCompareNWord
	newMap["s"] = GetLinesCompareNChar
	newMap["i"] = GetLinesWithoutRegister
	return &Handler{
		HandleMap: newMap,
	}
}
