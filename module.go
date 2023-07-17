package anomaly

import (
	"fmt"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/anomaly", new(Anomaly))
}

type Anomaly struct{}

func (*Anomaly) Lof(data []DataPoint, threshold float64) []LOFResult {
	var anomalies []LOFResult

	lofResults := LocalOutlierFactor(data)
	stdDev := CalculateStandardDeviation(lofResults)
	medianLOFScore := GetMedianFromLofResults(lofResults)
	threshold =  medianLOFScore - (threshold * stdDev) // Próg jako 2 odchylenia standardowe powyżej średniej

	fmt.Println("Test")
	for _, result := range lofResults {
		if result.LofScore < threshold {
			fmt.Println(result)
			anomalies = append(anomalies, result)
		}
	}

	fmt.Println("anomalies")
	fmt.Println(anomalies)
	return anomalies
}

// TODO
// func (*Anomaly) OneClassSvm(trainingData []DataPoint, data []DataPoint,  threshold) []OneClassSvmResult {
// var anomalies []OneClassSvmResult
// }



