package geography

import (
	"testing"
)

func TestIsInCircleFence(t *testing.T) {

	if IsInCircleFence(50000, 116.398232, 39.926224, 116.510341, 39.900102) {
		t.Log("In circle!")
	} else {
		t.Log("Not in circle")
	}

}
