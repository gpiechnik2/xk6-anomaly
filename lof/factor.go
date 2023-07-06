package lof

import (
    "log"
    "math"
)

const (
    // It might be the case that many samples in the training set have exactly
    // the same features. This results in zero distances, which may produce
    // Inf values for float division in getDensity(); we set a minimal sum value
    // to avoid that.
    CMinimalSum = 10e-15
)

const (
    LNotEnoughSamples = "LOF: not enough samples to train!"
)

//////////////////////////////////////////////////////////////////////////////
//
//  A Local Outlier Factor algorithm (`Markus M. Breunig) implementation.
//  This implementation allows you to choose between updating nearest
//  neighbors for all the samples OR updating nearest neighbors only for
//  nearest neighbors of the added sample; See @mode parameter in GetLOFs()
//  and GetLOF(). 
//
//////////////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////////////
//
//  This type is used to keep information anout distance between
//  (sample1, sample2) and sample2's index
// 
//////////////////////////////////////////////////////////////////////////////

type DistItem struct {
    Value       float64
    Index       int
}

//////////////////////////////////////////////////////////////////////////////
//
//  Main type; populate instances with lof.NewLof().
// 
//////////////////////////////////////////////////////////////////////////////

type LOF struct {
    // Slice of training samples 
    TrainingSet []ISample
    // Distances between the samples
    Distances   [][]DistItem
    // [sample index][slice of sample indices], used to keep
    // nearest neighbors info for each sample (first neighbor
    // is the nearest)
    KNNs        [][]int
    // Keeps the LOF.KNNs state just after training; see Reset()
    KNNsBackup  [][]int
    // Not used yet
    MinTrain    []float64
    // Not used yet
    MaxTrain    []float64
    // MinPts parameter from the paper, same as kNN in other
    // implementations
    MinPts      int
    NumSamples  int
    AddedIndex  int
}

// Constructor for LOF type. Check out ./samples.GetSamplesFromFloat64s()
// for fast [][]float64 -> []ISample conversion.
func NewLOF(minPts int) *LOF {
    // Create the LOF object
    lof := &LOF{
        MinPts: minPts,
    }
    return lof
}

// Pre-compute distances between training samples and store their
// nearest neighbors in LOF.KNNs.
func (lof *LOF) Train(samples []ISample) {
    lof.checkSamples(samples)
    numSamples := len(samples)
    // After training we want to compute LOF values for
    // new samples, and we need some space for their
    // distances; if we find LOF for one new sample at a
    // time, a single additional slot will be enough.
    addedIndex := len(samples) + 1
    lof.TrainingSet = samples
    lof.NumSamples = numSamples
    lof.AddedIndex = addedIndex        
    // Prepare storage between training samples
    lof.Distances = make([][]DistItem, addedIndex)
    for idx := 0; idx < addedIndex; idx++ {
        lof.Distances[idx] = make([]DistItem, addedIndex)
    }
    // Prepare storage for each sample's k-neighbors (and backup)
    lof.KNNs = make([][]int, addedIndex)
    lof.KNNsBackup = make([][]int, addedIndex)
    for idx := 0; idx < addedIndex; idx++ {
        lof.KNNs[idx] = make([]int, lof.MinPts)
        lof.KNNsBackup[idx] = make([]int, lof.MinPts)
    }
    // Throughout the train() method this value  is used for direct indexing
    // (i.e., not inside a for ...;...;... statement), so we need
    // to subtract 1 in order not to get out of range 
    addedIndex = lof.AddedIndex - 1
    numSamples = lof.NumSamples
    for idx, sample := range samples {
        sample.SetId(idx)  // Just additional info
    }
    // Compute distances between training samples
    for i := 0; i < numSamples; i++ {
        for j := 0; j < numSamples; j++ {
            if i == j {
                lof.Distances[i][j].Value = -1  // This is distinctive
                lof.Distances[i][j].Index = j
            } else {
                lof.Distances[i][j].Value = SampleDist(samples[i], samples[j])
                lof.Distances[j][i].Value = lof.Distances[i][j].Value
                lof.Distances[i][j].Index = j
                lof.Distances[j][i].Index = i
            }
        }
        // Set the additional slot's last value
        lof.Distances[addedIndex][addedIndex].Value = 0
        lof.Distances[addedIndex][addedIndex].Index = addedIndex
        lof.updateNNTable(i, "train")
    }
    // Save the nearest neighbors table state in the backup storage 
    for i := 0; i < numSamples; i++ {
        for k := 0; k < lof.MinPts; k++ {
            lof.KNNsBackup[i][k] = lof.KNNs[i][k] 
        } 
    }
}

// Shortcut for getting LOF for many samples. See GetLOF() method.
func (lof *LOF) GetLOFs(samples []ISample, mode string) map[ISample]float64 {
    output := make(map[ISample]float64)
    for _, sample := range samples {
        output[sample] = lof.GetLOF(sample, mode)
    }
    return output
}

