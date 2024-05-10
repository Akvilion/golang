package prose

import "testing"

func TestTwoElements(t *testing.T) {
	//t.Error("no test written yet")
	list := []string{"apple", "orange"}
	if JoinWithCommas(list) != "apple and orange" {
		t.Error("didn't match expected value")
	}
}
func TestThreeElements(t *testing.T) {
	//t.Error("no test written yet")
}
