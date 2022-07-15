package Features

func feature2(configurationParams string) func() {

	return func() {
		println(configurationParams)
	}
}
