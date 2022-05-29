package modules

import "testing"

var Debug bool = false

func TestAvarage(t *testing.T) {
	if Debug {
		t.Skip("Do Skip!!")
	}

	v := Average([]int{1, 2, 3, 4, 5})

	if v != 3 {
		t.Errorf("%v != %v", v, 3)
	}
}
