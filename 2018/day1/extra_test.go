package day1

import "testing"

func TestTwiceCalibration(t *testing.T) {
	test1 := TwiceCalibration("+1\n-1")
	assertEquals(t, 1, test1)

	test2 := TwiceCalibration("+3\n+3\n+4\n-2\n-4")
	assertEquals(t, 10, test2)

	test3 := TwiceCalibration("-6\n+3\n+8\n+5\n-6")
	assertEquals(t, 5, test3)

	test4 := TwiceCalibration("+7\n+7\n-2\n-7\n-4")
	assertEquals(t, 14, test4)
}
