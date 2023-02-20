package route

import (
	"fmt"
	"strings"

	handler "github.com/Zela2520/backend-park-mail-ru-go-course.git/uniq/handlers"
	"github.com/Zela2520/backend-park-mail-ru-go-course.git/uniq/param"
	"github.com/pkg/errors"
)

func Route(options []param.Param) error {

	var (
		writeBuffer []string
	)

	err := checkBoolFlags(options)
	if err != nil {
		return errors.Wrap(err, "Route function")
	}

	input, output, err := checkFileParam()
	if err != nil {
		return errors.Wrap(err, "checkFileParm function error")
	}

	defer input.Close()
	defer output.Close()

	optionsHandler := handler.NewHandler()

	countOfActiveFlags := 0

	for _, val := range options {
		switch paramValue := val.OptionValue.(type) {
		case bool:
			{
				if val.OptionValue != false {
					countOfActiveFlags++
					writeBuffer, err = optionsHandler.HandleMap[val.Option](input, paramValue, writeBuffer)
					if err != nil {
						return errors.Wrap(err, "handler error:"+val.OptionMessage)
					}
				}
			}

		case int:
			{
				if val.OptionValue != 0 {
					countOfActiveFlags++
					writeBuffer, err = optionsHandler.HandleMap[val.Option](input, val.OptionValue, writeBuffer)
					if err != nil {
						return errors.Wrap(err, "handler error:"+val.OptionMessage)
					}
				}
			}
		}
	}

	if countOfActiveFlags == 0 {
		handler.Uniq(input, output)
	} else {
		fmt.Fprintln(output, strings.Join(writeBuffer, ""))
	}

	return nil
}
