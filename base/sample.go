package base

import (
	"errors"
	"github.com/gonum/matrix/mat64"
	"go-ml/base/linear_algebra/distance"
)

type DataSamples struct {
	Data DataSet
	Fps  []FeaturePointer
}

func (samples *DataSamples) Get(i int) (error, float64) {
	output := make([]float64, len(samples.Fps))
	for j, fp := range samples.Fps {
		err, feature := samples.Data.GetFeatureFromFp(fp)
		if err != nil {
			return err, output
		}
		val := samples.Data.Get(fp, i)
		switch f := feature.(type) {
		case *ContinuousFeature:
			_, output[j] = f.GetFloatFromSys(val)
		default:
			return errors.New("input feature must all be continuous"), output
		}
	}
	return nil, output
}

func (samples *DataSamples) CalDistance(a, b mat64.Dense, distance linear_algebra.DistanceMeasure) (error, float64) {
	return distance(a, b)
}

func NewDataSamples(data DataSet, fps []Feature) *DataSamples {
	return &DataSamples{
		Data: data,
		Fps:  fps,
	}
}
