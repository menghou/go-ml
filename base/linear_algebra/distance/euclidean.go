package linear_algebra

import (
	"github.com/gonum/matrix/mat64"
	"math"
)

type EuclideanMetric struct {
}

func (e *EuclideanMetric) InnerProduct(x, y *mat64.Dense) float64 {
	container := mat64.NewDense(0, 0, nil)
	container.MulElem(x, y)

	return mat64.Sum(container)
}

func (e *EuclideanMetric) CalDistance(x, y *mat64.Dense) (error, float64) {
	container := mat64.NewDense(0, 0, nil)
	container.Sub(x, y)

	return nil, math.Sqrt(e.InnerProduct(container, container))
}
