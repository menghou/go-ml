package neighbors

type DistanceMeasure interface {
	CalDistance(a, b interface{}) (error, float64)
}
