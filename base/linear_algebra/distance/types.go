package linear_algebra

import "github.com/gonum/matrix/mat64"

type DistanceMeasure interface {
	CalDistance(a, b *mat64.Dense) (error, float64)
}
