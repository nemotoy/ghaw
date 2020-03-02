package main

import (
	"fmt"

	"github.com/pkg/errors"
)

var e = errors.New("")

func main() {
	fmt.Println(e)
}
