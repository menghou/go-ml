package base

type FeatureGroup struct {
	fs []Feature
}
func (fg *FeatureGroup) AddFeature(f Feature) {
	fg.fs = append(fg.fs, f)
}

