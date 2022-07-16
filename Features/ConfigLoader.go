package Features

import "sync"

// Load
//Reads the configuration, based on given JSON file and then creates modules.
//Returns created modules, with added WaitGroup to support concurrency.
//
//**
func Load() []func(group *sync.WaitGroup) {
	var model []func(group *sync.WaitGroup)
	//TODO add JSON reader
	model = append(model, feature1("Feature 1"))
	model = append(model, feature2("Feature 2"))

	return model
}
