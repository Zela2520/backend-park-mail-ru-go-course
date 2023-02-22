package handler

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
)

func skipWords(curLine string, numberOfSkipWords int) (string, error) {
	for i := 0; i < numberOfSkipWords; i++ {
		separatorIndex := strings.Index(curLine, " ")
		if separatorIndex == -1 {
			return curLine, nil
		}

		if separatorIndex < len(curLine)-separatorIndex-1 {
			curLine = curLine[separatorIndex+1:]
		} else {
			return "", errors.Wrap(errors.New("seperator index out of range"), "skipWords error:")
		}
	}

	return curLine, nil
}

func skipSymbols(curLine string, numberOfSkipSymbols int) string {
	if numberOfSkipSymbols <= len(curLine) {
		return curLine[numberOfSkipSymbols:]
	}

	return curLine
}

func processModifyingOptions(curString string, prevString string, numberOfSkipWords int, numberOfSkipChar int, register bool) (string, string, error) {
	var (
		err error
	)

	if register == true {
		curString = strings.ToLower(curString)
		prevString = strings.ToLower(prevString)
	}

	curString, err = skipWords(curString, numberOfSkipWords)
	if err != nil {
		return "", "", errors.Wrap(err, "processingModifyOptions function")
	}
	prevString, err = skipWords(prevString, numberOfSkipWords)
	if err != nil {
		return "", "", errors.Wrap(err, "processingModifyOptions function")
	}

	curString = skipSymbols(curString, numberOfSkipChar)
	if err != nil {
		return "", "", errors.Wrap(err, "processingModifyOptions function")
	}
	prevString = skipSymbols(prevString, numberOfSkipChar)
	if err != nil {
		return "", "", errors.Wrap(err, "processingModifyOptions function")
	}

	return curString, prevString, nil
}

func selectReader(writeBuffer []string, input io.Reader) *bufio.Scanner {
	if len(writeBuffer) != 0 {
		fmt.Println("Cur buffer: ", writeBuffer)
		curReader := bytes.NewReader(bytes.NewBufferString(strings.Join(writeBuffer, "\n")).Bytes())
		return bufio.NewScanner(curReader)
	} else {
		return bufio.NewScanner(input)
	}
}
