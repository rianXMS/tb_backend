package Features

type config struct {
	ModelFeatureOne []modelFeatureOne `json:"featureOne"`
	ModelFeatureTwo modelFeatureTwo   `json:"featureTwo"`
}
type modelFeatureOne struct {
	State    bool   `json:"active"`
	Interval int    `json:"interval"`
	Param    string `json:"param"`
}
type modelFeatureTwo struct {
	State bool   `json:"active"`
	Param string `json:"param"`
}
