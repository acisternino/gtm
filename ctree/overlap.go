// Copyright (c) 2017 Andrea Cisternino. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// This file contains code for handling the overlap of memory Frames

package ctree

const (
	right = iota
	left
	touchRight
	touchLeft
	overlaps
)

// position returns a value that encodes the relative position of "other"
// with respect to "f".
func (f *Frame) position(other *Frame) int {

	end := f.Address + f.Length
	otherEnd := other.Address + other.Length

	switch {
	case other.Address > end:
		return right
	case otherEnd < f.Address:
		return left
	case other.Address == end:
		return touchRight
	case otherEnd == f.Address:
		return touchLeft
	default:
		return overlaps
	}
}
