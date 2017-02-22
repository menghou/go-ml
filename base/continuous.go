package base
//implement Feature => base/types.go
type ContinuousFeature struct {
	Name string
	Precision int
}

func (feature *ContinuousFeature) setName(name string) {
	feature.Name = name
}

func NewContinuousFeature(name string) *ContinuousFeature {
	return &ContinuousFeature{
		Name:name,
	}
}

