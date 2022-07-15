package Features

func Load() []func() {
	var model []func()
	//TODO add JSON reader
	model = append(model, feature1("Feature 1"))
	model = append(model, feature2("Feature 2"))

	return model
}
