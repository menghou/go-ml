package base

import (
	"errors"
	"fmt"
)

type FeatureGroup struct {
	fs   []Feature
	size int
	vals []byte
}

func (fg *FeatureGroup) AddFeature(f Feature) {
	fg.fs = append(fg.fs, f)
}
func (fg *FeatureGroup) set(col, row int, val []byte) error {
	if len(val) != fg.size {
		return errors.New(fmt.Sprintf("val size can't mach :%d, should be %d", len(val), fg.size))
	}

	offset := fg.offset(col, row)

	copied := copy(fg.vals[offset:], val)
	if copied != fg.size {
		return errors.New(fmt.Sprintf("set only copy %d, should be %d", copied, fg.size))
	}
	return nil
}
func (fg *FeatureGroup) get(col, row int) []byte {
	offset := fg.offset(col, row)
	return fg.vals[offset : offset+fg.size]
}
func (fg *FeatureGroup) offset(col, row int) int {
	return row*len(fg.fs)*fg.size + col*fg.size
}
func (fg *FeatureGroup) AllFeatures() []Feature {
	fs := make([]Feature, len(fg.fs))
	for i, v := range fg.fs {
		fs[i] = v
	}
	return fs
}
func (fg *FeatureGroup) RowSizeInByte() int {
	return len(fg.fs) * fg.size
}

func (fg *FeatureGroup) resize(add int) {
	newVals := make([]byte, len(fg.vals)+add)
	copy(newVals, fg.vals)
	fg.vals = newVals
}
