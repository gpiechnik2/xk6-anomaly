package anomaly

import (
	"fmt"

	"anomaly/lof"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/anomaly", new(Anomaly))
}

// Httpagg is the k6 extension
type Anomaly struct{}

type LofWithTimestamp struct {
	Value float64
	Timestamp string
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (*Anomaly) Lof(data float64[], distance int) float64[], nil {
	anomalies := float64[]
	samples := lof.GetSamplesFromFloat64s(data)
	lofGetter := lof.NewLOF(len(data))
	lofGetter.Train(samples)

	mapping   := lofGetter.GetLOFs(samples, "strict")
	for sample, factor := range mapping {
		if factor > distance {
			anomalies = append(anomalies, sample)
		}
	}

	return anomalies
}

func (*Anomaly) LofWithTimestamps(data LofWithTimestamp[], distance int) LofWithTimestamp[], nil {
	anomalies := LofWithTimestamp[]
	
	dataWithoutTimestamp := float64[]
	for lofWithTimestamp := range data {
		dataWithoutTimestamp = append(dataWithoutTimestamp, data.Value)
	}

	samples := lof.GetSamplesFromFloat64s(dataWithoutTimestamp)
	lofGetter := lof.NewLOF(len(dataWithoutTimestamp))
	lofGetter.Train(samples)
	mapping := lofGetter.GetLOFs(samples, "strict")

	for sample, factor := range mapping {
		if factor > distance {
			for lofWithTimestamp := range data {
				if sample.GetPoint()[1] == lofWithTimestamp.Value {
					anomalies = append(anomalies, lofWithTimestamp)
				}
			}
		}
	}

	return anomalies
}
