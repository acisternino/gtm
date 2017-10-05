// Copyright (c) 2017 Andrea Cisternino. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package ctree

import (
	"fmt"
	"testing"
)

func (f *Frame) checkNode(expectedLeft, expectedRight *Frame, name string) error {
	if f.right != expectedRight {
		return fmt.Errorf("%s right leaf doesn't match. Was: %v", name, f.right)
	}
	if f.left != expectedLeft {
		return fmt.Errorf("%s left leaf doesn't match. Was: %v", name, f.left)
	}
	return nil
}

func TestInsertFirst(t *testing.T) {
	tree := New()

	f := &Frame{Length: 100, Address: 0}
	tree.Add(f)

	if tree.right != f {
		t.Errorf("Root right leaf doesn't match. Was: %v", tree.right)
	}
	if f.right != lowestRight {
		t.Errorf("Inserted node right leaf doesn't match. Was: %v", f.right)
	}
	if f.left != lowestLeft {
		t.Errorf("Inserted node left leaf doesn't match. Was: %v", f.left)
	}
}

func TestInsertLeft(t *testing.T) {
	tree := New()

	f1 := &Frame{Length: 1000, Address: 1000}
	f2 := &Frame{Length: 100, Address: 100}
	tree.Add(f1)
	tree.Add(f2)

	if tree.right != f1 {
		t.Errorf("Root right leaf doesn't match. Was: %v", tree.right)
	}

	if err := f1.checkNode(f2, lowestRight, "Top node"); err != nil {
		t.Errorf(err.Error())
	}
	if err := f2.checkNode(lowestLeft, lowestRight, "Leaf node"); err != nil {
		t.Errorf(err.Error())
	}
}

func TestInsertRight(t *testing.T) {
	tree := New()

	f1 := &Frame{Length: 1000, Address: 100}
	f2 := &Frame{Length: 100, Address: 2000}
	tree.Add(f1)
	tree.Add(f2)

	if tree.right != f1 {
		t.Errorf("Root right leaf doesn't match. Was: %v", tree.right)
	}

	if err := f1.checkNode(lowestLeft, f2, "Top node"); err != nil {
		t.Errorf(err.Error())
	}
	if err := f2.checkNode(lowestLeft, lowestRight, "Leaf node"); err != nil {
		t.Errorf(err.Error())
	}
}
