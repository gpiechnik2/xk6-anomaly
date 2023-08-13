package anomaly

import (
	"math"
	"sort"
)

type LOFResult struct {
	Value    float64
	LofScore float64
}

func GetMedianFromLofResults(lofResults []LOFResult) float64 {
	var data []float64

	for _, result := range lofResults {
		data = append(data, result.LofScore)
	}

	sort.Float64s(data)

	length := len(data)
	if length%2 == 1 {
		return data[length/2]
	}
	return (data[length/2-1] + data[length/2]) / 2
}

func LocalOutlierFactor(data []float64) []LOFResult {
	var lofResults []LOFResult

	for _, point := range data {
		reachDistances := CalculateReachabilityDistances(point, data)
		lof := CalculateLOF(reachDistances)

		newLofResult := LOFResult{Value: point, LofScore: lof}
		lofResults = append(lofResults, newLofResult)
	}

	return lofResults
}

func CalculateReachabilityDistances(point float64, data []float64) []float64 {
	distances := make([]float64, len(data))

	for i, neighbor := range data {
		distances[i] = EuclideanDistance(point, neighbor)
	}

	return distances
}

func CalculateLOF(reachDistances []float64) float64 {
	lrd := 0.0

	for _, reachDistance := range reachDistances {
		lrd += reachDistance
	}

	lrd /= float64(len(reachDistances))

	lof := 0.0

	for _, reachDistance := range reachDistances {
		lof += reachDistance / lrd
	}

	lof /= float64(len(reachDistances))
	lof /= lrd

	return lof
}

func EuclideanDistance(p1, p2 float64) float64 {
	return math.Abs(p1 - p2)
}

func CalculateStandardDeviation(lofResults []LOFResult) float64 {
	scores := make([]float64, len(lofResults))
	sum := 0.0

	// mean
	for i, result := range lofResults {
		scores[i] = result.LofScore
		sum += result.LofScore
	}
	mean := sum / float64(len(lofResults))

	// standard deviation
	varianceSum := 0.0
	for _, score := range scores {
		varianceSum += math.Pow(score-mean, 2)
	}

	variance := varianceSum / float64(len(lofResults))
	return math.Sqrt(variance)
}
