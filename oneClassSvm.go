package anomaly

import (
	"math"
	"sort"
)

type OneClassSVM struct {
	SupportVectors [][]float64
	Bias           float64
	Kernel         func(x, y []float64) float64
	Gamma          float64
}

func NewOneClassSVM(kernel func(x, y []float64) float64, gamma float64) *OneClassSVM {
	return &OneClassSVM{
		SupportVectors: nil,
		Bias:           0,
		Kernel:         kernel,
		Gamma:          gamma,
	}
}

func (svm *OneClassSVM) Fit(X [][]float64, nu float64) {
	numSamples := len(X)

	// Initialization of the gramian variable
	gramMatrix := make([][]float64, numSamples)
	for i := range gramMatrix {
		gramMatrix[i] = make([]float64, numSamples)
	}

	// Calculation of the gramma
	for i := 0; i < numSamples; i++ {
		for j := 0; j < numSamples; j++ {
			gramMatrix[i][j] = svm.Kernel(X[i], X[j])
		}
	}

	// Calculating the bias factor
	sumAlphas := 0.0
	for i := 0; i < numSamples; i++ {
		sumAlphas += gramMatrix[i][i]
	}
	svm.Bias = sumAlphas / float64(numSamples)

	// Writing the support vectors
	svm.SupportVectors = make([][]float64, 0)
	for i := 0; i < numSamples; i++ {
		if gramMatrix[i][i] >= svm.Bias {
			svm.SupportVectors = append(svm.SupportVectors, X[i])
		}
	}

	numSupportVectors := len(svm.SupportVectors)
	numSupportVectorsInBoundary := int(math.Round(nu * float64(numSupportVectors)))
	if numSupportVectorsInBoundary == 0 {
		numSupportVectorsInBoundary = 1
	}

	// Sort the support vectors by distance from the decision boundary.
	distancesFromBoundary := make([]float64, numSupportVectors)
	for i, sv := range svm.SupportVectors {
		dist := svm.Kernel(sv, sv) - 2*svm.Bias
		distancesFromBoundary[i] = dist
	}
	sort.Float64s(distancesFromBoundary)

	// Set the decision boundary to correspond to the distance from the boundary
	svm.Bias = distancesFromBoundary[numSupportVectorsInBoundary-1]
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
	return math.Exp(-squaredEuclideanDistance)
}
