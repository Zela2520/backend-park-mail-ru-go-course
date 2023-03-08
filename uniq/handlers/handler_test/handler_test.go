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
					"3 I love music.",
					"1 ",
					"2 I love music of Kartik.",
					"1 Thanks.",
					"2 I love music of Kartik.",
				}, ""),
			},
			expErr: "Without option test failed",
		},
		{
			name: "With register option",
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
				ExpectedData: "3 I LOVE MUSIC.1 2 I love MuSIC of Kartik.1 Thanks.2 I love music of kartik.",
			},
			expErr: "With register option test failed",
		},
		{
			name: "With integer options",
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
				ExpectedData: "3 I love music.1 2 I Home music of Kartik.1 Thanks.",
			},
			expErr: "Unic count with integer options test failed",
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			r := strings.NewReader(tCase.in.InitData)
			writeBuffer, err = handler.CountUniq(r, writeBuffer, tCase.in.ModifyOptions[0], tCase.in.ModifyOptions[1], tCase.in.Register)
			if err != nil {
				t.Errorf("CountUniq method error: %s", tCase.expErr)
			}

			output := strings.Join(writeBuffer, "")
			require.NotEqual(t, len(output), 0, tCase.expErr)

			require.Equal(t, tCase.in.ExpectedData, output, "should be equal")
			writeBuffer = writeBuffer[:0]
		})
	}
}

func TestRepeatedLines(t *testing.T) {
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
					"I love music of Kartik.",
					"I love music of Kartik.",
				}, ""),
			},
			expErr: "Without option test failed",
		},
		{
			name: "With register option",
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
					"I love MuSIC of Kartik.",
					"I love music of kartik.",
				}, ""),
			},
			expErr: "With register option test failed",
		},
		{
			name: "With integer options",
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
					"THanks.",
				}, "\n"),
				ExpectedData: strings.Join([]string{"I love music.",
					"I Home music of Kartik.",
					"Thanks.",
				}, ""),
			},
			expErr: "Unic count with integer options test failed",
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			r := strings.NewReader(tCase.in.InitData)
			writeBuffer, err = handler.GetRepeatedLines(r, writeBuffer, tCase.in.ModifyOptions[0], tCase.in.ModifyOptions[1], tCase.in.Register)
			if err != nil {
				t.Errorf("CountUniq method error: %s", tCase.expErr)
			}

			output := strings.Join(writeBuffer, "")
			require.NotEqual(t, len(output), 0, tCase.expErr)

			require.Equal(t, tCase.in.ExpectedData, output, "should be equal")
			writeBuffer = writeBuffer[:0]
		})
	}
}

func TestGetNotRepeatedLines(t *testing.T) {
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
					"",
					"Thanks.",
				}, ""),
			},
			expErr: "Without option test failed",
		},
		{
			name: "With register option",
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
					"",
					"Thanks.",
				}, ""),
			},
			expErr: "With register option test failed",
		},
		{
			name: "With integer options",
			in: &Case{
				Register:      true,
				ModifyOptions: intOptions[2],
				InitData: strings.Join([]string{
					"I love music.",
					"A LoVe music.",
					"C lOvE music.",
					"",
					"Thanks.",
					"I Home music of Kartik.",
					"We Lome music of Kartik.",
					"Pilot samoleta.",
				}, "\n"),
				ExpectedData: strings.Join([]string{
					"",
					"Thanks.",
					"Pilot samoleta.",
				}, ""),
			},
			expErr: "Unic count with integer options test failed",
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			r := strings.NewReader(tCase.in.InitData)
			writeBuffer, err = handler.GetNotRepeatedLines(r, writeBuffer, tCase.in.ModifyOptions[0], tCase.in.ModifyOptions[1], tCase.in.Register)
			if err != nil {
				t.Errorf("CountUniq method error: %s", tCase.expErr)
			}

			output := strings.Join(writeBuffer, "")
			require.NotEqual(t, len(output), 0, tCase.expErr)

			require.Equal(t, tCase.in.ExpectedData, output, "should be equal")
			writeBuffer = writeBuffer[:0]
		})
	}
}
