package neighbors

import "go-ml/base"

type KDTree struct {
	X         base.DataSet
	fps       []base.FeaturePointer
	leaf_size int
	metrics   DistanceMeasure
	root      *KDNode
}

//TODO 如何在kd树中描述“上次未采用的子树”
type KDNode struct {
	split        int
	sv           float64    //split value
	vr           ValueRange // value range
	leftSubTree  *KDNode
	rightSubTree *KDNode
	parent       *KDNode
	//lastVisitor
}

type ValueRange struct {
	min float64
	max float64
}
