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

func (*Anomaly) OneClassSvm(trainData []DataPoint, testData []DataPoint) []DataPoint {
	var anomalies []DataPoint
	
	convertedTrainData := make([][]float64, len(trainData))
	for i, dp := range trainData {
		convertedTrainData[i] = []float64{dp.X, dp.Y}
	}

	// Tworzenie i dopasowywanie modelu One-Class SVM
	ocsvm := NewOneClassSVM(rbfKernel)
	ocsvm.Fit(convertedTrainData)

	// Wykrywanie anomalii w danych testowych
	for _, dp := range testData {
		convertedInstance := []float64{dp.X, dp.Y}
		predicted := ocsvm.Predict(convertedInstance)

		if predicted == -1 {
			anomalies = append(anomalies, instance)
		}
	}

	return anomalies
}
