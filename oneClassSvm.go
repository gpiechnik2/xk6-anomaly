package anomaly

import (
	"math"
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

// func main() {
// 	// Przykładowe dane treningowe
// 	trainData := [][]float64{
// 		{1.0, 2.0},
// 		{2.0, 3.0},
// 		{3.0, 4.0},
// 		{4.0, 5.0},
// 		{5.0, 6.0},
// 	}

// 	// Tworzenie modelu One-Class SVM z jądrem RBF
// 	ocsvm := NewOneClassSVM(rbfKernel, 0.1)

// 	// Trenowanie modelu
// 	ocsvm.Fit(trainData, 0.01)

// 	// Dane testowe
// 	testData := []DataPoint{
// 		{X: 6.0, Y: 7.0, Timestamp: "2023-07-17T12:00:00"},
// 		{X: 3.0, Y: 4.0, Timestamp: "2023-07-17T12:01:00"},
// 		{X: 1.5, Y: 2.5, Timestamp: "2023-07-17T12:02:00"},
// 	}

// 	// Konwersja danych testowych na format [][]float64
// 	convertedTestData := make([][]float64, len(testData))
// 	for i, point := range testData {
// 		convertedTestData[i] = []float64{point.X, point.Y}
// 	}

// 	// Klasyfikowanie danych testowych
// 	for i, instance := range convertedTestData {
// 		predicted := ocsvm.Predict(instance)
// 		if predicted == -1 {
// 			fmt.Printf("Normal data (%s): %+v\n", testData[i].Timestamp, testData[i])
// 		} else {
// 			fmt.Printf("Anomaly detected (%s): %+v\n", testData[i].Timestamp, testData[i])
// 		}
// 	}
// }

// RBF (Radial Basis Function) kernel
func rbfKernel(x, y []float64) float64 {
	squaredEuclideanDistance := 0.0
	for i := 0; i < len(x); i++ {
		diff := x[i] - y[i]
		squaredEuclideanDistance += diff * diff
	}
	return math.Exp(-squaredEuclideanDistance)
}
