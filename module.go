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

	if threshold == 0.0 {
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

func (*Anomaly) OneClassSvm(trainData []DataPoint, data []DataPoint) []DataPoint {
	var anomalies []DataPoint
	
	convertedTrainData := make([][]float64, len(trainData))
	for i, point := range trainData {
		convertedTrainData[i] = []float64{point.X, point.Y}
	}

	ocsvm := NewOneClassSVM(rbfKernel, 0.01)
	ocsvm.Fit(convertedTrainData, 0.01)

	convertedTestData := make([][]float64, len(data))
	for i, point := range data {
		convertedTestData[i] = []float64{point.X, point.Y}
	}

	for _, instance := range data {
		convertedInstance := []float64{instance.X, instance.Y}
		predicted := ocsvm.Predict(convertedInstance)
		if predicted != -1 {
			anomalies = append(anomalies, instance)
		}
	}

	return anomalies
}
