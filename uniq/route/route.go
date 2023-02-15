package route

import (
	"fmt"
	"os"

	"github.com/Zela2520/backend-park-mail-ru-go-course.git/uniq"
	handler "github.com/Zela2520/backend-park-mail-ru-go-course.git/uniq/handlers"
)

// будем  передавать срез из основного слайса параметров
// func checkBoolParams(boolParam ...[]uniq.Param) {
// 	for _, val := range boolParam {
// 		os.Stdout = boolHandler(os.Stdin, os.Stdout, val) // if для первых трех параметров
// 		os.Stdout = boolHandler(os.Stdin, os.Stdout, val) // второй if глупо делать наверное, поэтому внутри хендлеров надо будет разбить на главные хендлеры и вспомагательные
// 	}
// }

// GetHandle | возвращать надо строку, которая собственно будет равна флагу
// В этой функции как раз-таки можно валидировать параметры.
func Route(paramList []uniq.Param) {
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
}

func getHandler(params []uniq.Param) {
	for _, val := range params {
		fmt.Println(val)
	}

	handler.CountUniq(os.Stdin, os.Stdout) // что передавать в handler, продумать структуру handler-a
}
