package base

type DataSamples struct {
	Data DataSet
	Fps  []FeaturePointer
}

func (samples *DataSamples) Get(i int) (error, []interface{}) {
	output := make([]interface{}, len(samples.Fps))
	for j, fp := range samples.Fps {
		err, feature := samples.Data.GetFeatureFromFp(fp)
		if err != nil {
			return err, output
		}
		val := samples.Data.Get(fp, i)
		switch f := feature.(type) {
		case *ContinuousFeature:
			_, output[j] = f.GetFloatFromSys(val)
		case *DiscreteFeature:
			err, output[j] = f.GetStringFromSysVal(val)
			if err != nil {
				return err, output
			}
		}
	}
	return nil, output
}

func (samples *DataSamples) CalDistance(a, b interface{}, distance func(a, b interface{}) (error, float64)) (error, float64) {
	return distance(a, b)
}

func NewDataSamples(data DataSet, fps []Feature) *DataSamples {
	return &DataSamples{
		Data: data,
		Fps:  fps,
	}
}
