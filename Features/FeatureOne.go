package Features

import (
	"log"
	"sync"
)

func configureFeatureOne(model modelFeatureOne) func(group *sync.WaitGroup) {

	return func(group *sync.WaitGroup) {

		log.Println(model.Param)
		group.Done()
	}
}
