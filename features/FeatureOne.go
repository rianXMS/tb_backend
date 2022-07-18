package features

import (
	"log"
	"sync"
)

func configureFeatureOne(model modelFeatureOne) func(group *sync.WaitGroup, terminate *bool) {

	return func(group *sync.WaitGroup, terminate *bool) {

		log.Println(model.Param)
		group.Done()
	}
}
