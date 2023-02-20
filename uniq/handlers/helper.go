package handler

import (
	"bufio"
	"bytes"
	"io"
	"strings"

	"github.com/pkg/errors"
)

func selectReader(writeBuffer []string, input io.Reader) *bufio.Scanner {
	if len(writeBuffer) != 0 {
		// writeBuffer = append(writeBuffer, "\n") // буффер не пустой - запись должна начаться с новой строки
		curReader := bytes.NewReader(bytes.NewBufferString(strings.Join(writeBuffer, "\n")).Bytes()) // тут скорее всего "\n"
		return bufio.NewScanner(curReader)
	} else {
		return bufio.NewScanner(input)
	}
}

func skipWords(curLine string, numberOfSkipWords int) (string, error) {
	for i := 0; i < numberOfSkipWords; i++ {
		separatorIndex := strings.Index(curLine, " ")
		// fmt.Println("CurLine: ", curLine)
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
