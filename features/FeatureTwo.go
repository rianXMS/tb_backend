package features

import (
	"log"
	"sync"
)

func configureFeatureTwo(model modelFeatureTwo) func(group *sync.WaitGroup, terminate *bool) {

	return func(group *sync.WaitGroup, terminate *bool) {

		log.Println(model.Param)
		group.Done()
	}
}
