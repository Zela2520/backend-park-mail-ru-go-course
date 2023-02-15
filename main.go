package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	check "github.com/Zela2520/backend-park-mail-ru-go-course.git/uniq/check"
	handler "github.com/Zela2520/backend-park-mail-ru-go-course.git/uniq/handlers"
	param "github.com/Zela2520/backend-park-mail-ru-go-course.git/uniq/param"
)

func main() {
	var (
		input  string
		output string
		err    error
	)

	options, err := param.GetParams()
	if err != nil {
		log.Fatal(errors.Unwrap(err))
	}

	// че должно сюда вернуться ? Было бы пиздато, если бы вернулся указатель на какой-нибудь handler.
	// А ну в принипе можно возвращать строку, а в самом handler сделать мапу.
	err = check.CheckBoolFlags(options)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i, val := range flag.Args() {
		switch i {
		case 0:
			input = val
		case 1:
			output = val
		}
	}

	// все это вынести в check пакет, хотя это route по сути
	for _, val := range options {
		switch v := val.OptionValue.(type) {
		case bool:
			{
				if val.OptionValue != false {
					fmt.Println("Need to call handler. Value", v) // вызываем обработчик для заданного флага
				}
			}

		case int:
			{
				if val.OptionValue != 0 {
					fmt.Println("Need to call handler. Value", v) // вызываем обработчик для заданного флага
				}
			}

		default:
			{
				if val.OptionValue != "" { // этому хендлеру передавать в любом случае input и output
					if input == "" {
						// input = os.Stdin()
					}
					if output == "" {

					}
					fmt.Println("Need to call handler. Value", v) // вызываем обработчик для заданного флага map[value.Option](arg1, arg2)
				}
			}
		}
	}

	// после того как прочекали опции идем по массиву
	// и вызываем соотвествующие хендлеры, обращаясь к мапе
	// вызываем ток то что true, то есть имеет не дефолтные значения, через if-ик можно чекнуть
	// можно как сразу вызывать хендлеры так и просто говорить, что вот по таким то ключам они есть, хотя это глупо
	// лучше сразу вызывать обращаться по ключу мапы, чтобы вызвался хендлер и вернул результат своей работы
	// но куда он будет возвращать результат своей работы ??? думаю в поток либо создать строку и посмотреть что будет.
	// (но лучше просто поток по указателю везде таскать)

	err = handler.CountUniq(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
