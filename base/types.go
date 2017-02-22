package base

type DataSet interface {
	AddFeature(Feature) error
}
type DataReader interface {
}
type Feature interface {
	setName(name string)
}
