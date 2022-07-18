package features

import (
	"encoding/json"
	"io"
	"os"
	"tb_backend/util"
)

const configPath = "config.json"

var containers []Container

// Load
//Reads the configuration, based on given JSON file and then creates modules.
//Returns created modules, with added WaitGroup to support concurrency.
//**
func load() {
	store := loadFile(configPath)

	//Load Module One
	for i := range store.ModelFeatureOne {
		containers = append(containers, Container{
			ID:        store.ModelFeatureOne[i].ID,
			Terminate: false,
			Runnable:  configureFeatureOne(store.ModelFeatureOne[i]),
		})
	}
	//Load Module Two
	containers = append(containers, Container{
		ID:        store.ModelFeatureTwo.ID,
		Terminate: false,
		Runnable:  configureFeatureTwo(store.ModelFeatureTwo),
	})
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
func GetContainers() []Container {
	load()
	return containers
}
