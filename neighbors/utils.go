package neighbors

import (
	"errors"
	"go-ml/base"
)

const (
	NotContinuousFeatureErr = errors.New("please make sure all input feature pointer is continuous")
)

func CheckParam(x base.DataSet, fps []base.FeaturePointer) error {
	for _, fp := range fps {
		err, feature := x.GetFeatureFromFp(fp)
		if err != nil {
			return err
		} else if _, ok := feature.(base.ContinuousFeature); !ok {
			return NotContinuousFeatureErr
		} else {
			continue
		}
	}
	return nil
}
