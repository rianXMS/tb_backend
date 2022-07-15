package Features

import (
	"log"
)

func feature1(configurationParams string) func() {

	return func() {
		log.Println(configurationParams)
	}
}
