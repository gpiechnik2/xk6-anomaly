package lof

//////////////////////////////////////////////////////////////////////////////
//
//  Interface for anything that must be treated as a point in a sample space.
//
//////////////////////////////////////////////////////////////////////////////

type ISample interface {
    GetId() int
    SetId(int)
    GetPoint() []float64
    SetPoint([]float64)
}

////////////////////////////////////////////////////////////////////////////// 
//
//  This type is only used with GetSamplesFromFloat64s().                   
//
////////////////////////////////////////////////////////////////////////////// 

type BasicSample struct {
    Id          int
    Point       []float64
    Distance    float64
}

// Constructor for BasicSample type.
func NewBasicSample(id int, point []float64) *BasicSample {

    return &BasicSample{Point: point, Id: id}
}

// Given a slice of float64 slices, build BasicSamples
// treating each float64 slice as a point in sample space.
func GetSamplesFromFloat64s(points [][]float64) []ISample {

    bSpl := []ISample{}
    for idx, pt := range points {
        bSpl = append(bSpl, NewBasicSample(idx, pt))
    }

    return bSpl
}

// Satisfies ISample interface.
func (bs *BasicSample) GetId() int {

    return bs.Id
}

// Satisfies ISample interface.
func (bs *BasicSample) SetId(id int) {

    bs.Id = id
}

// Satisfies ISample interface.
func (bs *BasicSample) GetPoint() []float64 {

    return bs.Point
}

// Satisfies ISample interface.
func (bs *BasicSample) SetPoint(point []float64) {

    bs.Point = point
}
