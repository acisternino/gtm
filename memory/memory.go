// Copyright (c) 2017 Andrea Cisternino. All rights reserved.

package memory

type frame struct {
	length  int32
	address int32
	left    *frame
	right   *frame
}

type tree struct {
	root *frame
}

const (
	Right = iota
	Left
	TouchRight
	TouchLeft
	Overlaps
)

func (f *frame) position(other *frame) int {

	end := f.address + f.length
	otherEnd := other.address + other.length

	switch {
	case other.address > end:
		return Right
	case otherEnd < f.address:
		return Left
	case other.address == end:
		return TouchRight
	case otherEnd == f.address:
		return TouchLeft
	default:
		return Overlaps
	}
}
