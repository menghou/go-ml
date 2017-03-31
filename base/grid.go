package base

import (
	"bytes"
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
	fgRevMap     map[int]string
	fgs          []*FeatureGroup
	fixed        bool
	maxRow       int
	classFeature map[FeaturePointer]bool
	lock         sync.Mutex
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
		err = errors.New("can't add feature")
		return
	}
	if _, ok = d.fgMap[ag]; !ok {
		err = d.createFeatureGroup(ag, 8)
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
func (d *DataGrid) createFeatureGroup(name string, size int) error {
	fg := &FeatureGroup{
		fs:   make([]Feature, 0),
		size: size,
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
	for i, f := range d.features {
		err, fp := d.GetFeaturePoint(f)
		if err != nil {
			return err, re
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
	return errors.New(fmt.Sprintf("can't resolve %s", what)), FeaturePointer{WhichFeatureGroup: -1, WhichFeatureInGroup: -1, fea: nil}
}
func (d *DataGrid) FixSize(rowCount int) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	for _, fs := range d.fgs {
		rowSize := fs.RowSizeInByte()
		totalSize := rowCount * rowSize
		fs.resize(totalSize)
	}
	d.fixed = true
	d.maxRow += rowCount
	return nil
}
func (d *DataGrid) AddClassFeature(f Feature) error {
	err, fp := d.GetFeaturePoint(f)
	if err != nil {
		return err
	}
	d.lock.Lock()
	defer d.lock.Unlock()

	d.classFeature[fp] = true
	return nil
}
func (d *DataGrid) Size() (int, int) {
	return len(d.features), d.maxRow
}
func (d *DataGrid) Get(fp FeaturePointer, row int) []byte {
	return d.fgs[fp.WhichFeatureGroup].get(fp.WhichFeatureInGroup, row)
}

func (d *DataGrid) GetFeatureFromFp(fp FeaturePointer) (error, Feature) {
	if len(d.fgs) < fp.WhichFeatureGroup {
		return errors.New(fmt.Sprintf("total feature group length %d, get %d", len(d.fgs), fp.WhichFeatureGroup)), nil
	}
	if len(d.fgs[fp.WhichFeatureGroup].fs) < fp.WhichFeatureInGroup {
		return errors.New(fmt.Sprintf("total feature in group length %d, get %d", len(d.fgs[fp.WhichFeatureGroup].fs), fp.WhichFeatureInGroup)), nil
	}
	return nil, d.fgs[fp.WhichFeatureGroup].fs[fp.WhichFeatureInGroup]
}

func (d *DataGrid) String() string {
	var buffer bytes.Buffer

	err, fps := d.GetAllFeaturePoints()
	if err != nil {
		buffer.WriteString(fmt.Sprintf("Get feature error :%v\n", err))
		return buffer.String()
	}
	cols, row := d.Size()
	buffer.WriteString("DataSet: ")
	buffer.WriteString(fmt.Sprintf("%d row ", row))
	buffer.WriteString(fmt.Sprintf("%d features \n", cols))

	buffer.WriteString("Features: \n")
	for _, p := range fps {
		prefix := "\t"
		if d.classFeature[p] {
			prefix = "*\t"
		}
		buffer.WriteString(fmt.Sprintf("%s%s\n", prefix, p.fea))
	}
	maxRow := 20
	if row < maxRow {
		maxRow = row
	}
	for i := 0; i < maxRow; i++ {
		buffer.WriteString("\t")
		for _, a := range fps {
			val := d.Get(a, i)
			err, str := a.fea.GetStringFromSysVal(val)
			if err != nil {
				buffer.WriteString(fmt.Sprintf("%d row err", i))
				return buffer.String()
			}
			buffer.WriteString(fmt.Sprintf("%s ", str))
		}
		buffer.WriteString("\n")
	}
	if row-maxRow == 0 {
		buffer.WriteString("All rows displayed")
	} else {
		buffer.WriteString(fmt.Sprintf("\t...\n%d row(s) undisplayed", row-maxRow))
	}
	return buffer.String()
}
func NewDataGrid() *DataGrid {
	return &DataGrid{
		fgMap:        make(map[string]int),
		fgRevMap:     make(map[int]string),
		features:     make([]Feature, 0),
		fgs:          make([]*FeatureGroup, 0),
		fixed:        false,
		classFeature: make(map[FeaturePointer]bool),
	}
}
