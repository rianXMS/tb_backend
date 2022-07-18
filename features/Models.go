package features

import (
	"github.com/google/uuid"
	"sync"
)

type config struct {
	ModelFeatureOne []modelFeatureOne `json:"featureOne"`
	ModelFeatureTwo modelFeatureTwo   `json:"featureTwo"`
}
type modelFeatureOne struct {
	State    bool      `json:"active"`
	ID       uuid.UUID `json:"ID"`
	Interval int       `json:"interval"`
	Param    string    `json:"param"`
}
type modelFeatureTwo struct {
	State bool      `json:"active"`
	ID    uuid.UUID `json:"ID"`
	Param string    `json:"param"`
}
type Container struct {
	ID        uuid.UUID
	Terminate bool
	Runnable  func(group *sync.WaitGroup, terminate *bool)
}
