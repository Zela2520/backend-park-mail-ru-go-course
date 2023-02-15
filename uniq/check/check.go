package check

import (
	param "github.com/Zela2520/backend-park-mail-ru-go-course.git/uniq/param"
	"github.com/pkg/errors"
)

// GetHandle | возвращать надо строку, которая собственно будет равна флагу
// В этой функции как раз-таки можно валидировать параметры.
func CheckBoolFlags(paramList []param.Param) error {
	boolParams := paramList[:3]
	trueFlag := 0
	for _, val := range boolParams {
		if val.OptionValue == true {
			trueFlag++
		}
	}

	if trueFlag > 1 {
		return errors.Wrap(errors.New("CheckBool"), "CheckBool function")
	}

	return nil
}

// func CheckInt(paramList []param.Param) error {
// 	intParams := paramList[4:6]
// 	return nil
// }

// func CheckStream(paramList []param.Param) error {
// 	if len(paramList) == 7 {
// 		// вызвать один обработчик ()
// 	}
// 	if len(paramList) == 8 {
// 		// вызвать другой обработчик ()
// 	}
// 	return nil
// }

// func Route() {

// }

// func Check(paramList []param.Param) error {
// 	err := CheckBool(paramList)
// 	if err != nil {
// 		return errors.Wrap(err, "Check function error")
// 	}

// 	err := CheckInt(paramList)
// 	if err != nil {
// 		return errors.Wrap(err, "Check function error")
// 	}

// 	err := CheckStream(paramList)
// 	if err != nil {
// 		return errors.Wrap(err, "Check function error")
// 	}

// 	return nil
// }
