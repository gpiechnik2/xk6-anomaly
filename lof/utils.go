package lof

import "math"

func SampleDist(sample1, sample2 ISample) float64 {
    features1 := sample1.GetPoint()
    features2 := sample2.GetPoint()
    var d float64
    for i := 0; i < len(features1); i++ {
        d += math.Pow(features1[i] - features2[i], 2)
    }
    return math.Sqrt(d)
}
