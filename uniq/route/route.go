package route

import (
	"fmt"

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

	for _, val := range options {
		switch v := val.OptionValue.(type) {
		case bool:
			{
				if val.OptionValue != false {
					optionsHandler.HandleMap[val.Option](input, output, val.OptionValue)
					fmt.Println(val.OptionMessage, v) // отладка
				}
			}

		case int:
			{
				if val.OptionValue != 0 {
					optionsHandler.HandleMap[val.Option](input, output, val.OptionValue)
					fmt.Println(val.OptionMessage, v) // отладка
				}
			}

		default: // Значит хотя бы один file был передан. Или input или output
			{
				if val.OptionValue != "" {
					fmt.Println("Need to call handler. Value", val.OptionValue)
				}
			}
		}
	}

	return nil
}
