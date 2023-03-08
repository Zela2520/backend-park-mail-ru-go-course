package calc

import (
	"fmt"

	src "github.com/Zela2520/backend-park-mail-ru-go-course.git/calc/src"
)

func Calc() {
	input := src.GetInput()

	res, err := src.Calc(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%.3f\n", res)
}
