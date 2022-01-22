package fever

const threshold float64 = 99.5

func Determine(temp float64) bool {
	return temp > threshold
}
