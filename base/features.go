package base

type FeatureGroup struct {

}
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

type DiscreteFeature struct {
	Name string
}

func (feature *DiscreteFeature)setName(name string) {
	feature.Name = name
}

func NewDiscreteFeature (name string) *DiscreteFeature {
	return &DiscreteFeature{
		Name:name,
	}
}
