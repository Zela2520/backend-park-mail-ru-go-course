package handler

import (
	"bufio"
	"bytes"
	"io"
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
