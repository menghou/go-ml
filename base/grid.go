package base

import (
	"errors"
	"fmt"
	"sync"
)

// grid implement DataSet in types.go
type DataGrid struct {
	features []Feature
	//fgMap means index of the feature group in fgs
	fgMap map[string]int
	//fgRevMap is the rev of fgMap
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
func (d *DataGrid) createFeatureGroup(name string) error {
	fg := &FeatureGroup{
		fs: make([]Feature, 0),
	}
	if d.fixed {
		return errors.New("data grid is fixed")
	}
	d.fgMap[name] = len(d.fgMap)
	d.fgRevMap[len(d.fgMap)] = name
	d.fgs = append(d.fgs, fg)
	return nil
}
func (d *DataGrid) Set(fp FeaturePointer, row int, val []byte) error {
	return d.fgs[fp.WhichFeatureGroup].set(fp.WhichFeatureInGroup, row, val)
}
func (d *DataGrid) GetAllFeaturePoints() (error, []FeaturePointer) {
	re := make([]FeaturePointer, len(d.features))
	for i, f := range re {
		err, fp := d.GetFeaturePoint(f)
		if err != nil {
			return err
		}
		re[i] = fp
	}
	return nil, re
}
func (d *DataGrid) GetFeaturePoint(what Feature) (error, FeaturePointer) {
	for i, fg := range d.fgs {
		for j, f := range fg.AllFeatures() {
			if f.Equal(what) {
				return nil, FeaturePointer{WhichFeatureGroup: i, WhichFeatureInGroup: j, fea: f}
			}
		}
	}
	return FeaturePointer{WhichFeatureGroup: -1, WhichFeatureInGroup: -1, fea: nil}, errors.New(fmt.Sprintf("can't resolve %s", what))
}
func NewDataGrid() *DataGrid {
	return &DataGrid{
		fgMap:    make(map[string]int),
		fgRevMap: make(map[int]string),
		features: make([]Feature, 0),
		fgs:      make(map[string]*FeatureGroup),
		fixed:    false,
	}
}
