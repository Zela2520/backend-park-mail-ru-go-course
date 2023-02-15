package iostream

import (
	"fmt"
)

// TODO: в аргументах должен быть интерфейс на поток os.Stdin || os.Stdout
func Read() {
	fmt.Println("Read from stream")
}

func Write() {
	fmt.Println("WriteFromStream")
}

func GreeetIostream() {
	fmt.Println("iostream package has been imported")
}
