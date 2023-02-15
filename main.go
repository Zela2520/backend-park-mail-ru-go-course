package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Zela2520/backend-park-mail-ru-go-course.git/uniq"
	handler "github.com/Zela2520/backend-park-mail-ru-go-course.git/uniq/handlers"
)

func main() {

	// - создадим объект uniq.NewUniq() *Uniq
	// - data := uniq.ReadData(os.Stdin) // можно сделать так, чтобы в струкуре uniq было поле buffer
	// - сделать switch case для выбора handler-a
	// - res := handler.Hanler1(data) // вместо data можно передать uniq.Data
	// - uniq.WrireData(res, os.Stdout)

	// 1) считать входные данные
	// 2) обработать входные данные
	// 3) вернуть данные

	var err error

	options, err := uniq.GetParams()
	if err != nil {
		log.Fatal(errors.Unwrap(err))
	}

	// отладка флагов
	for _, val := range options {
		fmt.Println(val.OptionValue)
	}

	err = handler.CountUniq(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
