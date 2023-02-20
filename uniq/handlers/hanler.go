package handler

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func Uniq(input io.Reader, output []string) ([]string, error) {
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
		output = append(output, txt)
	}

	return output, nil
}

func CountUniq(input io.Reader, val interface{}, writeBuffer []string) ([]string, error) {
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
			writeBuffer = append(writeBuffer, strconv.Itoa(counts[prev])+" "+prev)

			delete(counts, prev)
		}

		prev = curText
	}

	writeBuffer = append(writeBuffer, strconv.Itoa(counts[prev])+" "+prev) // добавили последний ключ
	return writeBuffer, nil
}

func GetRepeatedLines(input io.Reader, val interface{}, writeBuffer []string) ([]string, error) {
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
		if exist == true && counts[prev] > 1 {
			writeBuffer = append(writeBuffer, prev)

			delete(counts, prev)
		}

		if exist == true {
			delete(counts, prev)
		}

		prev = curText
	}

	if counts[prev] > 1 {
		writeBuffer = append(writeBuffer, prev)
	}

	return writeBuffer, nil
}

func GetNotRepeatedLines(input io.Reader, val interface{}, writeBuffer []string) ([]string, error) {
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
		if exist == true && counts[prev] == 1 {
			writeBuffer = append(writeBuffer, prev)

			delete(counts, prev)
		}

		if exist == true {
			delete(counts, prev)
		}

		prev = curText
	}

	if counts[prev] == 1 {
		writeBuffer = append(writeBuffer, prev)
	}

	return writeBuffer, nil
}

func GetLinesCompareNWord(input io.Reader, val interface{}, writeBuffer []string) ([]string, error) {
	fmt.Println("GetLinesCompareNWord has been called")
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
