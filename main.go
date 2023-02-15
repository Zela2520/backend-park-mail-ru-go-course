package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	check "github.com/Zela2520/backend-park-mail-ru-go-course.git/uniq/check"
	handler "github.com/Zela2520/backend-park-mail-ru-go-course.git/uniq/handlers"
	param "github.com/Zela2520/backend-park-mail-ru-go-course.git/uniq/param"
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

	options, err := param.GetParams()
	if err != nil {
		log.Fatal(errors.Unwrap(err))
	}

	// че должно сюда вернуться ? Было бы пиздато, если бы вернулся указатель на какой-нибудь handler.
	// А ну в принипе можно возвращать строку, а в самом handler сделать мапу.
	err = check.Check(options)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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
