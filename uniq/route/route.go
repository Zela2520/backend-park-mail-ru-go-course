package route

import (
	handler "github.com/Zela2520/backend-park-mail-ru-go-course.git/uniq/handlers"
	"github.com/Zela2520/backend-park-mail-ru-go-course.git/uniq/param"
	"github.com/pkg/errors"
)

func Route(options []param.Param) error {
	err := checkBoolFlags(options)
	if err != nil {
		return errors.Wrap(err, "Route function")
	}

	input, output, err := checkFileParam()
	if err != nil {
		return errors.Wrap(err, "checkFileParm function error")
	}

	optionsHandler := handler.NewHandler()

	countOfFiles := 0

	for _, val := range options {
		switch paramValue := val.OptionValue.(type) {
		case bool:
			{
				if val.OptionValue != false {
					optionsHandler.HandleMap[val.Option](input, output, paramValue)
					// fmt.Println(val.OptionMessage, v) // отладка
				}
			}

		case int:
			{
				if val.OptionValue != 0 {
					optionsHandler.HandleMap[val.Option](input, output, paramValue)
					// fmt.Println(val.OptionMessage, v) // отладка
				}
			}

		default:
			{
				countOfFiles++
				// все что ниже можно убрать
				// if val.OptionValue != "" {
				// 	if len(flag.Args()) == 1 {
				// 		handler.CountUniq(input, output, nil) // defer.close()
				// 		defer input.Close()
				// 	}
				// 	if len(flag.Args()) == countOfFiles && countOfFiles == 2 {
				// 		fmt.Println("Need to call handler input output. Value", paramValue)
				// 		handler.CountUniq(input, output, nil) // отладка
				// 		defer output.Close()
				// 	}
				// }
			}
		}
	}

	if countOfFiles == 0 {
		handler.Uniq(input, output, nil)
	}

	return nil
}
