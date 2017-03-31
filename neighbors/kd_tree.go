package neighbors

import "go-ml/base"

type KDTree struct {
	X       base.DataSet
	fps     []base.FeaturePointer
	metrics DistanceMeasure
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
	cols         []int
}

//TODO kd tree build
func (tree *KDTree) Build() error {
	//select most variable feature -> tree.split
	//
	for i := range tree.fps {

	}
}

//TODO recursion build tree
func (tree *KDTree) build()

//TODO select most variable feature
func (tree *KDTree) selectMostVariableFeature(fps []base.FeaturePointer) int {

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

func NewKDTree(x base.DataSet, fps []base.FeaturePointer, metrics DistanceMeasure) *KDTree {
	init_instance := &KDTree{
		X:       x,
		fps:     fps,
		metrics: metrics,
		root:    *KDNode{},
	}
	init_instance.Build()
	return init_instance
}
