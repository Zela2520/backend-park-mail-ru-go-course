package route

import (
	"fmt"
	"strings"

	handler "github.com/Zela2520/backend-park-mail-ru-go-course.git/uniq/handlers"
	"github.com/Zela2520/backend-park-mail-ru-go-course.git/uniq/param"
	"github.com/pkg/errors"
)

func Route(options []param.Param) error {

	var writeBuffer []string

	err := checkBoolFlags(options)
	if err != nil {
		return errors.Wrap(err, "Route function")
	}

	input, output, err := getStreams()
	if err != nil {
		return errors.Wrap(err, "checkFileParm function error")
	}

	defer input.Close()
	defer output.Close()

	modifyingOptions := options[3:6] // -i -f -s
	optionsHandler := handler.NewHandler(modifyingOptions...)
	countOfActiveFlags := 0

	for _, val := range options[:3] {
		if val.OptionValue != false {
			countOfActiveFlags++
			writeBuffer, err = optionsHandler.HandleMap[val.Option](input, writeBuffer, options[4].OptionValue.(int), options[5].OptionValue.(int), options[3].OptionValue.(bool)) // TODO: убарть передачу функций, сделать все через конструктор
			if err != nil {
				return errors.Wrap(err, "handler error:"+val.OptionMessage)
			}
		}
	}

	if countOfActiveFlags == 0 {
		writeBuffer, err = handler.Uniq(input, writeBuffer, modifyingOptions[1].OptionValue.(int), modifyingOptions[2].OptionValue.(int), modifyingOptions[0].OptionValue.(bool))
		if err != nil {
			return errors.Wrap(err, "handler error: default case error")
		}
	}

	fmt.Fprintln(output, strings.Join(writeBuffer, "\n"))

	return nil
}
