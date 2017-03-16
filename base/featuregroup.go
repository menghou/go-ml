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
func (fg *FeatureGroup) offset(col, row int) int {
	return col*len(fg.fs)*fg.size + row*fg.size
}
func (fg *FeatureGroup) AllFeatures() []Feature {
	fs := make([]Feature, len(fg.fs))
	for i, v := range fg.fs {
		fs[i] = v
	}
	return fs
}
