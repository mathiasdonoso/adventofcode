package day2

import "testing"

func TestDivisibleChecksum(t *testing.T) {
	data := "5\t9\t2\t8\n9\t4\t7\t3\n3\t8\t6\t5"
	test := DivisibleChecksum(data)
	assertEquals(t, 9, test)
}
