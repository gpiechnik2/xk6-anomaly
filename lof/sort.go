package lof

import (
    "sort"
)

// This function allows you to sort a slice of DistItem values in ascending
// order by the .Value field; the implementation is inspired by (or even
// copied from) https://golang.org/pkg/sort/#example__sortKeys  
func SortDistItems(distances []DistItem) {
    distance := func(s1 DistItem, s2 DistItem) bool {
        return s1.Value < s2.Value
    }
    By(distance).Sort(distances)
}

// By is the type of a "less" function that defines the ordering of its
// Sample arguments.
type By func(p1, p2 DistItem) bool

// SampleSorter joins a By function and a slice of Distances to be sorted.
type SampleSorter struct {
    Distances []DistItem
    By      func(p1, p2 DistItem) bool  // Closure used in the Less method
}

// Len is part of sort.Interface.
func (s *SampleSorter) Len() int {

    return len(s.Distances)
}

// Swap is part of sort.Interface.
func (s *SampleSorter) Swap(i, j int) {

    s.Distances[i], s.Distances[j] = s.Distances[j], s.Distances[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" 
// closure in the sorter.
func (s *SampleSorter) Less(i, j int) bool {

    return s.By(s.Distances[i], s.Distances[j])
}

// Sort is a method on the function type, By, that sorts the argument slice
// according to the function.
func (by By) Sort(distances []DistItem) {

    ps := &SampleSorter{
        Distances: distances,
        // The Sort method's receiver is the function
        // (closure) that defines the sort order
        By: by, 
    }
    sort.Sort(ps)
}

