// Copyright (c) 2017 Andrea Cisternino. All rights reserved.

package memory

type frame struct {
	length  int32
	address int32
	left    *frame
	right   *frame
}

type tree struct {
	root *frame
}

func newTree(initialSize int32) (t *tree) {
	f := &frame{
		length:  initialSize,
		address: 0,
	}
	t = &tree{
		root: f,
	}
	return
}
