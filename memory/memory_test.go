// Copyright (c) 2017 Andrea Cisternino. All rights reserved.

package memory

import "testing"

var (
	names = map[int]string{
		right:      "Right",
		left:       "Left",
		touchRight: "TouchRight",
		touchLeft:  "TouchLeft",
		overlaps:   "Overlaps",
	}
)

func TestPositionRight(t *testing.T) {
	t.Log("Other frame is to the right")

	this := &frame{40, 10, nil, nil}
	other := &frame{50, 80, nil, nil}

	if res := this.position(other); res != right {
		t.Errorf("Expected \"Right\", but was %v instead.", names[res])
	}
}

func TestPositionTouchesRight(t *testing.T) {
	t.Log("Other frame touches to the right")

	this := &frame{40, 10, nil, nil}
	other := &frame{50, 50, nil, nil}

	if res := this.position(other); res != touchRight {
		t.Errorf("Expected \"TouchRight\", but was %v instead.", names[res])
	}
}

func TestPositionLeft(t *testing.T) {
	t.Log("Other frame is to the left")

	this := &frame{40, 100, nil, nil}
	other := &frame{20, 10, nil, nil}

	if res := this.position(other); res != left {
		t.Errorf("Expected \"Left\", but was %v instead.", names[res])
	}
}

func TestPositionTouchesLeft(t *testing.T) {
	t.Log("Other frame touches to the left")

	this := &frame{40, 100, nil, nil}
	other := &frame{20, 80, nil, nil}

	if res := this.position(other); res != touchLeft {
		t.Errorf("Expected \"TouchLeft\", but was %v instead.", names[res])
	}
}

func TestPositionOverlapsLeft(t *testing.T) {
	t.Log("Other frame overlaps left")

	this := &frame{40, 100, nil, nil}
	other := &frame{40, 80, nil, nil}

	if res := this.position(other); res != overlaps {
		t.Errorf("Expected \"Overlaps\", but was %v instead.", names[res])
	}
}

func TestPositionOverlapsRight(t *testing.T) {
	t.Log("Other frame overlaps right")

	this := &frame{40, 100, nil, nil}
	other := &frame{40, 120, nil, nil}

	if res := this.position(other); res != overlaps {
		t.Errorf("Expected \"Overlaps\", but was %v instead.", names[res])
	}
}

func TestPositionOverlapsFully(t *testing.T) {
	t.Log("Other frame overlaps fully")

	this := &frame{100, 100, nil, nil}
	other := &frame{40, 120, nil, nil}

	if res := this.position(other); res != overlaps {
		t.Errorf("Expected \"Overlaps\", but was %v instead.", names[res])
	}
}
