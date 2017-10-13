// Copyright (c) 2017 Andrea Cisternino. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

/*
Package ctree implements a Cartesian Tree optimized for nodes with two
integer attributes: length and address.

This implementation is heavily inspired by the original implementation in
Tierra and is tailored for managing memory segments. A characteristic of this
implementation is that two constraints are always true.

1. The length of a node is always equal or larger to that of its children.

2. The address of a node is always larger or equal to those of its entire
   left subtree. Resp. smaller than those of its right subtree.

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

//---- Frame -------------------------------------------------------------

// Frame represents a segment of free memory.
type Frame struct {
	Address int32
	Length  int32
	left    *Frame
	right   *Frame
}

// Add adds a new Frame to the subtree anchored at this Frame.
func (f *Frame) Add(nf, parent *Frame) {
	current := f
	loop := true

	for loop {
		if current != nil && nf.Length <= current.Length {
			// new frame is smaller than current, descend according to address
			parent = current
			if nf.Address <= current.Address {
				current = current.left
			} else {
				current = current.right
			}
		} else {
			// new frame is larger than current or we reached the end: insert here
			if nf.Address <= parent.Address {
				// append left
				nf.left = current // was parent.left
				nf.right = nil
				parent.left = nf
				// rebalance current.right
				nf.rebalanceLeft(parent)
			} else {
				// append right
				nf.right = current // was parent.right
				nf.left = nil
				parent.right = nf
				// rebalance current.left
			}
			loop = false
		}
	}
}

// Traverse the subtree anchored at f in pre-order depth-first order.
// Each node is provided to function visit for processing.
func (f *Frame) Traverse(visit func(*Frame) error) error {
	s := newStackWith(f)
	for !s.empty() {
		current, err := s.pop()
		if err != nil {
			return errors.Wrap(err, "stack error")
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

// rebalanceLeft re-distributes child nodes so that their addresses are
// correctly sorted.
func (f *Frame) rebalanceLeft(parent *Frame) {
	if f.left == nil || f.left.right == nil {
		return
	}
	// detach affected nodes
	branch := f.left.right
	f.left.right = nil

	branch.Traverse(func(tf *Frame) error {
		fmt.Println("rebalancing ", tf)
		f.Add(tf, parent)
		return nil
	})
}

// String returns a string representation of the Frame.
func (f Frame) String() string {
	return fmt.Sprintf("[%d,%d]", f.Address, f.Length)
}

//---- CTree -------------------------------------------------------------

// CTree is the cartesian tree containing the free memory segments.
type CTree struct {
	root   *Frame
	Frames int
}

// New return a new tree with an initial root node of the given length
// and located at address 0.
func New(length int32) *CTree {
	return &CTree{
		root:   &Frame{Address: 0, Length: length},
		Frames: 1,
	}
}

// Add inserts a frame to the tree. This function handles cases related to
// the root of the tree. Once the root cases have been handled, it delegates
// to the Frame's method with the same name.
func (t *CTree) Add(nf *Frame) {
	if nf.Length >= t.root.Length {
		// root insertion
		if nf.Address <= t.root.Address {
			// new is to the left
			nf.right = t.root
		} else {
			// new is to the right
			nf.left = t.root
		}
		t.Frames++
		t.root = nf
		return
	}

	// root is larger, add to one of the subtrees
	if nf.Address <= t.root.Address {
		if t.root.left == nil {
			t.root.left = nf
		} else {
			t.root.left.Add(nf, t.root)
		}
	} else {
		if t.root.right == nil {
			t.root.right = nf
		} else {
			t.root.right.Add(nf, t.root)
		}
	}
	t.Frames++
}
