package src

import (
	"bufio"
	"errors"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getCharacters(expression string) ([]string, error) {
	strWithoutSpaces := strings.ReplaceAll(expression, " ", "")
	res, err := regexp.Compile(`\d*\.\d*|\D|\d*`)
	if err != nil {
		return make([]string, 0), err
	}
	tokkens := res.FindAllString(strWithoutSpaces, -1)

	return tokkens, nil
}

func calcReversePolishNotation(charaters []string) (float64, error) {
	var (
		operators      = map[string]int{"+": 1, "-": 1, "*": 2, "/": 2, "(": -1, ")": -1}
		operationStack = Stack{}
		valuesStack    = Stack{}
	)

	for _, token := range charaters {
		number, parsingError := strconv.ParseFloat(token, 64)
		if parsingError == nil {
			valuesStack.Push(number)
			continue
		}

		_, isExist := operators[token]
		if !isExist {
			return 0, errors.New("invalid input data")
		}

		if operationStack.IsEmpty() || token == "(" {
			operationStack.Push(token)
			continue
		}

		curOperator := operationStack.Top().(string)
		if operators[curOperator] < operators[token] {
			operationStack.Push(token)
			continue
		}

		if token == ")" {
			for curOperator != "(" {
				tmp, err := getOperation(
					valuesStack.Pop().(float64),
					valuesStack.Pop().(float64),
					operationStack.Pop().(string))
				if err != nil {
					return 0, err
				}

				valuesStack.Push(tmp)

				if operationStack.IsEmpty() {
					return 0, errors.New("syntax Error")
				}

				curOperator = operationStack.Top().(string)
			}

			operationStack.Pop()

			continue
		}

		for curOperator != "(" && operators[curOperator] >= operators[token] {
			tmp, err := getOperation(
				valuesStack.Pop().(float64),
				valuesStack.Pop().(float64),
				operationStack.Pop().(string))
			if err != nil {
				return 0, err
			}
			valuesStack.Push(tmp)
			if operationStack.IsEmpty() {
				break
			}

			curOperator = operationStack.Top().(string)
		}
		operationStack.Push(token)
	}

	for !operationStack.IsEmpty() {
		tmp, err := getOperation(
			valuesStack.Pop().(float64),
			valuesStack.Pop().(float64),
			operationStack.Pop().(string))
		if err != nil {
			return 0, err
		}
		valuesStack.Push(tmp)
	}

	if math.IsNaN(valuesStack.Top().(float64)) {
		return valuesStack.Top().(float64), errors.New("is NaN")
	}
	return valuesStack.Top().(float64), nil
}

func getOperation(value1 float64, value2 float64, operation string) (float64, error) {
	switch operation {
	case "+":
		return value2 + value1, nil
	case "-":
		return value2 - value1, nil
	case "*":
		return value2 * value1, nil
	case "/":
		return value2 / value1, nil
	case ")":
		return 0, errors.New("syntax Error")
	case "(":
		return 0, errors.New("syntax Error")
	default:
		return 0, errors.New("unknown operator")
	}
}

func GetInput() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func Calc(expression string) (float64, error) {
	charaters, err := getCharacters(expression)
	if err != nil {
		return 0, err
	}
	return calcReversePolishNotation(charaters)
}
