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
	"fmt"

	"github.com/pkg/errors"
)

// Frame represents a segment of free memory.
type Frame struct {
	Address int32
	Length  int32
	left    *Frame
	right   *Frame
}

// Traverse the subtree anchored at f in pre-order depth-first order.
// Each node is provided to function visit for processing.
func (f *Frame) Traverse(visit func(*Frame) error) error {
	s := newStackWith(f)
	for !s.empty() {
		current, err := s.pop()
		if err != nil {
			return errors.Wrap(err, "read failed")
		}
		if err = visit(current); err != nil {
			return err
		}
		// by pushing right first, we visit the left subtree first
		if current.right != nil {
			s.push(current.right)
		}
		if current.left != nil {
			s.push(current.left)
		}
	}
	return nil
}

// String returns a string representation of the Frame.
func (f Frame) String() string {
	return fmt.Sprintf("[%d,%d]", f.Address, f.Length)
}

// CTree is the cartesian tree containing the free memory segments.
type CTree struct {
	root   *Frame
	Frames int
}

// New return a new empty tree.
func New() *CTree {
	return new(CTree)
}

// NewWithLength return a new tree with an initial root node already
// initialized with the given length and address 0.
func NewWithLength(length int32) *CTree {
	ct := New()
	ct.Add(&Frame{Address: 0, Length: length})
	return ct
}

// Add inserts a frame to the tree.
// TODO() check for overlap of ranges
func (t *CTree) Add(f *Frame) {
	if t.root == nil {
		t.root = f
		t.Frames++
		return
	}

	current := t.root
	parent := (*Frame)(nil)

	loop := true

	for loop {
		if current != nil && f.Length <= current.Length {
			// new frame is smaller than current, descend according to address
			parent = current
			if f.Address <= current.Address {
				// descend left
				current = current.left
			} else {
				// descend right
				current = current.right
			}
		} else {
			// new frame is larger than current or we reached the end: insert here
			if parent == nil {
				// root insertion
				t.root = f
				// here current points to the old root
				if current.Address <= f.Address {
					f.left = current
				} else {
					f.right = current
				}
			} else {
				if f.Address <= parent.Address {
					// append left
					f.left = parent.left
					f.right = nil
					parent.left = f
				} else {
					// append right
					f.right = parent.right
					f.left = nil
					parent.right = f
				}
			}
			loop = false
		}
	}
	t.Frames++
}
