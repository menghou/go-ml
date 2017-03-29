package neighbors

import "go-ml/base"

type DistanceMeasure interface {
	CalDistance(base.DataSet, int, int, []base.FeaturePointer) int
}
