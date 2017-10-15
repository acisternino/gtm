// Copyright (c) 2017 Andrea Cisternino. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package ctree

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootAddAboveRight(t *testing.T) {
	tree := New(50) // [0,50]
	r := tree.root

	f := &Frame{Address: 100, Length: 100}
	tree.Add(f)

	assert.Equal(t, 2, tree.Frames)
	assert.Equal(t, f, tree.root)
	assert.Equal(t, r, tree.root.left)
	assert.Nil(t, tree.root.right)
}

func TestRootAddAboveLeft(t *testing.T) {
	tree := New(50) // [0,50]
	// first add something to the right
	r := &Frame{Address: 200, Length: 100}
	tree.Add(r)
	// now add something to the left
	f := &Frame{Address: 100, Length: 200}
	tree.Add(f)

	assert.Equal(t, 3, tree.Frames)
	assert.Equal(t, f, tree.root)
	assert.Equal(t, r, tree.root.right)
}

func TestRootAddRight(t *testing.T) {
	tree := New(100) // [0,100]

	f := &Frame{Address: 200, Length: 50}
	tree.Add(f)

	assert.Equal(t, 2, tree.Frames)
	assert.Equal(t, int32(100), tree.root.Length)
	assert.Equal(t, f, tree.root.right)
	assert.Nil(t, tree.root.left)
}

func TestRootAddLeft(t *testing.T) {
	tree := New(100)        // [0,100]
	tree.root.Address = 100 // move root to the right

	f := &Frame{Address: 0, Length: 50}
	tree.Add(f)

	assert.Equal(t, 2, tree.Frames)
	assert.Equal(t, int32(100), tree.root.Length)
	assert.Equal(t, f, tree.root.left)
	assert.Nil(t, tree.root.right)
}

func TestAddLeft(t *testing.T) {
	tree := New(200)        // [0,200]
	tree.root.Address = 200 // move root to the right -> [200,200]
	tree.Add(&Frame{Address: 0, Length: 50})
	tree.Add(&Frame{Address: 300, Length: 100})

	f := &Frame{Address: 100, Length: 100}
	tree.Add(f)

	assert.Equal(t, 4, tree.Frames)

	actual := tree.root.left
	assert.Equal(t, f, actual)
	assert.Equal(t, int32(100), actual.Length)
	assert.Equal(t, int32(50), actual.left.Length)
	assert.Nil(t, actual.right)
}

func TestAddRight(t *testing.T) {
	tree := New(300)        // [0,300]
	tree.root.Address = 100 // move root to the right -> [100,300]
	tree.Add(&Frame{Address: 0, Length: 50})
	tree.Add(&Frame{Address: 600, Length: 20})

	f := &Frame{Address: 500, Length: 50}
	tree.Add(f)

	assert.Equal(t, 4, tree.Frames)

	actual := tree.root.right
	assert.Equal(t, f, actual)
	assert.Equal(t, int32(50), actual.Length)
	assert.Equal(t, int32(20), actual.right.Length)
	assert.Nil(t, actual.left)
}

func TestTraversePre(t *testing.T) {
	tree := New(100)
	tree.Add(&Frame{Address: 200, Length: 80})
	tree.Add(&Frame{Address: 500, Length: 300})
	tree.Add(&Frame{Address: 1000, Length: 80})
	tree.Add(&Frame{Address: 1500, Length: 200})
	tree.Add(&Frame{Address: 2000, Length: 100})

	nodes := make([]string, 0, 10)
	tree.root.TraversePre(func(f *Frame) error {
		nodes = append(nodes, f.String())
		return nil
	})
	result := strings.Join(nodes, "")
	assert.Equal(t, "[500,300][0,100][200,80][1500,200][1000,80][2000,100]", result)
}

func TestTraversePost(t *testing.T) {
	tree := New(100)
	tree.Add(&Frame{Address: 200, Length: 80})
	tree.Add(&Frame{Address: 500, Length: 300})
	tree.Add(&Frame{Address: 1000, Length: 80})
	tree.Add(&Frame{Address: 1500, Length: 200})
	tree.Add(&Frame{Address: 2000, Length: 100})

	nodes := make([]string, 0, 10)
	tree.root.TraversePost(func(f *Frame) error {
		nodes = append(nodes, f.String())
		return nil
	})
	result := strings.Join(nodes, "")
	assert.Equal(t, "[200,80][0,100][1000,80][2000,100][1500,200][500,300]", result)
}

func NewFrame(address, length int32) *Frame {
	return &Frame{address, length, nil, nil}
}

func TestRebalanceRight(t *testing.T) {
	tree := New(5)
	tree.Add(NewFrame(250, 20))
	tree.Add(NewFrame(280, 10))
	tree.Add(NewFrame(350, 15))
	tree.Add(NewFrame(470, 35))

	tree.Add(NewFrame(300, 25)) // triggers rebalance

	nodes := make([]string, 0, 10)
	tree.root.TraversePre(func(f *Frame) error {
		nodes = append(nodes, f.String())
		return nil
	})
	result := strings.Join(nodes, "")
	assert.Equal(t, "[470,35][300,25][250,20][0,5][280,10][350,15]", result)
}

func TestRebalanceLeft(t *testing.T) {
	tree := New(20)
	tree.Add(NewFrame(130, 40))
	tree.Add(NewFrame(410, 5))
	tree.Add(NewFrame(210, 20))
	tree.Add(NewFrame(180, 25))
	tree.Add(NewFrame(500, 30))
	tree.Add(NewFrame(630, 10))
	tree.Add(NewFrame(700, 20))

	tree.Add(NewFrame(300, 35)) // triggers rebalance

	nodes := make([]string, 0, 12)
	tree.root.TraversePre(func(f *Frame) error {
		nodes = append(nodes, f.String())
		return nil
	})
	result := strings.Join(nodes, "")
	assert.Equal(t, "[130,40][0,20][300,35][180,25][210,20][500,30][410,5][700,20][630,10]", result)
}
