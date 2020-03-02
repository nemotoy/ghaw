package main

import (
	"fmt"

	"github.com/pkg/errors"
)

var Err = errors.New("")

func main() {
	fmt.Println(Err)
}
