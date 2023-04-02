package files

import (
	"errors"
	"fmt"
)

func Upload(filename string, size int, options ...string) (myfile string, err error) {
	fmt.Println(filename, size)
	fmt.Println(options)

	myfile = "new-todo.txt"

	if filename != "todo.txt" {
		err = errors.New("Invalid filename")
	}

	return
}
