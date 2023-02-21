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

	type Case struct {
		Register      bool
		ModifyOptions []int
		InitData      string
		ExpectedData  string
	}

	intOptions := [][]int{
		{1, 0},
		{0, 1},
		{1, 1},
	}

	cases := []struct {
		name   string
		in     *Case
		expErr string
	}{
		{
			name: "Uniq without options",
			in: &Case{
				Register:      false,
				ModifyOptions: make([]int, 2),
				InitData: strings.Join([]string{
					"I love music.",
					"I love music.",
					"I love music.",
					"",
					"I love music of Kartik.",
					"I love music of Kartik.",
					"Thanks.",
					"I love music of Kartik.",
					"I love music of Kartik.",
				}, "\n"),
				ExpectedData: strings.Join([]string{
					"I love music.",
					"",
					"I love music of Kartik.",
					"Thanks.",
					"I love music of Kartik.",
				}, ""),
			},
			expErr: "Uniq test failed",
		},
		{
			name: "GetLinesWithoutRegister",
			in: &Case{
				Register:      true,
				ModifyOptions: make([]int, 2),
				InitData: strings.Join([]string{
					"I LOVE MUSIC.",
					"I love music.",
					"I LoVe MuSiC.",
					"",
					"I love MuSIC of Kartik.",
					"I love music of kartik.",
					"Thanks.",
					"I love music of kartik.",
					"I love MuSIC of Kartik.",
				}, "\n"),
				ExpectedData: strings.Join([]string{
					"I LOVE MUSIC.",
					"",
					"I love MuSIC of Kartik.",
					"Thanks.",
					"I love music of kartik.",
				}, ""),
			},
			expErr: "GetLinesWithoutRegister test failed",
		},
		{
			name: "GetLinesCompareNWord",
			in: &Case{
				Register:      false,
				ModifyOptions: intOptions[0],
				InitData: strings.Join([]string{
					"We love music.",
					"I love music.",
					"They love music.",
					"",
					"I love music of Kartik.",
					"We love music of Kartik.",
					"Thanks.",
				}, "\n"),
				ExpectedData: strings.Join([]string{
					"We love music.",
					"",
					"I love music of Kartik.",
					"Thanks.",
				}, ""),
			},
			expErr: "GetLinesCompareNWord test failed",
		},
		{
			name: "GetLinesCompareNChar",
			in: &Case{
				Register:      false,
				ModifyOptions: intOptions[1],
				InitData: strings.Join([]string{
					"I love music.",
					"A love music.",
					"C love music.",
					"",
					"I love music of Kartik.",
					"We love music of Kartik.",
					"Thanks.",
				}, "\n"),
				ExpectedData: strings.Join([]string{
					"I love music.",
					"",
					"I love music of Kartik.",
					"We love music of Kartik.",
					"Thanks.",
				}, ""),
			},
			expErr: "GetLinesCompareNChar test failed",
		},
		{
			name: "СombinationOfChangeableOptions",
			in: &Case{
				Register:      true,
				ModifyOptions: intOptions[2],
				InitData: strings.Join([]string{
					"I love music.",
					"A LoVe music.",
					"C lOvE music.",
					"",
					"I Home music of Kartik.",
					"We Lome music of Kartik.",
					"Thanks.",
				}, "\n"),
				ExpectedData: strings.Join([]string{
					"I love music.",
					"",
					"I Home music of Kartik.",
					"Thanks.",
				}, ""),
			},
			expErr: "GetLinesCompareNChar test failed",
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			r := strings.NewReader(tCase.in.InitData)
			writeBuffer, err = handler.Uniq(r, writeBuffer, tCase.in.ModifyOptions[0], tCase.in.ModifyOptions[1], tCase.in.Register)
			if err != nil {
				t.Errorf("Uniq method error: %s", "")
			}

			output := strings.Join(writeBuffer, "")
			require.NotEqual(t, len(output), 0, tCase.expErr)

			require.Equal(t, tCase.in.ExpectedData, output, "should be equal")
			writeBuffer = writeBuffer[:0]
		})
	}
}

func TestIoControl(t *testing.T) {

}

