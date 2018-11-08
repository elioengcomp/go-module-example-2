package anotherpackage

import (
	"fmt"
)

const version = "v1.0.3"

func Exec() string {

	return fmt.Sprintf("This is go-module-example-2 %s", version)
}
