package handler

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func selectReader(writeBuffer []string, input io.Reader) *bufio.Scanner {
	if len(writeBuffer) != 0 {
		writeBuffer = append(writeBuffer, "\n") // буффер не пустой - запись должна начаться с новой строки
		curReader := bytes.NewReader(bytes.NewBufferString(strings.Join(writeBuffer, ",")).Bytes())
		return bufio.NewScanner(curReader)
	} else {
		return bufio.NewScanner(input)
	}
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

func CountUniq(input io.Reader, val interface{}, writeBuffer []string) ([]string, error) {

	var (
		in      *bufio.Scanner
		curText string
		prev    string
	)

	in = selectReader(writeBuffer, input)

	writeBuffer = writeBuffer[:0] // чистим буффер - будем принимать отфильтровнные данные из Readera

	counts := make(map[string]int)

	for in.Scan() {
		curText = in.Text()
		counts[curText]++

		if prev != curText && counts[prev] != 0 && prev != "" { // counts[prev] != 0 - не выводим сразу новое значение, проверяем есть ли дальше повторы
			for line, n := range counts {
				writeBuffer = append(writeBuffer, strconv.Itoa(n), " ", line, "\n")
			}

			for k := range counts {
				delete(counts, k)
			}
		}

		if curText == prev {
			continue
		}

		prev = curText
	}

	if counts[prev] > 1 {
		writeBuffer = append(writeBuffer, strconv.Itoa(counts[prev]), " ", curText, "\n")
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
	var (
		in      *bufio.Scanner
		curText string
		prev    string
	)

	in = selectReader(writeBuffer, input)

	counts := make(map[string]int)

	writeBuffer = writeBuffer[:0]

	for in.Scan() {
		curText = in.Text()
		counts[curText]++

		if prev == curText {
			continue
		}

		_, exist := counts[prev]
		if exist == true {
			writeBuffer = append(writeBuffer, strconv.Itoa(counts[prev]), " ", prev, "\n")

			delete(counts, prev)
		}

		prev = curText
	}

	writeBuffer[len(writeBuffer)-1] = ""

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
	newMap["c"] = GetLinesCompareNWord // CountUniq
	newMap["d"] = GetRepeatedLines
	newMap["u"] = GetNotRepeatedLines
	newMap["f"] = GetLinesCompareNWord
	newMap["s"] = GetLinesCompareNChar
	newMap["i"] = GetLinesWithoutRegister
	return &Handler{
		HandleMap: newMap,
	}
}
