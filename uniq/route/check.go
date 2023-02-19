package route

import (
	"flag"
	"os"

	param "github.com/Zela2520/backend-park-mail-ru-go-course.git/uniq/param"
	"github.com/pkg/errors"
)

func checkBoolFlags(paramList []param.Param) error {
	trueFlag := 0
	for _, val := range paramList[:3] {
		if val.OptionValue == true {
			trueFlag++
		}
	}

	if trueFlag > 1 {
		return errors.Wrap(errors.New("CheckBool"), "CheckBool function")
	}

	return nil
}

func checkFileParam() (*os.File, *os.File, error) {
	var (
		input  *os.File
		output *os.File
		err    error
	)

	streamHandler := func(filePath string) (*os.File, error) {
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			return file, errors.Wrap(err, "file does not exist or cannot be opened")
		}
		return file, nil
	}

	for i, val := range flag.Args() {
		switch i {
		case 0:
			{
				input, err = streamHandler(val)
				if err != nil {
					return nil, nil, errors.Wrap(err, "input file incorrected")
				}
			}

		case 1:
			{
				output, err = streamHandler(val)
				if err != nil {
					return nil, nil, errors.Wrap(err, "ouput file incorrected")
				}
			}
		}
	}

	if input == nil {
		input = os.Stdin
	}

	if output == nil {
		output = os.Stdout
	}

	return input, output, nil
}
