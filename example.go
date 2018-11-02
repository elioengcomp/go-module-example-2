package example

import (
	"fmt"

	example_3 "github.com/elioengcomp/go-module-example-3"
)

const version = "v1.0.0"

func Exec() string {

	return fmt.Sprintf("This is go-module-example-2 %s consuming go-module-example-3: %s", version, example_3.Exec())
}
