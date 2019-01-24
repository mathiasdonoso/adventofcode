package main

import "testing"

func assertEquals(t *testing.T, spected interface{}, returned interface{}) {
	if spected != returned {
		t.Fatal("It should be", spected, "but instead it returns", returned)
	}
}

func TestInverseCaptcha(t *testing.T) {
	test1 := InverseCaptcha("1122")
	assertEquals(t, 3, test1)

	test2 := InverseCaptcha("1111")
	assertEquals(t, 4, test2)

	test3 := InverseCaptcha("1234")
	assertEquals(t, 0, test3)

	test4 := InverseCaptcha("91212129")
	assertEquals(t, 9, test4)
}
