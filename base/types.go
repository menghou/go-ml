package base

type DataSet interface {
	AddFeature(Feature) error
	Set(FeaturePointer, int, []byte) error
	GetAllFeaturePoints() (error, []FeaturePointer)
	GetFeaturePoint(Feature) (error, FeaturePointer)
	Size() (row, col int)
}

//data read interface, hope to support mongo, mysql, rpc
type DataReader interface {
	Read() (error, DataSet)
}
type Feature interface {
	setName(name string)
	GetSysValFromString(string) (error, []byte)
	Equal(Feature) bool
	GetName() string
}
