package base

import (
	"errors"
	"fmt"
)

//implement Feature => base/types.go
type DiscreteFeature struct {
	Name string
	vals []string
}

func (feature *DiscreteFeature) setName(name string) {
	feature.Name = name
}
func (feature *DiscreteFeature) GetSysValFromString(str string) (error, []byte) {
	existIndex := -1
	for i, s := range feature.vals {
		if s == str {
			existIndex = i
			break
		}
	}
	if existIndex == -1 {
		feature.vals = append(feature.vals, str)
		existIndex = len(feature.vals)
	}

	return nil, PackU64ToBytes(uint64(existIndex))
}
func (feature *DiscreteFeature) GetStringFromSysVal(val []byte) (error, string) {
	existIndex := int(UnPackBytesToFloat(val))
	if existIndex >= len(feature.vals) {
		return errors.New(fmt.Sprintf("Out of range: %d in %d (%s)", existIndex, len(feature.vals), feature)), ""
	}
	return nil, feature.vals[existIndex]
}

func (feature *DiscreteFeature) Equal(what Feature) bool {
	f, ok := what.(*DiscreteFeature)
	if !ok {
		return false
	}
	if feature.GetName() != f.GetName() {
		return false
	}

	if len(feature.vals) != len(f.vals) {
		return false
	}

	for i, v := range f.vals {
		if v != feature.vals[i] {
			return false
		}
	}
	return true
}
func (feature *DiscreteFeature) GetName() string {
	return feature.Name
}
func (feature *DiscreteFeature) String() string {
	return fmt.Sprintf("DiscreteFeature(%s|%s)", feature.Name, feature.vals)
}
func NewDiscreteFeature(name string) *DiscreteFeature {
	return &DiscreteFeature{
		Name: name,
	}
}
