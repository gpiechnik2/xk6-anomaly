package anomaly

import (
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/anomaly", new(Anomaly))
}

type Anomaly struct{}

func (*Anomaly) Lof(data []float64, threshold float64) []LOFResult {
	var anomalies []LOFResult

	if threshold == 0.0 {
		threshold = 1
	}

	lofResults := LocalOutlierFactor(data)
	median := GetMedianFromLofResults(lofResults)
	stdDev := CalculateStandardDeviation(lofResults)
	threshold = median - (threshold * stdDev)

	for _, result := range lofResults {
		if result.LofScore < threshold {
			anomalies = append(anomalies, result)
		}
	}

	return anomalies
}

func (*Anomaly) OneClassSvm(trainData []float64, testData []float64, gammaValue float64) []float64 {
	if gammaValue == 0.0 {
		gammaValue = 50.0
	}

	svm := NewOneClassSVM(gammaValue, rbfKernel)
	svm.Fit(trainData)
	anomalies := svm.Predict(testData)

	return anomalies
}
