package base

type DataSet interface {
	AddFeature(Feature) error
}
//data read interface, hope to support mongo, mysql, rpc
type DataReader interface {
	Read() (err error, dataSet DataSet)
}
type Feature interface {
	setName(name string)
}
