package anomaly

import (
	// "fmt"
	"math"
	"sort"
)

type DataPoint struct {
	X, Y float64
}

type LOFResult struct {
	DataPoint
	LofScore float64
	Timestamp string
}

type Point struct {
	X float64
	Y float64
	Timestamp string
}

// func main() {
// 	// Przykładowe dane
// 	data := []DataPoint{
// 		{0, 0},
// 		{1, 0},
// 		{2, 0},
// 		{3, 0},
// 		{2, 0},
// 		{25, 0},
//         {120, 0},
// 		{0, 0},
// 	}

// 	lofResults := LocalOutlierFactor(data)

// 	stdDev := CalculateStandardDeviation(lofResults)
// 	medianLofScore := GetMedianFromLofResults(lofResults)

// 	// mean or medianLofScore? 
// 	threshold :=  medianLofScore - (1 * stdDev) // Próg jako 2 odchylenia standardowe powyżej średniej

// 	fmt.Println("LOF wyniki:")
// 	for _, result := range lofResults {
// 		fmt.Printf("Punkt (%.5f, %.5f), LOF: %.5f\n", result.X, result.Y, result.LofScore)
// 	}

// 	fmt.Printf("\nAnaliza statystyczna:\nOdchylenie standardowe: %.5f\n", stdDev)
// 	fmt.Printf("Próg wartości odstępstwa: %.5f\n", threshold)

// 	// Sprawdzanie odstępstw na podstawie progu
// 	fmt.Println("\nOdstępstwa:")
// 	for _, result := range lofResults {
// 		if result.LofScore < threshold {
// 			fmt.Printf("Punkt (%.5f, %.5f), LOF: %.5f\n", result.X, result.Y, result.LofScore)
// 		}
// 	}
// }

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
	lofResults := make([]LOFResult, len(data))

	for i, point := range data {
		reachDistances := calculateReachabilityDistances(point, data)
		lof := calculateLOF(reachDistances)

		lofResults[i] = LOFResult{
			DataPoint: point,
			LofScore:  lof,
			Timestamp: point.Timestamp
		}
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


	// 17;51;10
	// 18;10;01
	// 02;10;27
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

// func LOFDetection(data []DataPoint, threshold float64) {

// 	// const data = [
// 		// { x: float64, y: float64, timestamp: string },
// 		// { x: float64, y: float64, timestamp: string }
// 	// ]

// 	lofResults := LocalOutlierFactor(data)
// 	stdDev := CalculateStandardDeviation(lofResults)
// 	medianLofScore := GetMedianFromLofResults(lofResults)
// 	threshold :=  medianLofScore - (1 * stdDev) // Próg jako 2 odchylenia standardowe powyżej średniej

// 	// Sprawdzanie odstępstw na podstawie progu
// 	fmt.Println("\nOdstępstwa:")

// 	for _, result := range lofResults {
// 		if result.LofScore < threshold {
// 			fmt.Printf("Punkt (%.5f, %.5f), LOF: %.5f\n", result.X, result.Y, result.LofScore)
// 		}
// 	}

// }

