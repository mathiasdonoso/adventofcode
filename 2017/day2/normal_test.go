package day2

import "testing"

func assertEquals(t *testing.T, spected interface{}, returned interface{}) {
	if spected != returned {
		t.Fatal("It should be", spected, "but instead it returns", returned)
	}
}

func TestCorruptionChecksum(t *testing.T) {
	data := "5\t1\t9\t5\n7\t5\t3\n2\t4\t6\t8"

	test1 := CorruptionChecksum(data)
	assertEquals(t, 18, test1)
}
