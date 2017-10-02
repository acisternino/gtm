// Copyright (c) 2017 Andrea Cisternino. All rights reserved.

package memory

import (
	"fmt"
	"testing"
)

func (f *frame) checkNode(expectedLeft, expectedRight *frame, name string) error {
	if f.right != expectedRight {
		return fmt.Errorf("%s right leaf doesn't match. Was: %v", name, f.right)
	}
	if f.left != expectedLeft {
		return fmt.Errorf("%s left leaf doesn't match. Was: %v", name, f.left)
	}
	return nil
}

func TestInsertFirst(t *testing.T) {
	tree := newTree()

	f := &frame{length: 100, address: 0}
	tree.insert(f)

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
	tree := newTree()

	f1 := &frame{length: 1000, address: 1000}
	f2 := &frame{length: 100, address: 100}
	tree.insert(f1)
	tree.insert(f2)

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
	tree := newTree()

	f1 := &frame{length: 1000, address: 100}
	f2 := &frame{length: 100, address: 2000}
	tree.insert(f1)
	tree.insert(f2)

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
