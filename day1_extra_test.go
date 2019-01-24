package main

import "testing"

func TestHalfwayAroundCaptcha(t *testing.T) {
	test1 := HalfwayAroundCaptcha("1212")
	assertEquals(t, 6, test1)

	test2 := HalfwayAroundCaptcha("1221")
	assertEquals(t, 0, test2)

	test3 := HalfwayAroundCaptcha("123425")
	assertEquals(t, 4, test3)

	test4 := HalfwayAroundCaptcha("123123")
	assertEquals(t, 12, test4)

	test5 := HalfwayAroundCaptcha("12131415")
	assertEquals(t, 4, test5)
}
