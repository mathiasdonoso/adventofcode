package day1

import "testing"

func assertEquals(t *testing.T, spected interface{}, returned interface{}) {
	if spected != returned {
		t.Fatal("It should be", spected, "but instead it returns", returned)
	}
}

func TestChronalCalibration(t *testing.T) {
	test1 := ChronalCalibration("+1\n+1\n+1")
	assertEquals(t, 3, test1)

	test2 := ChronalCalibration("+1\n+1\n-2")
	assertEquals(t, 0, test2)

	test3 := ChronalCalibration("-1\n-2\n-3")
	assertEquals(t, -6, test3)
}
