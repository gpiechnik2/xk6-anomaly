package anomaly

import (
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/anomaly", new(Anomaly))
}

type Anomaly struct{}

func (*Anomaly) Lof(data []DataPoint, threshold float64) []LOFResult {
	var anomalies []LOFResult

	if threshold == nil {
		threshold = 1
	}

	lofResults := LocalOutlierFactor(data)
	stdDev := CalculateStandardDeviation(lofResults)
	medianLOFScore := GetMedianFromLofResults(lofResults)
	threshold =  medianLOFScore - (threshold * stdDev) 

	for _, result := range lofResults {
		if result.LofScore < threshold {
			anomalies = append(anomalies, result)
		}
	}

	return anomalies
}

func (*Anomaly) OneClassSvm(trainData [][]float64, data []DataPoint,  threshold float64) []DataPoint {
	var anomalies []DataPoint
	
	// trainData := [][]float64{
	// 	{1.0, 2.0},
	// 	{2.0, 3.0},
	// 	{3.0, 4.0},
	// 	{4.0, 5.0},
	// 	{5.0, 6.0},
	// }

	ocsvm := NewOneClassSVM(rbfKernel, 0.1)
	ocsvm.Fit(trainData, 0.01)

	// testData := []DataPoint{
	// 	{X: 6.0, Y: 7.0, Timestamp: "2023-07-17T12:00:00"},
	// 	{X: 3.0, Y: 4.0, Timestamp: "2023-07-17T12:01:00"},
	// 	{X: 1.5, Y: 2.5, Timestamp: "2023-07-17T12:02:00"},
	// }

	convertedTestData := make([][]float64, len(data))
	for i, point := range data {
		convertedTestData[i] = []float64{point.X, point.Y}
	}

	for i, instance := range convertedTestData {
		predicted := ocsvm.Predict(instance)
		if predicted != -1 {
			anomalies = append(anomalies, data)
		}
	}

}



