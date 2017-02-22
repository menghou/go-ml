package base
//implement Feature => base/types.go
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
