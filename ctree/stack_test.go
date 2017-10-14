// Copyright (c) 2017 Andrea Cisternino. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package ctree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	wrongDepthMsg  = "Wrong stack depth"
	wrongObjectMsg = "Wrong object"
)

func TestPopEmpty(t *testing.T) {
	stack := newStack()

	assert.True(t, stack.empty())

	_, err := stack.pop()
	assert.EqualError(t, err, errEmptyStack.Error())
	assert.Equalf(t, 0, stack.depth(), wrongDepthMsg)
}

func TestPushOne(t *testing.T) {
	stack := newStack()

	fp := &Frame{Address: 0, Length: 100}
	stack.push(fp)

	assert.Equalf(t, 1, stack.depth(), wrongDepthMsg)

	result, err := stack.pop()
	if assert.NoError(t, err) {
		assert.Equalf(t, fp, result, wrongObjectMsg)
	}
}

func TestPushTwo(t *testing.T) {
	stack := newStack()

	fp1 := &Frame{Address: 0, Length: 100}
	stack.push(fp1)
	fp2 := &Frame{Address: 10, Length: 1000}
	stack.push(fp2)

	assert.Equalf(t, 2, stack.depth(), wrongDepthMsg)

	result, err := stack.pop()
	if assert.NoError(t, err) {
		assert.Equalf(t, fp2, result, wrongObjectMsg)
		assert.Equalf(t, 1, stack.depth(), wrongDepthMsg)
		assert.False(t, stack.empty())
	}
}

func TestAfterPop(t *testing.T) {
	stack := newStack()

	fp := &Frame{Address: 0, Length: 100}
	stack.push(fp)

	_, err := stack.pop()
	if assert.NoError(t, err) {
		assert.Equalf(t, 0, stack.depth(), wrongDepthMsg)
		assert.True(t, stack.empty())
	}
}

func TestPeekEmpty(t *testing.T) {
	stack := newStack()

	assert.Nil(t, stack.peek())
}

func TestPeek(t *testing.T) {
	stack := newStack()

	fp1 := &Frame{Address: 0, Length: 100}
	stack.push(fp1)
	fp2 := &Frame{Address: 10, Length: 1000}
	stack.push(fp2)

	assert.Equal(t, fp2, stack.peek(), wrongObjectMsg)
	_, err := stack.pop()
	if assert.NoError(t, err) {
		assert.Equal(t, fp1, stack.peek(), wrongObjectMsg)
	}
}
