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
	containers := features.GetContainers()
	for i := range containers {
		featureGroup.Add(1)
		log.Printf("Executing %d\n", i+1)
		go containers[i].Runnable(&featureGroup, &containers[i].Terminate)
		featureGroup.Wait()
	}
}
