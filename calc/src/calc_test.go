package src

import (
	"errors"
	"math"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
	rnd "github.com/wadey/go-rounding"
)

const precision = 3

func round(received float64, prec int) (rounded float64) {
	x := new(big.Rat).SetFloat64(received)
	rnd.Round(x, prec, rnd.HalfUp)
	rounded, _ = x.Float64()
	return
}

func TestCalcPlus(t *testing.T) {
	cases := map[string]struct {
		in       string
		expected float64
	}{
		"success: simlpe operation": {
			in:       "2+2",
			expected: 4,
		},
		"success: 3 operands": {
			in:       "  2 + 2+ 4  ",
			expected: 2 + 2 + 4.0,
		},
		"success: 4 float digits": {
			in:       "1.93 + 3.233 + 6.129603 + 2.961735",
			expected: 1.93 + 3.233 + 6.129603 + 2.961735,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			received, err := Calc(tc.in)
			require.Equal(t, round(tc.expected, precision), round(received, precision))
			require.NoError(t, err)
		})
	}
}

func TestCalcMinus(t *testing.T) {
	cases := map[string]struct {
		in       string
		expected float64
	}{
		"success: 3 simlpe operands": {
			in:       "10.0 - 6 - 1",
			expected: 10.0 - 6 - 1,
		},
		"success: 3 float digits": {
			in:       "6.3291 - 8.28381 - 3.323214",
			expected: 6.3291 - 8.28381 - 3.323214,
		},
		"success: 2 small digits without first 0": {
			in:       ".93252   -.345262",
			expected: 0.93252 - 0.345262,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			received, err := Calc(tc.in)
			require.Equal(t, round(tc.expected, precision), round(received, precision))
			require.NoError(t, err)
		})
	}
}

func TestCalcMul(t *testing.T) {
	cases := map[string]struct {
		in       string
		expected float64
	}{
		"success: simlpe case": {
			in:       "2.5 * 4",
			expected: 2.5 * 4,
		},
		"success: 3 float digits": {
			in:       "12.366 * 2.26 * 0.3599",
			expected: 12.366 * 2.26 * 0.3599,
		},
		"success: small digits": {
			in:       "0.02 * 0.000398",
			expected: 0.02 * 0.000398,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			received, err := Calc(tc.in)
			require.Equal(t, round(tc.expected, precision), round(received, precision))
			require.NoError(t, err)
		})
	}
}

func TestCalcDiv(t *testing.T) {
	cases := map[string]struct {
		in       string
		expected float64
	}{
		"success: simlpe case": {
			in:       "10.0 / 4.0",
			expected: 10.0 / 4.0,
		},
		"success: 3 float digits": {
			in:       "623.2125 / 2.35 / 10.342325",
			expected: 623.2125 / 2.35 / 10.342325,
		},
		"success: with periodic result": {
			in:       "10.0 / 3.0",
			expected: 10.0 / 3.0,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			received, err := Calc(tc.in)
			require.Equal(t, round(tc.expected, precision), round(received, precision))
			require.NoError(t, err)
		})
	}
}

func TestCalcFunctional(t *testing.T) {
	cases := map[string]struct {
		in       string
		expected float64
	}{
		"success: plus-mul": {
			in:       "2 + 2*2.0",
			expected: 2 + 2*2.0,
		},
		"success: brackets-div-minus": {
			in:       "2/(2-2.8)",
			expected: 2 / (2 - 2.8),
		},
		"success: all operators, simple digits": {
			in:       "2*(2-2.0) + 15/(2+1)",
			expected: 2*(2-2.0) + 15/(2+1),
		},
		"success: all operators, embedded brackets, simple digits": {
			in:       "1.0 + 2*(3+4/2-(1+2))*2 + 1",
			expected: 1.0 + 2*(3+4/2-(1+2))*2 + 1,
		},
		"success: more complex experssion, all cases together": {
			in:       "345.2928 + (7374.4332-1.) - 847.958222*23.949032 - ((763.541 - .531) / 82.09093 * (522.00001 + 22.4373*88.146 - 115.7*2.323467))",
			expected: 345.2928 + (7374.4332 - 1.) - 847.958222*23.949032 - ((763.541 - .531) / 82.09093 * (522.00001 + 22.4373*88.146 - 115.7*2.323467)),
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			received, err := Calc(tc.in)
			require.Equal(t, round(tc.expected, precision), round(received, precision))
			require.NoError(t, err)
		})
	}
}

func TestCalcZeroMul(t *testing.T) {
	cases := map[string]struct {
		in       string
		expected float64
	}{
		"success: simple mul": {
			in:       "25*0/5.0",
			expected: 0.0,
		},
		"success: more compicated digits": {
			in:       "(1.9999999-1.9999999)/(0.3-0.29999)*999999999",
			expected: 0.0,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			received, err := Calc(tc.in)
			require.Equal(t, round(tc.expected, precision), round(received, precision))
			require.NoError(t, err)
		})
	}
}

func TestCalcZeroDiv(t *testing.T) {
	expected := math.Inf(-1)
	received, err := Calc("(2-3)/(2-2)")
	require.Equal(t, expected, received)
	require.NoError(t, err)
}

func TestCalcNaN(t *testing.T) {
	cases := map[string]struct {
		in       string
		expected float64
	}{
		"success: operation 0/0": {
			in:       "(1-1)/0.0",
			expected: math.NaN(),
		},
		"success: operation 0/0, more compicated digits": {
			in:       "(1.000001-1.000001)/(9.999999-9.999999)",
			expected: math.NaN(),
		},
		"success: operation Inf/Inf": {
			in:       "(1/0)/(2/0)",
			expected: math.NaN(),
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			received, err := Calc(tc.in)
			require.True(t, math.IsNaN(received))
			errors.Is(err, errors.New("is NaN"))
		})
	}
}

func TestCalcValidInput(t *testing.T) {
	cases := map[string]struct {
		in       string
		expected string
	}{
		"success: unknown operators": {
			in:       "abc*123",
			expected: "invalid input data",
		},
		"success: unknown brackets": {
			in:       "[2-1]/2",
			expected: "invalid input data",
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			_, err := Calc(tc.in)
			require.Equal(t, err.Error(), tc.expected)
		})
	}

}

func TestCreateTokkensFromString(t *testing.T) {
	cases := map[string]struct {
		in       string
		expected []string
	}{
		"success: any set of numbers and symbols utf-8": {
			in:       "   23+325   + 1-3.51(*1+3.))0.0",
			expected: []string{"23", "+", "325", "+", "1", "-", "3.51", "(", "*", "1", "+", "3.", ")", ")", "0.0"},
		},
		"success: any set of numbers and symbols utf-32": {
			in:       "шЛ8.23#MWzX]   🍺",
			expected: []string{"ш", "Л", "8.23", "#", "M", "W", "z", "X", "]", "🍺"},
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			received, err := getCharacters(tc.in)
			require.Equal(t, tc.expected, received)
			require.NoError(t, err)
		})
	}
}

func TestCalcReversePolishNotation(t *testing.T) {
	expected := 17.0
	received, err := calcReversePolishNotation([]string{"2.0", "5.0", "+", "3.0", "*"})
	require.Equal(t, round(expected, precision), round(received, precision))
	require.NoError(t, err)
}
