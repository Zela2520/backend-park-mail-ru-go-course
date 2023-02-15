package check

import (
	"fmt"
	"os"

	handler "github.com/Zela2520/backend-park-mail-ru-go-course.git/uniq/handlers"
	param "github.com/Zela2520/backend-park-mail-ru-go-course.git/uniq/param"
)

// будем  передавать срез из основного слайса параметров. Нет тут чекаем опции bool, потом опции int, опции ioput
// func checkBoolParams(boolParam ...param.Param) {
// 	for _, val := range boolParam {
// 		os.Stdout = boolHandler(os.Stdin, os.Stdout, val) // if для первых трех параметров
// 		os.Stdout = boolHandler(os.Stdin, os.Stdout, val) // второй if глупо делать наверное, поэтому внутри хендлеров надо будет разбить на главные хендлеры и вспомагательные
// 	}
// }

// GetHandle | возвращать надо строку, которая собственно будет равна флагу
// В этой функции как раз-таки можно валидировать параметры.
func Check(paramList []param.Param) error {
	for _, val := range paramList {
		var curVal interface{} = val.OptionValue
		switch v := curVal.(type) {
		case bool:
			fmt.Println("bool. Need to call boolHandler:", v) // будет три главных функции handling-a. С
		case int:
			fmt.Println("int. Need to call intHanler:", v)
		default:
			fmt.Println("string. Need to call StdStream handler", v)
		}
	}
	return nil
}

func getHandler(params []param.Param) {
	for _, val := range params {
		fmt.Println(val)
	}

	handler.CountUniq(os.Stdin, os.Stdout) // что передавать в handler, продумать структуру handler-a
}
