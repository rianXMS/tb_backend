package Features

import (
	"log"
	"os"
)

func feature2(configurationParams string) func() {

	return func() {
		log.SetOutput(os.Stdout)
		println(configurationParams)
	}
}
