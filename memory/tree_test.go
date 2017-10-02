// Copyright (c) 2017 Andrea Cisternino. All rights reserved.

package memory

import "testing"

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
	if f1.right != lowestRight {
		t.Errorf("Top node right leaf doesn't match. Was: %v", f1.right)
	}
	if f1.left != f2 {
		t.Errorf("Top node left leaf doesn't match. Was: %v", f1.left)
	}
	if f2.right != lowestRight {
		t.Errorf("Leaf node right leaf doesn't match. Was: %v", f2.right)
	}
	if f2.left != lowestLeft {
		t.Errorf("Leaf node left leaf doesn't match. Was: %v", f2.left)
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
	if f1.right != f2 {
		t.Errorf("Top node right leaf doesn't match. Was: %v", f1.right)
	}
	if f1.left != lowestLeft {
		t.Errorf("Top node left leaf doesn't match. Was: %v", f1.left)
	}
	if f2.right != lowestRight {
		t.Errorf("Leaf node right leaf doesn't match. Was: %v", f2.right)
	}
	if f2.left != lowestLeft {
		t.Errorf("Leaf node left leaf doesn't match. Was: %v", f2.left)
	}
}
