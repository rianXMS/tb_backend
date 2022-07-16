package Features

import (
	"log"
	"sync"
)

func configureFeatureTwo(model modelFeatureTwo) func(group *sync.WaitGroup) {

	return func(group *sync.WaitGroup) {

		log.Println(model.Param)
		group.Done()
	}
}
