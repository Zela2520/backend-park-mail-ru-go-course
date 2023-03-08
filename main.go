package main

import (
	"errors"
	"log"

	param "github.com/Zela2520/backend-park-mail-ru-go-course.git/uniq/param"
	route "github.com/Zela2520/backend-park-mail-ru-go-course.git/uniq/route"
)

func main() {
	var (
		err error
	)

	options, err := param.GetParams()
	if err != nil {
		log.Fatal(errors.Unwrap(err))
	}

	err = route.Route(options)
	if err != nil {
		log.Fatal(errors.Unwrap(err))
	}
}
