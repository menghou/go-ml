package base

import (
	"fmt"
	"strconv"
)

//implement Feature => base/types.go
type ContinuousFeature struct {
	Name      string
	Precision int
}

func (feature *ContinuousFeature) setName(name string) {
	feature.Name = name
}
func (feature *ContinuousFeature) GetSysValFromString(str string) (error, []byte) {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return err, nil
	}
	re := PackFloatToBytes(f)
	return nil, re
}
func (feature *ContinuousFeature) GetFloatFromSys(val []byte) (error, float64) {
	return nil, UnPackBytesToFloat(val)
}
func (feature *ContinuousFeature) GetStringFromSysVal(val []byte) (error, string) {
	err, float := feature.GetFloatFromSys(val)
	if err != nil {
		return err, ""
	}
	formatString := fmt.Sprintf("%%.%df", feature.Precision)
	return err, fmt.Sprintf(formatString, float)
}
func (feature *ContinuousFeature) Equal(what Feature) bool {
	f, ok := what.(*ContinuousFeature)
	if !ok {
		return false
	}
	if f.GetName() != feature.GetName() {
		return false
	}
	return true
}
func (feature *ContinuousFeature) String() string {
	return fmt.Sprintf("ContinuousFeature(%s)", feature.Name)
}
func (feature *ContinuousFeature) GetName() string {
	return feature.Name
}
func NewContinuousFeature(name string) *ContinuousFeature {
	return &ContinuousFeature{
		Name: name,
	}
}
