package counters

import "testing"

func TestByteCounter_EqualSha1(t *testing.T) {
	t.Skip("debugging")
	const (
		size    = 100
		content = "test content"
	)
	var lhs ByteCounter
	var rhs ByteCounter

	t.Run("create two equal counters", func(t *testing.T) {
		lhs, err := NewByteCounter(size)
		if err != nil {
			t.Fatal(err)
		}
		rhs, err := NewByteCounter(size)
		if err != nil {
			t.Fatal(err)
		}
		if err := lhs.SetBytes(0, []byte(content)); err != nil {
			t.Fatal(err)
		}
		if err := rhs.SetBytes(0, []byte(content)); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ensure the two counters return the same hash", func(t *testing.T) {
		// Ensure that EqualSha1 returns true for equal content
		if !lhs.EqualSha1(&rhs) {
			t.Fatalf("Expected EqualSha1 to return true, but got false")
		}
	})

	t.Run("ensure the two counters return the same hash and the operation is commutative", func(t *testing.T) {
		// Ensure that EqualSha1 returns true for equal content
		if !rhs.EqualSha1(&lhs) {
			t.Fatalf("Expected EqualSha1 to return true, but got false")
		}
	})

	t.Run("Modify RHS and confirm hash equality fails", func(t *testing.T) {
		// Modify the content of counter2
		if err := rhs.Add(1); err != nil {
			t.Fatalf("%v", err)
		}
		// Ensure that EqualSha1 returns false for different content
		if lhs.EqualSha1(&rhs) {
			t.Fatalf("Expected EqualSha1 to return false, but got true.\n"+
				"lhs: '%s' (%s)\n"+
				"rhs: '%s' (%s)\n",
				lhs.String(), lhs.Sha1(),
				rhs.String(), rhs.Sha1())
		}
	})

	t.Run("ensure equality fails and commutative property persists", func(t *testing.T) {
		// Ensure that EqualSha1 returns true for equal content
		if rhs.EqualSha1(&lhs) {
			t.Fatalf("Expected EqualSha1 to return true, but got false")
		}
	})

	t.Run("Modify LHS in the same way and confirm hash equality is restored", func(t *testing.T) {
		// Modify the content of counter2
		if err := lhs.Add(1); err != nil {
			t.Fatal(err)
		}
		// Ensure that EqualSha1 returns false for different content
		if lhs.EqualSha1(&rhs) {
			t.Fatalf("Expected EqualSha1 to return false, but got true (lhs:rhs)")
		}
		if rhs.EqualSha1(&lhs) {
			t.Fatalf("Expected EqualSha1 to return false, but got true (rhs:lhs)")
		}
	})

}
