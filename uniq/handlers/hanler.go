package handler

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// надо чтобы хендлеры принимали еще и buffer какой-то
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

// readBuffer не нужен завести временный буффер с отфильрованными байтами и возвращать его по значению
func CountUniq(input io.Reader, val interface{}, writeBuffer []string) ([]string, error) {

	var (
		in   *bufio.Scanner
		prev string
	)

	if len(writeBuffer) != 0 {
		curReader := bytes.NewReader(bytes.NewBufferString(strings.Join(writeBuffer, ",")).Bytes())
		writeBuffer = writeBuffer[:0]
		in = bufio.NewScanner(curReader)
	} else {
		in = bufio.NewScanner(input)
	}

	counts := make(map[string]int)

	for in.Scan() {
		txt := in.Text()
		counts[txt]++

		if prev != txt && prev != "" {
			for line, n := range counts {
				writeBuffer = append(writeBuffer, strconv.Itoa(n), " ", line, "\n")
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

	writeBuffer[len(writeBuffer)-1] = ""

	return writeBuffer, nil
}

func GetRepeatedLines(input io.Reader, val interface{}, writeBuffer []string) ([]string, error) {
	return writeBuffer, nil
}

func GetNotRepeatedLines(input io.Reader, val interface{}, writeBuffer []string) ([]string, error) {
	fmt.Println("GetNotRepeatedLines has been called")
	return writeBuffer, nil
}

func GetLinesCompareNWord(input io.Reader, val interface{}, writeBuffer []string) ([]string, error) {
	// fmt.Println("GetLinesCompareNWord has been called")
	// var (
	// 	in   *bufio.Scanner
	// 	prev string
	// )

	// if len(writeBuffer) != 0 {
	// 	curReader := bytes.NewReader(bytes.NewBufferString(strings.Join(writeBuffer, ",")).Bytes())
	// 	writeBuffer = writeBuffer[:0]
	// 	in = bufio.NewScanner(curReader)
	// } else {
	// 	in = bufio.NewScanner(input)
	// }

	// for in.Scan() {
	// 	txt := in.Text()
	// 	if txt == prev {
	// 		continue
	// 	}
	// 	if txt == io.EOF.Error() {
	// 		break
	// 	}

	// 	writeBuffer = append(writeBuffer, txt, "\n")
	// }

	writeBuffer = append(writeBuffer, "addd some string")

	return writeBuffer, nil
}

func GetLinesCompareNChar(input io.Reader, val interface{}, writeBuffer []string) ([]string, error) {
	fmt.Println("GetLinesCompareNChar has been called")
	return writeBuffer, nil
}

func GetLinesWithoutRegister(input io.Reader, val interface{}, writeBuffer []string) ([]string, error) {
	fmt.Println("GetLinesWithoutRegister has been called")
	return writeBuffer, nil
}

type Handler struct {
	HandleMap map[string]func(input io.Reader, val interface{}, writeBuffer []string) ([]string, error)
}

func NewHandler() *Handler {
	newMap := make(map[string]func(input io.Reader, val interface{}, writeBuffer []string) ([]string, error))
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
