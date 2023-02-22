package handler

import (
	"bufio"
	"io"
	"strconv"

	"github.com/Zela2520/backend-park-mail-ru-go-course.git/uniq/param"
	"github.com/pkg/errors"
)

func Uniq(input io.Reader, output []string, numberOfSkipWords int, numberOfSkipChar int, register bool) ([]string, error) {
	in := bufio.NewScanner(input)
	var (
		prev            string
		curCompareLine  string
		prevCompareLine string
		err             error
	)

	for in.Scan() {
		txt := in.Text()

		curCompareLine, prevCompareLine, err = processModifyingOptions(txt, prev, numberOfSkipWords, numberOfSkipChar, register)
		if err != nil {
			return nil, errors.Wrap(err, "processModifyingOptions error:")
		}

		if curCompareLine == prevCompareLine {
			continue
		}
		if curCompareLine == io.EOF.Error() {
			break
		}

		prev = txt
		output = append(output, txt)
	}

	return output, nil
}

// TODO: сделать компаратор - передавать его в функию, поскольку 3 обработчика отличаются тока правилом сравнения.
func CountUniq(input io.Reader, writeBuffer []string, numberOfSkipWords int, numberOfSkipChar int, register bool) ([]string, error) {
	var (
		in              *bufio.Scanner
		curText         string
		prev            string
		curCompareLine  string
		prevCompareLine string
		err             error
	)

	in = selectReader(writeBuffer, input)
	writeBuffer = writeBuffer[:0]
	counts := make(map[string]int)

	for in.Scan() {
		curText = in.Text()

		curCompareLine, prevCompareLine, err = processModifyingOptions(curText, prev, numberOfSkipWords, numberOfSkipChar, register)
		if err != nil {
			return nil, errors.Wrap(err, "processModifyingOptions error:")
		}

		counts[curCompareLine]++

		if prevCompareLine == curCompareLine {
			continue
		}

		_, exist := counts[prevCompareLine]
		if exist == true {
			writeBuffer = append(writeBuffer, strconv.Itoa(counts[prevCompareLine])+" "+prev)

			delete(counts, prevCompareLine)
		}

		prev = curText
	}

	if counts[curCompareLine] == 1 && counts[prevCompareLine] == 0 {
		writeBuffer = append(writeBuffer, strconv.Itoa(counts[curCompareLine])+" "+prev)
	} else {
		writeBuffer = append(writeBuffer, strconv.Itoa(counts[prevCompareLine])+" "+prev)
	}

	return writeBuffer, nil
}

func GetRepeatedLines(input io.Reader, writeBuffer []string, numberOfSkipWords int, numberOfSkipChar int, register bool) ([]string, error) {
	var (
		in              *bufio.Scanner
		curText         string
		prev            string
		curCompareLine  string
		prevCompareLine string
		err             error
	)

	in = selectReader(writeBuffer, input)
	counts := make(map[string]int)
	writeBuffer = writeBuffer[:0]

	for in.Scan() {
		curText = in.Text()

		curCompareLine, prevCompareLine, err = processModifyingOptions(curText, prev, numberOfSkipWords, numberOfSkipChar, register)
		if err != nil {
			return nil, errors.Wrap(err, "processModifyingOptions error:")
		}

		counts[curCompareLine]++

		if prevCompareLine == curCompareLine {
			continue
		}

		_, exist := counts[prevCompareLine]
		if exist == true && counts[prevCompareLine] > 1 {
			writeBuffer = append(writeBuffer, prev)

			delete(counts, prevCompareLine)
		}

		if exist == true {
			delete(counts, prevCompareLine)
		}

		prev = curText
	}

	if counts[prev] > 1 || counts[curCompareLine] > 1 {
		writeBuffer = append(writeBuffer, prev)
	}

	return writeBuffer, nil
}

func GetNotRepeatedLines(input io.Reader, writeBuffer []string, numberOfSkipWords int, numberOfSkipChar int, register bool) ([]string, error) {
	var (
		in              *bufio.Scanner
		curText         string
		prev            string
		curCompareLine  string
		prevCompareLine string
		err             error
	)

	in = selectReader(writeBuffer, input)
	writeBuffer = writeBuffer[:0]
	counts := make(map[string]int)

	for in.Scan() {
		curText = in.Text()

		curCompareLine, prevCompareLine, err = processModifyingOptions(curText, prev, numberOfSkipWords, numberOfSkipChar, register)
		if err != nil {
			return nil, errors.Wrap(err, "processModifyingOptions error:")
		}

		counts[curCompareLine]++

		if prevCompareLine == curCompareLine {
			continue
		}

		_, exist := counts[prevCompareLine]
		if exist == true && counts[prevCompareLine] == 1 {
			writeBuffer = append(writeBuffer, prev)

			delete(counts, prevCompareLine)
		}

		if exist == true {
			delete(counts, prevCompareLine)
		}

		prev = curText
	}

	if counts[prev] == 1 || counts[curCompareLine] == 1 {
		writeBuffer = append(writeBuffer, prev)
	}

	return writeBuffer, nil
}

type Handler struct {
	HandleMap           map[string]func(input io.Reader, writeBuffer []string, numberOfSkipWords int, numberOfSkipChar int, register bool) ([]string, error)
	activeRegister      bool
	numberOfSkipWords   int
	numberOfSkipSymbols int
}

func NewHandler(args ...param.Param) *Handler { // _numberOfSkipWords int, _numberOfSkipSymbols int
	newMap := make(map[string]func(input io.Reader, writeBuffer []string, numberOfSkipWords int, numberOfSkipChar int, register bool) ([]string, error))
	// как привязать методы объекта мапы к ключам функции ?
	newMap["c"] = CountUniq
	newMap["d"] = GetRepeatedLines
	newMap["u"] = GetNotRepeatedLines
	return &Handler{
		HandleMap:           newMap,
		activeRegister:      args[0].OptionValue.(bool),
		numberOfSkipWords:   args[1].OptionValue.(int),
		numberOfSkipSymbols: args[2].OptionValue.(int),
	}
}
