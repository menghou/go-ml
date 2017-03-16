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
	数据读取接口，用来生成DataSetReader。目前希望以mongo, mysql, csv, rpc等格式的实现。
Feature:
	用来定义特征的接口。目前包括连续型和离散型特征两种
	SetName(name string)

需要实现的结构体：
DataGrid:
	实现DataSet。数据集的主要形式。其中包括以下功能：
		1.特征类型依据连续型、离散型进行分类保存。
		2.对每一种数据类型，可以进行直接定位检索。
	关于实现上面的第二个功能，需要实现一个函数，用来返回一个[]FeaturePoint
FeaturePoint:
        实现特征的快速检索，应该三个属性：
            1. 该特征应该归属于第几个FeatureGroup
            2. 该特征在这个FeatureGroup中归属于第几个特征
            3. 该特征本身（是个接口）


FeatureGroup
	实现特征的分类，应该有以下功能：
		1.返回自己包括的所有的特征
		2.set功能，需要输入第几个特征和val
DiscreteFeature 离散型特征
	实现Feature,  GetSysValFromString需要首先判断输入的string是否已经在自己的vals里了，如果不存在，新添加，返回其索引；如果存在，返回其索引，然后将索引进行pack转换，返回
ContinuousFeature 连续型特征
	实现Feature,  GetSysValFromString直接将float进行pack，返回
需要实现的utils：
func DataSetTestTrainSplit(*DataSet,float32) (*DataSet, *DataSet)
*/
