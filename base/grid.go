package base

import (
	"sync"
	"errors"
)
// grid implement DataSet in types.go
type DataGrid struct {
	features []Feature
	fgMap    map[string]int
	fgRevMap map[int]string
	fgs      []*FeatureGroup
	fixed    bool
	lock     sync.Mutex
}

func (d *DataGrid) AddFeature(f Feature) (err error) {
	d.lock.Lock()
	defer d.lock.Unlock()
	ag := "continuous"
	var ok bool
	if _, ok = f.(*DiscreteFeature); ok {
		ag = "discrete"
	} else if _, ok = f.(*ContinuousFeature); ok {
		ag = "continuous"
	} else {
		err = "can't add feature"
		return
	}
	if _, ok = d.fgMap[ag]; !ok {
		err = d.createFeatureGroup(ag)
		if err != nil {
			return
		}
	}
	id := d.fgMap[ag]
	a := d.fgs[id]
	a.AddFeature(f)
	d.features = append(d.features, f)
	return nil
}
func (d *DataGrid) createFeatureGroup(name string) (error) {
	fg := &FeatureGroup{
		fs:make([]Feature,0),
	}
	if d.fixed {
		return errors.New("data grid is fixed")
	}
	d.fgMap[name] = len(d.fgMap)
	d.fgRevMap[len(d.fgMap)] = name
	d.fgs = append(d.fgs, fg)
	return nil
}

func NewDataGrid() *DataGrid {
	return &DataGrid{
		fgMap:make(map[string]int),
		fgRevMap:make(map[int]string),
		features:make([]Feature,0),
		fgs:make(map[string]*FeatureGroup),
		fixed:false,
	}
}
