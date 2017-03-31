package base

type DataSet interface {
	AddFeature(Feature) error
	Set(FeaturePointer, int, []byte) error
	GetAllFeaturePoints() (error, []FeaturePointer)
	GetFeaturePoint(Feature) (error, FeaturePointer)
	GetFeatureFromFp(FeaturePointer) (error, Feature)
	Get(fp FeaturePointer, row int) []byte
	Size() (row, col int)
	FixSize(int) error
	AddClassFeature(Feature) error
	String() string
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
	String() string
	GetStringFromSysVal(val []byte) (error, string)
}
