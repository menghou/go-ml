package base

import "sync"
// grid implement DataSet in types.go
type DataGrid struct {
	featureList []Feature
	featureGroup map[string]*FeatureGroup
	fixed bool
	lock sync.Mutex
}
