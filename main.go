package main

import (
	"log"
	"sync"
	"tb_backend/Features"
)

func main() {
	runModules()
}

func runModules() {
	featureGroup := sync.WaitGroup{}
	a := Features.Load()
	for i := range a {
		featureGroup.Add(1)
		log.Printf("Executing %d\n", i+1)
		go a[i](&featureGroup)
		featureGroup.Wait()
	}
}
