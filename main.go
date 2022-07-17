package main

import (
	"log"
	"sync"
	"tb_backend/features"
)

func main() {
	runModules()
}

func runModules() {
	featureGroup := sync.WaitGroup{}
	a := features.Load()
	for i := range a {
		featureGroup.Add(1)
		log.Printf("Executing %d\n", i+1)
		go a[i](&featureGroup)
		featureGroup.Wait()
	}
}
