// Copyright (c) 2017 Andrea Cisternino. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

/*
Package ctree implements a Cartesian Tree optimized for nodes with two
integer attributes: length and address.

This implementation is heavily inspired by the original implementation in
Tierra and is tailored for managing memory segments. A characteristic of this
implementation is that two constraints are always true.

The length of a node is always equal or larger to that of its children.

The address of a node is always larger or equal to those of its entire left
subtree. Resp. smaller than those of its right subtree.

From rule 1 derives that the root of the tree always contains the node with
the greates length. This characteristic is useful when using ctree to store
memory segments.

See more at https://en.wikipedia.org/wiki/Cartesian_tree
*/
package ctree

import (
	"math"
)

// Frame represents a segment of free memory.
type Frame struct {
	Address int32
	Length  int32
	left    *Frame
	right   *Frame
}

// These two variables are used as terminals with values that place them
// below every other node and at the extreme left and right.
// They simplify tree traversal avoiding explicit checks for nil.

var (
	lowestLeft  = &Frame{math.MinInt32, math.MinInt32, nil, nil}
	lowestRight = &Frame{math.MaxInt32, math.MinInt32, nil, nil}
)

// New return a new empty tree. The Frame returned is not part of the tree
// and is placed at the hypotetical top-left corner of the coordinate space
// defined by the length and address attributes. The "real" tree hangs at
// its right.
func New() *Frame {
	return &Frame{
		Address: math.MinInt32,
		Length:  math.MaxInt32,
		left:    lowestLeft,
		right:   lowestRight,
	}
}

// NewWithValues return a new tree with an initial root node already initialized.
func NewWithValues(address, length int32) *Frame {
	return &Frame{
		Address: math.MinInt32,
		Length:  math.MaxInt32,
		left:    lowestLeft,
		right:   &Frame{address, length, lowestLeft, lowestRight},
	}
}

// Add inserts a frame into the tree below the current one.
func (f *Frame) Add(nf *Frame) {
	loop := true
	current := f
	parent := f

	for loop {
		if nf.Length <= current.Length {
			// new frame is smaller than current, descend according to address
			parent = current
			if nf.Address <= current.Address {
				// descend left
				current = current.left
			} else {
				// descend right
				current = current.right
			}
		} else {
			// new frame is larger than current, insert here
			if nf.Address <= parent.Address {
				// append left
				nf.left = parent.left
				nf.right = lowestRight
				parent.left = nf
			} else {
				// append right
				nf.right = parent.right
				nf.left = lowestLeft
				parent.right = nf
			}
			loop = false
		}
	}
}
