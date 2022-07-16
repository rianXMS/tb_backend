package Features

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"
)

const configPath = "config.json"

// Load
//Reads the configuration, based on given JSON file and then creates modules.
//Returns created modules, with added WaitGroup to support concurrency.
//**
func Load() []func(group *sync.WaitGroup) {
	var model []func(group *sync.WaitGroup)
	store := loadFile(configPath)

	//Load Module One
	for i := range store.ModelFeatureOne {
		model = append(model, configureFeatureOne(store.ModelFeatureOne[i]))
	}
	//Load Module Two
	model = append(model, configureFeatureTwo(store.ModelFeatureTwo))

	return model
}

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func loadFile(path string) config {
	var store config
	//Loads JSON file
	jsonFile, errJSON := os.Open(path)
	errorHandler(errJSON)
	//Creates byte-stream
	byteStream, errBYTE := io.ReadAll(jsonFile)
	errorHandler(errBYTE)
	//Interprets byte-stream to JSON
	errorHandler(json.Unmarshal(byteStream, &store))
	return store
}
