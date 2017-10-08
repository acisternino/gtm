// Copyright (c) 2017 Andrea Cisternino. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// A simple implementation of stack optimized for keeping pointers to Frames.
// Used when doing depth-first traversal of the tree. Loosely insipred by
// https://github.com/karalabe/cookiejar/tree/master/collections/stack

package ctree

import "errors"

var errorEmptyStack = errors.New("Illegal operation: stack is empty")

// blockSize is the default number of slots in the stack.
const blockSize = 64

// stack is a last-in-first-out (LIFO) stack of pointers to Frames.
type stack struct {
	frames  []*Frame
	current int
}

// newStack returns a new empty stack of default size (depth)
func newStack() *stack {
	s := new(stack)
	s.frames = make([]*Frame, blockSize)
	s.current = -1
	return s
}

// newStackWith returns a new stack of default size (depth) with the
// given item already at the top.
func newStackWith(f *Frame) *stack {
	s := newStack()
	s.push(f)
	return s
}

// push a Frame pointer onto the top of the stack.
func (s *stack) push(f *Frame) {
	s.current++
	s.frames[s.current] = f

	if s.current == len(s.frames)-1 {
		newSlice := make([]*Frame, cap(s.frames)+blockSize)
		copy(newSlice, s.frames)
		s.frames = newSlice
	}
}

// pop removes a Frame pointer from the top of the stack and returns it.
func (s *stack) pop() (*Frame, error) {
	if s.current < 0 {
		return nil, errorEmptyStack
	}
	f := s.frames[s.current]
	s.current--
	return f, nil
}

// depth return the number of items currently in the stack.
func (s stack) depth() int {
	return s.current + 1
}

// empty tests if this stack is empty.
func (s stack) empty() bool {
	return s.current == -1
}
