package Features

import (
	"fmt"
)

func feature1(configurationParams string) func() {

	return func() {

		fmt.Println(configurationParams)
	}
}
