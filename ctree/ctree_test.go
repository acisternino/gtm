// Copyright (c) 2017 Andrea Cisternino. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package ctree

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddFirst(t *testing.T) {
	tree := New()

	f := &Frame{Address: 0, Length: 100}
	tree.Add(f)

	assert.Equal(t, 1, tree.Frames)
	assert.Equal(t, f, tree.root)
	assert.Nil(t, f.left)
	assert.Nil(t, f.right)
}

func TestAddLeft(t *testing.T) {
	tree := New()

	f2 := &Frame{Address: 100, Length: 100}
	tree.Add(f2)
	f1 := &Frame{Address: 1000, Length: 1000}
	tree.Add(f1)

	assert.Equal(t, 2, tree.Frames)
	assert.Equal(t, f1, tree.root)
	assert.Equal(t, f1.left, f2)
	assert.Nil(t, tree.root.right)
}

func TestAddRight(t *testing.T) {
	tree := New()

	f2 := &Frame{Address: 2000, Length: 100}
	tree.Add(f2)
	f1 := &Frame{Address: 100, Length: 1000}
	tree.Add(f1)

	assert.Equal(t, 2, tree.Frames)
	assert.Equal(t, f1, tree.root)
	assert.Equal(t, f1.right, f2)
	assert.Nil(t, tree.root.left)
}

func TestTraverse(t *testing.T) {
	tree := New()

	tree.Add(&Frame{Address: 0, Length: 100})
	tree.Add(&Frame{Address: 500, Length: 300})
	tree.Add(&Frame{Address: 1000, Length: 80})
	tree.Add(&Frame{Address: 1500, Length: 200})

	nodes := make([]string, 0, 10)
	tree.root.Traverse(func(f *Frame) error {
		nodes = append(nodes, f.String())
		return nil
	})
	result := strings.Join(nodes, "")
	assert.Equal(t, "[500,300][0,100][1500,200][1000,80]", result)
}
