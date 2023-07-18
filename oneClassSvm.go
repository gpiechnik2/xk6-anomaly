package anomaly

import (
	"math"
)

type OneClassSVM struct {
	SupportVectors [][]float64
	Bias           float64
	Kernel         func(x, y []float64) float64
}

func NewOneClassSVM(kernel func(x, y []float64) float64) *OneClassSVM {
	return &OneClassSVM{
		SupportVectors: nil,
		Bias:           0,
		Kernel:         kernel,
	}
}

func (svm *OneClassSVM) Fit(X [][]float64) {
	numSamples := len(X)

	// initialization of the gramian matrix
	gramMatrix := make([][]float64, numSamples)
	for i := range gramMatrix {
		gramMatrix[i] = make([]float64, numSamples)
	}

	// calculation of the gramian value
	for i := 0; i < numSamples; i++ {
		for j := 0; j < numSamples; j++ {
			gramMatrix[i][j] = svm.Kernel(X[i], X[j])
		}
	}

	// calculating the bias value
	sumAlphas := 0.0
	for i := 0; i < numSamples; i++ {
		sumAlphas += gramMatrix[i][i]
	}
	svm.Bias = sumAlphas / float64(numSamples)

	// storing support vectors
	svm.SupportVectors = make([][]float64, 0)
	for i := 0; i < numSamples; i++ {
		if gramMatrix[i][i] >= svm.Bias {
			svm.SupportVectors = append(svm.SupportVectors, X[i])
		}
	}
}

func (svm *OneClassSVM) Predict(X []float64) int {
	result := -svm.Bias

	for _, sv := range svm.SupportVectors {
		result += svm.Kernel(sv, X)
	}

	if result >= 0 {
		return 1
	}
	return -1
}

// RBF (Radial Basis Function) kernel
func rbfKernel(x, y []float64) float64 {
	squaredEuclideanDistance := 0.0
	for i := 0; i < len(x); i++ {
		diff := x[i] - y[i]
		squaredEuclideanDistance += diff * diff
	}
	gamma := 1.0 / float64(len(x))
	return math.Exp(-gamma * squaredEuclideanDistance)
}
