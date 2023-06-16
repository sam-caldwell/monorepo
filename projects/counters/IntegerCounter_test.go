package counters

import "testing"

func TestIntegerCounter(t *testing.T) {

	func() {
		count := IntegerCounter(0, 1)

		test := func(n int) {
			if count() != n {
				t.Fatalf("Failed on %d", n)
			}
		}

		for i := 1; i <= 100; i++ {
			test(i)
		}
	}()
	//
	//func() {
	//	count := IntegerCounter(10, -1)
	//
	//	test := func(n int) {
	//		if count() != n {
	//			t.Fatalf("Failed on %d", n)
	//		}
	//	}
	//
	//	for i := 10; i > 0; i-- {
	//		test(i)
	//	}
	//}()
}
