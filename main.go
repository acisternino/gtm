// Copyright (c) 2017 Andrea Cisternino. All rights reserved.

package main

import (
	"fmt"
)

// Frame is a segment of memory
type Frame struct {
	addr, length int
	left, right  int
}

const (
	// SIZE is the dimension of the soup
	SIZE = 50
)

var (
	frames = make([]Frame, 0, 10)
	root   int
)

func main() {
	fmt.Println("start")

	fmt.Printf("soup size: %d\n", SIZE)

	// init data structures
	frames = append(frames, Frame{0, 0, 0, 0}) // tail frame
	frames = append(frames, Frame{0, SIZE, 0, 0})
	root = 1

	// try to alloc 10

	fmt.Println(frames)
}

func alloc(size int) {
	fmt.Printf("allocating %d bytes", size)

	currentFrame := root

	for {
		frame := frames[currentFrame]

		if frame.length < size {
			break
		}

		left := frames[frame.left]
		right := frames[frame.right]

		if left.length > size {
			currentFrame = frame.left
		}
		if right.length > size {
			currentFrame = frame.right
		}
	}

	// here currentFrame points to the right one
}

/*
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
*/
