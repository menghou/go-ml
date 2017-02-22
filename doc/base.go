package doc
/*
base库设计：
需要定义几个接口：
DataSet:
	数据集接口，用DataSetReader生成。
	应该实现的功能：
		1.对于连续型数据，可以做归一化、标准化、白化。
		2.提供sample个数和feature个数.
		3.提供所有的features.
DataSetReader:
	数据读取接口，用来生成DataSetReader。目前希望以mongo, mysql, csv等格式的实现。
Features:
	用来定义特征的接口。目前包括连续型和离散型特征两种
需要实现的结构体：
DataGrid:
	实现DataSet。数据集的主要形式。其中包括以下功能：
		1.特征类型依据连续型、离散型进行分类保存。
		2.对每一种数据类型，可以进行直接定位检索。
FeatureGroup
需要实现的utils：
func DataSetTestTrainSplit(*DataSet,float32) (*DataSet, *DataSet)
 */