func TestCountUniq(t *testing.T) {
	var (
		writeBuffer []string
		err         error
	)

	type Case struct {
		Register      bool
		ModifyOptions []int
		InitData      string
		ExpectedData  string
	}

	intOptions := [][]int{
		{1, 0},
		{0, 1},
		{1, 1},
	}

	cases := []struct {
		name   string
		in     *Case
		expErr string
	}{
		{
			name: "Without options",
			in: &Case{
				Register:      false,
				ModifyOptions: make([]int, 2),
				InitData: strings.Join([]string{
					"I love music.",
					"I love music.",
					"I love music.",
					"",
					"I love music of Kartik.",
					"I love music of Kartik.",
					"Thanks.",
					"I love music of Kartik.",
					"I love music of Kartik.",
				}, "\n"),
				ExpectedData: strings.Join([]string{
					"I love music.",
					"",
					"I love music of Kartik.",
					"Thanks.",
					"I love music of Kartik.",
				}, ""),
			},
			expErr: "Uniq test failed",
		},
		{
			name: "Negative case",
			in: &Case{
				Register:      true,
				ModifyOptions: make([]int, 2),
				InitData: strings.Join([]string{
					"I LOVE MUSIC.",
					"I love music.",
					"I LoVe MuSiC.",
					"",
					"I love MuSIC of Kartik.",
					"I love music of kartik.",
					"Thanks.",
					"I love music of kartik.",
					"I love MuSIC of Kartik.",
				}, "\n"),
				ExpectedData: strings.Join([]string{
					"I LOVE MUSIC.",
					"",
					"I love MuSIC of Kartik.",
					"Thanks.",
					"I love music of kartik.",
				}, ""),
			},
			expErr: "GetLinesWithoutRegister test failed",
		},
		{
			name: "With register option",
			in: &Case{
				Register:      false,
				ModifyOptions: intOptions[0],
				InitData: strings.Join([]string{
					"We love music.",
					"I love music.",
					"They love music.",
					"",
					"I love music of Kartik.",
					"We love music of Kartik.",
					"Thanks.",
				}, "\n"),
				ExpectedData: strings.Join([]string{
					"We love music.",
					"",
					"I love music of Kartik.",
					"Thanks.",
				}, ""),
			},
			expErr: "GetLinesCompareNWord test failed",
		},
		{
			name: "With integet options",
			in: &Case{
				Register:      false,
				ModifyOptions: intOptions[0],
				InitData: strings.Join([]string{
					"We love music.",
					"I love music.",
					"They love music.",
					"",
					"I love music of Kartik.",
					"We love music of Kartik.",
					"Thanks.",
				}, "\n"),
				ExpectedData: strings.Join([]string{
					"We love music.",
					"",
					"I love music of Kartik.",
					"Thanks.",
				}, ""),
			},
			expErr: "GetLinesCompareNWord test failed",
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			r := strings.NewReader(tCase.in.InitData)
			writeBuffer, err = handler.CountUniq(r, writeBuffer, tCase.in.ModifyOptions[0], tCase.in.ModifyOptions[1], tCase.in.Register)
			if err != nil {
				t.Errorf("Uniq method error: %s", "")
			}

			output := strings.Join(writeBuffer, "")
			require.NotEqual(t, len(output), 0, tCase.expErr)

			require.Equal(t, tCase.in.ExpectedData, output, "should be equal")
			writeBuffer = writeBuffer[:0]
		})
	}
}

// func TestRepeatedLines(t *testing.T) {
// 	var (
// 		writeBuffer []string
// 		err         error
// 	)

// 	initData := strings.Join([]string{
// 		"I love music.",
// 		"I love music.",
// 		"I love music.",
// 		"",
// 		"I love music of Kartik.",
// 		"I love music of Kartik.",
// 		"Thanks.",
// 		"I love music of Kartik.",
// 		"I love music of Kartik.",
// 		"sddsds",
// 		"sdsdsd",
// 		"",
// 		"",
// 		"111",
// 		"111",
// 		"",
// 		"",
// 		"1",
// 		"1",
// 	}, "\n")

// 	expectedData := strings.Join([]string{
// 		"I love music.",
// 		"I love music of Kartik.",
// 		"I love music of Kartik.",
// 		"",
// 		"111",
// 		"",
// 		"1",
// 	}, "")

// 	r := strings.NewReader(initData)

// 	writeBuffer, err = handler.GetRepeatedLines(r, true, writeBuffer)
// 	if err != nil {
// 		t.Errorf("Uniq method error: %s", "")
// 	}

// 	output := strings.Join(writeBuffer, "")

// 	if len(output) == 0 {
// 		t.Errorf("Uniq method error: %s", "")
// 	}

// 	require.Equal(t, expectedData, output, "should be equal")
// }

