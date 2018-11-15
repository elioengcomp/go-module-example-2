package example

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	example_3 "github.com/elioengcomp/go-module-example-3"
)

const version = "v1.0.3"

func Exec() string {

	log.SetLevel(log.WarnLevel)
	return fmt.Sprintf("This is go-module-example-2 %s consuming go-module-example-3: %s", version, example_3.Exec())
}
