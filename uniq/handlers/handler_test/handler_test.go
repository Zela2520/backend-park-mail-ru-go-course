package handler_test

import (
	"strings"
	"testing"

	handler "github.com/Zela2520/backend-park-mail-ru-go-course.git/uniq/handlers"
	"github.com/stretchr/testify/require"
)

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

func TestRepeatedLines(t *testing.T) {
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
		"sddsds",
		"sdsdsd",
		"",
		"",
		"111",
		"111",
		"",
		"",
		"1",
		"1",
	}, "\n")

	expectedData := strings.Join([]string{
		"I love music.",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"",
		"111",
		"",
		"1",
	}, "")

	r := strings.NewReader(initData)

	writeBuffer, err = handler.GetRepeatedLines(r, true, writeBuffer)
	if err != nil {
		t.Errorf("Uniq method error: %s", "")
	}

	output := strings.Join(writeBuffer, "")

	if len(output) == 0 {
		t.Errorf("Uniq method error: %s", "")
	}

	require.Equal(t, expectedData, output, "should be equal")
}

func TestGetNotRepeatedLines(t *testing.T) {
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
		"",
		"Thanks.",
	}, "")

	r := strings.NewReader(initData)

	writeBuffer, err = handler.GetNotRepeatedLines(r, true, writeBuffer)
	if err != nil {
		t.Errorf("Uniq method error: %s", "")
	}

	output := strings.Join(writeBuffer, "")

	if len(output) == 0 {
		t.Errorf("Uniq method error: %s", "")
	}

	require.Equal(t, expectedData, output, "should be equal")
}

func TestGetLinesWithoutRegister(t *testing.T) {
	var (
		writeBuffer []string
		err         error
	)

	initData := strings.Join([]string{
		"I LOVE MUSIC.",
		"I love music.",
		"I LoVe MuSiC.",

		"I love MuSIC of Kartik.",
		"I love music of kartik.",
		"Thanks.",
		"I love music of kartik.",
		"I love MuSIC of Kartik.",
	}, "\n")

	expectedData := strings.Join([]string{
		"I LOVE MUSIC.",
		"",
		"I love MuSIC of Kartik.",
		"Thanks.",
		"I love music of kartik.",
	}, "")

	r := strings.NewReader(initData)

	writeBuffer, err = handler.GetLinesWithoutRegister(r, true, writeBuffer)
	if err != nil {
		t.Errorf("Uniq method error: %s", "")
	}

	output := strings.Join(writeBuffer, "")

	if len(output) == 0 {
		t.Errorf("Uniq method error: %s", "")
	}

	require.Equal(t, expectedData, output, "should be equal")
}

func TestGetLinesCompareNWord(t *testing.T) {
	var (
		writeBuffer []string
		err         error
	)

	initData := strings.Join([]string{
		"We love music.",
		"I love music.",
		"They love music.",
		"",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}, "\n")

	initDataComareWord := 1

	expectedData := strings.Join([]string{
		"We love music.",
		"",
		"I love music of Kartik.",
		"Thanks.",
	}, "")

	r := strings.NewReader(initData)

	writeBuffer, err = handler.GetLinesCompareNWord(r, initDataComareWord, writeBuffer)
	if err != nil {
		t.Errorf("Uniq method error: %s", "")
	}

	output := strings.Join(writeBuffer, "")

	if len(output) == 0 {
		t.Errorf("Uniq method error: %s", "")
	}

	require.Equal(t, expectedData, output, "should be equal")
}
