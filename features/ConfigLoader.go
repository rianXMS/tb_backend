package features

import (
	"encoding/json"
	"io"
	"os"
	"sync"
	"tb_backend/util"
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

func loadFile(path string) config {
	var store config
	//Loads JSON file
	jsonFile, errJSON := os.Open(path)

	util.ErrorHandler(errJSON)
	//Creates byte-stream
	byteStream, errBYTE := io.ReadAll(jsonFile)
	util.ErrorHandler(errBYTE)
	//Interprets byte-stream to JSON
	util.ErrorHandler(json.Unmarshal(byteStream, &store))
	return store
}
