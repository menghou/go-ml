package neighbors

//暂时只支持float型的feature
import (
	"errors"
	"go-ml/base"
	"go-ml/base/linear_algebra/distance"
)

const (
	SonTreeEmpty = errors.New("son tree empty")
)

type KDTree struct {
	X       base.DataSet
	fps     []base.FeaturePointer
	metrics linear_algebra.DistanceMeasure
	root    *KDNode
}

// 如何在kd树中描述“上次未采用的子树” : 不满足父的那一边就是没去过的么。。
type KDNode struct {
	split        int
	sv           float64    //split value
	vr           ValueRange // value range
	leftSubTree  *KDNode
	rightSubTree *KDNode
	parent       *KDNode
	cols         []int // cols this node contains
}

//TODO kd tree build
func (tree *KDTree) Build() error {
	//select most variable feature -> tree.split
	//
}

//TODO recursion build tree
//给他父节点，让他自己生成左右子树
//father的父亲需要在上一层递归里规定好
func (tree *KDTree) build(father *KDNode, cols []int) {
	father.cols = cols
	father.split = tree.selectMostVariableFeature(father)
}

//TODO select most variable feature
//返回的是fps的下标
func (tree *KDTree) selectMostVariableFeature(father *KDNode) (error, int) {
	type MostVariableFeature struct {
		index    int
		variable float64
	}
	m := &MostVariableFeature{index: -1, variable: -1}
	for i := range tree.fps {
		//cal mean for i'st feature
		err, feature := tree.X.GetFeatureFromFp(tree.fps[i])
		if err != nil {
			return err, -1
		}
		con, ok := feature.(*base.ContinuousFeature)
		if !ok {
			return NotContinuousFeatureErr, -1

		}
		data := make([]float64, len(father.cols))
		var sum float64 = 0

		for j, v := range father.cols {
			val := tree.X.Get(tree.fps[i], v)
			_, f := con.GetFloatFromSys(val)
			sum += f
			data[j] = f
		}
	}

}

//TODO kd tree format
func (tree *KDTree) String() string {

}

//TODO kdtree search
func (tree *KDTree) Search(test base.DataSet, k int) base.Feature {

}

//func (tree *KDTree) search(testSamples []interface{}, k int) []interface{} {
//
//}

type ValueRange struct {
	min float64
	max float64
}

func NewKDTree(x base.DataSet, fps []base.FeaturePointer, metrics linear_algebra.DistanceMeasure) (error, *KDTree) {
	init_instance := &KDTree{
		X:       x,
		fps:     fps,
		metrics: metrics,
		root:    *KDNode{},
	}
	if err := CheckParam(x, fps); err != nil {
		return err, nil
	}
	init_instance.Build()
	return nil, init_instance
}