// Returns LOF value for a sample; does lots of things, read the original
// paper and see implemetation for details. Read below for @mode param
// explanation.
func (lof *LOF) GetLOF(added ISample, mode string) float64 {
    // Decision between updating nearest neighbors for all the samples
    // OR updating nearest neighbors only for nearest neighbors of the
    // added sample
    optimized := lof.checkOptimization(mode)
    // Throughout the GetLOF() method this value is mostly used for direct
    // indexing (i.e., not inside a for ...;...;... statement), so we
    // need to subtract 1 in order not to get out of range 
    addedIndex := lof.AddedIndex - 1
    // Update distances table with added sample
    for i := 0; i < lof.NumSamples; i++ {
        // Distance between current training
        // sample and the sample being added
        dist := DistItem {
            Value: SampleDist(added, lof.TrainingSet[i]),
            Index: addedIndex,
        }
        lof.Distances[i][addedIndex] = dist
        lof.Distances[addedIndex][i] = dist
        lof.Distances[addedIndex][i].Index = i
    }
    if optimized {
        // Fill nearest neighbors table for added sample
        // (but don't touch any other samples yet)
        lof.updateNNTable(addedIndex, "compute")
        // We want to update nearest neighbors table ONLY
        // for those samples that are the added sample's nearest
        // neighbors; this adds some error, but saves CPU time
        for _, neighborIndex := range lof.KNNs[addedIndex] {
            lof.updateNNTable(neighborIndex, "compute")
        } 
    } else {
        // We want to update nearest neighbors for ALL samples;
        // don't forget we subtracted one from addedIndex for
        // direct indexing
        for idx := 0; idx < addedIndex + 1; idx++ {
            lof.updateNNTable(idx, "compute")  
        }
    }
    // Compute the LOF value
    addedDensity := lof.getDensity(addedIndex)
    neighborDensitySum := .0
    for _, neighborIndex := range lof.KNNs[addedIndex] {
        neighborDensitySum += lof.getDensity(neighborIndex)
    }
    factor := (neighborDensitySum / addedDensity) / float64(lof.MinPts)
    return factor
}

// This function resets the LOF.KNNs table to the state right after
// training. This may be necessery because after each GetLOF() call
// in the "fast" mode the table may get a bit distorted if some of
// the new samples happen to be somebody's nearest neighbors; thus
// there will be an accumulated error. This requires linear time.
func (lof *LOF) Reset() {
    log.Println("LOF: resetting NNs")
    for i := 0; i < lof.NumSamples; i++ {
        for k := 0; k < lof.MinPts; k++ {
            lof.KNNs[i][k] = lof.KNNsBackup[i][k] 
        } 
    }
}

// Given a sample's index in Distance table, update this sample's
// row in the nearest neighbors table. The @mode parameter
// controls whether we use the whole table row length (with added
// sample's slots, for LOF computation)
func (lof *LOF) updateNNTable(sampleIndex int, mode string) {
    bound := 0
    switch mode {
    case "train":
        bound = lof.NumSamples
    case "compute":
        bound = lof.AddedIndex
    default:
        log.Fatal("LOF: @mode should be either \"train\" or \"compute\"")
    }
    // Find nearest samples for current one: sort distances
    sorted := make([]DistItem, bound)
    copy(sorted, lof.Distances[sampleIndex])
    SortDistItems(sorted)
    // Find nearest samples for current one: take MinPts nearest
    for k := 1; k <= lof.MinPts; k++ {
        lof.KNNs[sampleIndex][k - 1] = sorted[k].Index
    }
}

// Returns Local Reachability Density for a sample (see the original
// paper for Local Reachability Density).
func (lof *LOF) getDensity(sampleIdx int) float64 {
    var distanceSum float64
    lastNeighborIdx := lof.MinPts - 1
    for _, neighborIdx := range lof.KNNs[sampleIdx] {
        // This is pre-computed distance between target sample
        // and his current nearest neighbors
        distance := lof.Distances[sampleIdx][neighborIdx].Value
        // Index of the farthest sample among the current neigbor's
        // nearest neighbors (will be used to retrieve the actual)
        // distance
        kDistanceIdx := lof.KNNs[neighborIdx][lastNeighborIdx]
        kDistance := lof.Distances[sampleIdx][kDistanceIdx].Value
        distanceSum += math.Max(math.Max(distance, kDistance), CMinimalSum)
    }
    return float64(lof.MinPts) / distanceSum
}

func (lof *LOF) checkOptimization(mode string) bool {
    var optimized bool
    switch mode {
    case "fast":
        optimized = true
    case "strict":
        optimized = false
    default:
        log.Fatal("LOF: @mode should be either \"fast\" or \"strict\"")
    }
    return optimized
}

func (lof *LOF) checkSamples(samples []ISample) {
    if len(samples) < lof.MinPts {
        log.Fatal(LNotEnoughSamples)
    }
}
