package Features

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"
)

// Load
//Reads the configuration, based on given JSON file and then creates modules.
//Returns created modules, with added WaitGroup to support concurrency.
//
//**
func Load() []func(group *sync.WaitGroup) {
	var model []func(group *sync.WaitGroup)
	//TODO add JSON reader
	jsonFile, errJSON := os.Open("config.json")
	errorHandler(errJSON)

	byteStream, errBYTE := io.ReadAll(jsonFile)
	errorHandler(errBYTE)

	errorHandler(json.Unmarshal(byteStream))
	//model = append(model, configureFeatureOne("Feature 1"))
	//model = append(model, configureFeatureTwo("Feature 2"))

	return model
}
func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
