// Copyright (c) 2017 Andrea Cisternino. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package ctree

import (
	"testing"
)

func TestPopEmpty(t *testing.T) {
	stack := newStack()

	if !stack.empty() {
		t.Errorf("Expected an empty stack.")
	}
	if _, err := stack.pop(); err != errEmptyStack {
		t.Errorf("Popping from an empty stack. Expected \"%v\", was: \"%v\"", errEmptyStack, err)
	}
	if d := stack.depth(); d != 0 {
		t.Errorf("Wrong stack depth. Expected 0, was: \"%v\"", d)
	}
}

func TestPushOne(t *testing.T) {
	stack := newStack()

	fp := &Frame{Address: 0, Length: 100}
	stack.push(fp)

	if d := stack.depth(); d != 1 {
		t.Errorf("Wrong stack depth. Expected 1, was: \"%v\"", d)
	}
	result, err := stack.pop()
	if err != nil {
		t.Errorf("Unexpected error while popping: \"%v\"", err)
	}
	if result != fp {
		t.Errorf("Wrong frame: Expected \"%v\", was \"%v\"", fp, result)
	}
}

func TestPushTwo(t *testing.T) {
	stack := newStack()

	fp1 := &Frame{Address: 0, Length: 100}
	stack.push(fp1)
	fp2 := &Frame{Address: 10, Length: 1000}
	stack.push(fp2)

	if d := stack.depth(); d != 2 {
		t.Errorf("Wrong stack depth. Expected 2, was: \"%v\"", d)
	}
	result, err := stack.pop()
	if err != nil {
		t.Errorf("Unexpected error while popping: \"%v\"", err)
	}
	if result != fp2 {
		t.Errorf("Wrong frame: Expected \"%v\", was \"%v\"", fp2, result)
	}
	if d := stack.depth(); d != 1 {
		t.Errorf("Wrong stack depth. Expected 1, was: \"%v\"", d)
	}
	if stack.empty() {
		t.Errorf("Expected a non empty stack.")
	}
}

func TestAfterPop(t *testing.T) {
	stack := newStack()

	fp := &Frame{Address: 0, Length: 100}
	stack.push(fp)

	_, err := stack.pop()
	if err != nil {
		t.Errorf("Unexpected error while popping: \"%v\"", err)
	}

	if d := stack.depth(); d != 0 {
		t.Errorf("Wrong stack depth. Expected 0, was: \"%v\"", d)
	}
	if !stack.empty() {
		t.Errorf("Expected an empty stack.")
	}
}
