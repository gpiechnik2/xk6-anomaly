package anomaly

import (
	"math"
	"sort"
)

type LOFResult struct {
	X float64
	Y float64
	Timestamp string
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
		median := data[length/2]
		return median
	} else {
		median := (data[length/2-1] + data[length/2]) / 2
		return median
	}
}

func LocalOutlierFactor(data []DataPoint) []LOFResult {
	var lofResults []LOFResult

	for _, point := range data {
		reachDistances := calculateReachabilityDistances(point, data)
		lof := calculateLOF(reachDistances)
		
		newLofResult := LOFResult{X: point.X, Y: point.Y, LofScore: lof, Timestamp: point.Timestamp}
		lofResults = append(lofResults, newLofResult)
	}
	
	return lofResults
}

func calculateReachabilityDistances(point DataPoint, data []DataPoint) []float64 {
	distances := make([]float64, len(data))

	for i, neighbor := range data {
		distances[i] = EuclideanDistance(point, neighbor)
	}

	return distances
}

func calculateLOF(reachDistances []float64) float64 {
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

func EuclideanDistance(p1, p2 DataPoint) float64 {
	return math.Sqrt(math.Pow(p1.X-p2.X, 2) + math.Pow(p1.Y-p2.Y, 2))
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
	stdDev := math.Sqrt(variance)
	
	return stdDev
}