// func TestGetNotRepeatedLines(t *testing.T) {
// 	var (
// 		writeBuffer []string
// 		err         error
// 	)

// 	initData := strings.Join([]string{
// 		"I love music.",
// 		"I love music.",
// 		"I love music.",
// 		"",
// 		"I love music of Kartik.",
// 		"I love music of Kartik.",
// 		"Thanks.",
// 		"I love music of Kartik.",
// 		"I love music of Kartik.",
// 	}, "\n")

// 	expectedData := strings.Join([]string{
// 		"",
// 		"Thanks.",
// 	}, "")

// 	r := strings.NewReader(initData)

// 	writeBuffer, err = handler.GetNotRepeatedLines(r, true, writeBuffer)
// 	if err != nil {
// 		t.Errorf("Uniq method error: %s", "")
// 	}

// 	output := strings.Join(writeBuffer, "")

// 	if len(output) == 0 {
// 		t.Errorf("Uniq method error: %s", "")
// 	}

// 	require.Equal(t, expectedData, output, "should be equal")
// }

// func TestGetLinesWithoutRegister(t *testing.T) {
// 	var (
// 		writeBuffer []string
// 		err         error
// 	)

// 	initData := strings.Join([]string{
// 		"I LOVE MUSIC.",
// 		"I love music.",
// 		"I LoVe MuSiC.",

// 		"I love MuSIC of Kartik.",
// 		"I love music of kartik.",
// 		"Thanks.",
// 		"I love music of kartik.",
// 		"I love MuSIC of Kartik.",
// 	}, "\n")

// 	expectedData := strings.Join([]string{
// 		"I LOVE MUSIC.",
// 		"",
// 		"I love MuSIC of Kartik.",
// 		"Thanks.",
// 		"I love music of kartik.",
// 	}, "")

// 	r := strings.NewReader(initData)

// 	writeBuffer, err = handler.GetLinesWithoutRegister(r, true, writeBuffer)
// 	if err != nil {
// 		t.Errorf("Uniq method error: %s", "")
// 	}

// 	output := strings.Join(writeBuffer, "")

// 	if len(output) == 0 {
// 		t.Errorf("Uniq method error: %s", "")
// 	}

// 	require.Equal(t, expectedData, output, "should be equal")
// }

// func TestGetLinesCompareNWord(t *testing.T) {
// 	var (
// 		writeBuffer []string
// 		err         error
// 	)

// 	initData := strings.Join([]string{
// 		"We love music.",
// 		"I love music.",
// 		"They love music.",
// 		"",
// 		"I love music of Kartik.",
// 		"We love music of Kartik.",
// 		"Thanks.",
// 	}, "\n")

// 	initDataComareWord := 1

// 	expectedData := strings.Join([]string{
// 		"We love music.",
// 		"",
// 		"I love music of Kartik.",
// 		"Thanks.",
// 	}, "")

// 	r := strings.NewReader(initData)

// 	writeBuffer, err = handler.GetLinesCompareNWord(r, initDataComareWord, writeBuffer)
// 	if err != nil {
// 		t.Errorf("Uniq method error: %s", "")
// 	}

// 	output := strings.Join(writeBuffer, "")

// 	if len(output) == 0 {
// 		t.Errorf("Uniq method error: %s", "")
// 	}

// 	require.Equal(t, expectedData, output, "should be equal")
// }

// func TestGetLinesCompareNChar(t *testing.T) {
// 	var (
// 		writeBuffer []string
// 		err         error
// 	)

// 	initData := strings.Join([]string{
// 		"I love music.",
// 		"A love music.",
// 		"C love music.",
// 		"",
// 		"I love music of Kartik.",
// 		"We love music of Kartik.",
// 		"Thanks.",
// 	}, "\n")

// 	initDataComareWord := 1

// 	expectedData := strings.Join([]string{
// 		"I love music.",
// 		"",
// 		"I love music of Kartik.",
// 		"We love music of Kartik.",
// 		"Thanks.",
// 	}, "")

// 	r := strings.NewReader(initData)

// 	writeBuffer, err = handler.GetLinesCompareNChar(r, initDataComareWord, writeBuffer)
// 	if err != nil {
// 		t.Errorf("Uniq method error: %s", "")
// 	}

// 	output := strings.Join(writeBuffer, "")

// 	if len(output) == 0 {
// 		t.Errorf("Uniq method error: %s", "")
// 	}

// 	require.Equal(t, expectedData, output, "should be equal")
// }
