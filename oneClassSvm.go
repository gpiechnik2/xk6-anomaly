package anomaly

import (
	"math"
)

type OneClassSVM struct {
	SupportVectors []float64
	Bias           float64
	Gamma          float64
	Kernel         func(x, y, gamma float64) float64
}

func NewOneClassSVM(gamma float64, kernel func(x, y, gamma float64) float64) *OneClassSVM {
	return &OneClassSVM{
		SupportVectors: nil,
		Bias:           0,
		Gamma:          gamma,
		Kernel:         kernel,
	}
}

func (svm *OneClassSVM) Fit(normalData []float64) {
	numSamples := len(normalData)

	// Initialization of the gramian matrix
	gramMatrix := make([][]float64, numSamples)
	for i := range gramMatrix {
		gramMatrix[i] = make([]float64, numSamples)
	}

	// Calculation of the gramian value
	for i := 0; i < numSamples; i++ {
		for j := 0; j < numSamples; j++ {
			gramMatrix[i][j] = svm.Kernel(normalData[i], normalData[j], svm.Gamma)
		}
	}

	// Calculating the bias value
	sumAlphas := 0.0
	for i := 0; i < numSamples; i++ {
		sumAlphas += gramMatrix[i][i]
	}
	svm.Bias = sumAlphas / float64(numSamples)

	// Storing support vectors
	svm.SupportVectors = make([]float64, 0)
	for i := 0; i < numSamples; i++ {
		if gramMatrix[i][i] >= svm.Bias {
			svm.SupportVectors = append(svm.SupportVectors, normalData[i])
		}
	}
}

func (svm *OneClassSVM) Predict(potentialAnomalies []float64) []float64 {
	var anomalies []float64

	for _, dataPoint := range potentialAnomalies {
		result := -svm.Bias

		for _, sv := range svm.SupportVectors {
			result += svm.Kernel(sv, dataPoint, svm.Gamma)
		}

		if result < 0 {
			anomalies = append(anomalies, dataPoint)
		}
	}

	return anomalies
}

// RBF (Radial Basis Function) kernel
func rbfKernel(x, y, gamma float64) float64 {
	diff := x - y
	squaredEuclideanDistance := diff * diff
	return math.Exp(-gamma * squaredEuclideanDistance)
}
