package handler_test

import (
	"strings"
	"testing"

	handler "github.com/Zela2520/backend-park-mail-ru-go-course.git/uniq/handlers"
	"github.com/stretchr/testify/require"
)

func TestCountUniq(t *testing.T) {
	var (
		writeBuffer []string
		err         error
	)

	initData := strings.Join([]string{
		"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik.",
		"I love music of Kartik.",
	}, "\n")

	expectedData := strings.Join([]string{
		"3 I love music.",
		"1 ",
		"2 I love music of Kartik.",
		"1 Thanks.",
		"2 I love music of Kartik.",
	}, "")

	r := strings.NewReader(initData)

	writeBuffer, err = handler.CountUniq(r, true, writeBuffer)
	if err != nil {
		t.Errorf("CountUniq method error: %s", "")
	}

	output := strings.Join(writeBuffer, "")

	if len(output) == 0 {
		t.Errorf("CountUniq method error: %s", "")
	}

	require.Equal(t, expectedData, output, "should be equal")
}

func TestUniq(t *testing.T) {
	var (
		writeBuffer []string
		err         error
	)

	initData := strings.Join([]string{
		"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik.",
		"I love music of Kartik.",
	}, "\n")

	expectedData := strings.Join([]string{
		"I love music.",
		"",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik.",
	}, "")

	r := strings.NewReader(initData)

	writeBuffer, err = handler.Uniq(r, writeBuffer)
	if err != nil {
		t.Errorf("Uniq method error: %s", "")
	}

	output := strings.Join(writeBuffer, "")

	if len(output) == 0 {
		t.Errorf("Uniq method error: %s", "")
	}

	require.Equal(t, expectedData, output, "should be equal")
}
