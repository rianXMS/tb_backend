package Features

import (
	"log"
	"sync"
)

func feature1(configurationParams string) func(group *sync.WaitGroup) {

	return func(group *sync.WaitGroup) {

		log.Println(configurationParams)
		group.Done()
	}
}
