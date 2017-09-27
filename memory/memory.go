// Copyright (c) 2017 Andrea Cisternino. All rights reserved.

package memory

import (
	"math"
)

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
	right = iota
	left
	touchRight
	touchLeft
	overlaps
)

var (
	lowestLeft  = &frame{math.MinInt32, math.MinInt32, nil, nil}
	lowestRight = &frame{math.MinInt32, math.MaxInt32, nil, nil}
)

func newTree() *frame {
	return &frame{
		length:  math.MaxInt32,
		address: math.MinInt32,
		left:    lowestLeft,
		right:   lowestRight,
	}
}

// insert inserts a frame into the tree below the current one
func (f *frame) insert(newFrame *frame) {

	loop := true
	node := f

	for loop {
		if newFrame.length <= node.length {
			// new node is smaller than current, go down according to address

			if newFrame.address <= node.address {
				// descend left
				node = node.left
			} else {
				// descend right
				node = node.right
			}

		} else {
			// new node is larger than current, insert here

			f.right = newFrame
			newFrame.right = node
			newFrame.left = lowestLeft
			loop = false

			newFrame.left = node
			newFrame.right = lowestRight
			// how do I change the parent?  parent.left = newFrame
		}
	}
}

func (f *frame) position(other *frame) int {

	end := f.address + f.length
	otherEnd := other.address + other.length

	switch {
	case other.address > end:
		return right
	case otherEnd < f.address:
		return left
	case other.address == end:
		return touchRight
	case otherEnd == f.address:
		return touchLeft
	default:
		return overlaps
	}
}
