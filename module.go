package anomaly

import (
	"go.k6.io/k6/js/modules"
	"fmt"
)

func init() {
	modules.Register("k6/x/anomaly", new(Anomaly))
}

// Httpagg is the k6 extension
type Anomaly struct{}

func (*Anomaly) Lof(data []DataPoint) {
	var anomalies []LOFResult

	lofResults := LocalOutlierFactor(data)
	stdDev := CalculateStandardDeviation(lofResults)
	medianLOFScore := GetMedianFromLofResults(lofResults)
	threshold :=  medianLOFScore - (1 * stdDev) // Próg jako 2 odchylenia standardowe powyżej średniej

	fmt.Println("\nOdstępstwa:")
	for _, result := range lofResults {
		if result.LOFScore < threshold {
			anomalies = append(anomalies, result)
			// fmt.Printf("Punkt (%.5f, %.5f), LOF: %.5f\n", result.X, result.Y, result.LOFScore)
		}
	}

	return anomalies
}



