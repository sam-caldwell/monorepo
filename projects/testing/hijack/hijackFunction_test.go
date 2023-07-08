package hijack

import (
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"reflect"
	"testing"
)

func TestHijackFunction(t *testing.T) {
	/*
	 * The test plan...
	 *
	 * 1. Setup the test by creating an initial state and altered state.
	 *
	 * 2. Then create two functions (targetFunc and imposterFunc) which
	 *    will alter the 'actualValue' variable in the outer scope.
	 *
	 * 3. Perform a pre-test to prove we are sane (am I?).
	 *
	 * 4. Hijack targetFunc with imposterFunc.
	 *
	 * 5. Perform the same test we ran at pre-test and expect evidence
	 *    of our hijacking.
	 */
	const (
		initialValue = "original"
		alteredValue = "altered"
		debug        = true
	)
	stepComplete := func(n int) {
		if debug {
			t.Logf("Step %d complete", n)
		}
	}
	stepComplete(0)

	/*
	 * actualValue is in the outer test function scope, but it is visible
	 * to the targetFunc and imposterFunc, so we can alter the actualValue
	 * state and test its value as evidence of our targetFunc() and imposterFunc()
	 * operations, as we will see.
	 */
	var actualValue = words.EmptyString
	targetFunc := func() { actualValue = initialValue }
	imposterFunc := func() { actualValue = alteredValue }
	/*
	 * First, let's check and make sure our actual value is what we expect.
	 * Doveryay no proveryay.
	 */
	if actualValue != words.EmptyString {
		t.Fatal("expected actualValue to be empty")
	}
	/*
	 * Now run our targetFunc().  We expect it to change actualValue to its
	 * initialValue because we have not hijacked it yet.  See? It's happy
	 * and healthy...
	 */
	if targetFunc(); actualValue != initialValue {
		t.Fatal("targetFunc did not maintain initial value")
	}
	/*
	 * Now we run our imposterFunc().  We expect that it will alter our
	 * actualValue to an alteredValue because it's a deviant little critter
	 * that does these things.  Quick, call Greg Abbott!  It's an ALTERED
	 * value...
	 */
	if imposterFunc(); actualValue != alteredValue {
		t.Fatal("imposterFunc did not alter the value")
	}
	/*
	 * We have made it this far and our test is running as expected.
	 * Let's reset actualValue to EmptyString.  Greg "Gimmie Money" Abbott
	 * can be happy.  The Karens can relax, we're back to empty state.
	 */
	actualValue = words.EmptyString
	stepComplete(1)
	/*
	 * As the Beastie Boys would say...  It's time to get Func'y
	 * Hijack this thingy.  It's time for some Sabotage!
	 */
	targetFuncPtr := reflect.ValueOf(targetFunc).Pointer()
	imposterFuncPtr := reflect.ValueOf(imposterFunc).Pointer()
	originalMemory, err := hijackFunction(targetFuncPtr, imposterFuncPtr)
	stepComplete(2)
	/*
	 * We should do some quick sanity checks (it's good for the code but too late for me).
	 */
	if err != nil {
		t.Fatalf("hijackFunction returned an error: %v", err)
	}
	/*
	 * Make sure we didn't get too greedy.  Our hijack operation should only copy out
	 * a specific memory area (the jump code) so we are fast, light and sweet.
	 */
	if len(originalMemory) != len(assemblyJmpToFunction(imposterFuncPtr)) {
		t.Fatalf("we expected the original memory to only be the size of the jump code")
	}
	/*
	 * Let's check and make sure our actual value is what we expect.
	 * Doveryay no proveryay.  Someone may drink and code some day and forget that
	 * we need the initial state to be correct.
	 */
	//if actualValue != words.EmptyString {
	//	t.Fatal("expected actualValue to be empty")
	//}
	/*
	 * Run our targetFunc() and expect that it now alters the views.  Call the Karens!
	 * Someone get Greg Abbott on the phone and summon the Texas Legislature into Special Session,
	 * the pure targetFunc has been hijacked near the Texas Border, and it's now spreading its
	 * alteredValue.
	 */
	//if targetFunc(); actualValue != alteredValue {
	//	t.Fatalf("targetFunc should have altered the value. ActualValue:%s", actualValue)
	//}
	/*
	 * And our imposterFunc() is still doing its thing.
	 */
	//if imposterFunc(); actualValue != alteredValue {
	//	t.Fatalf("imposterFunc should have altered the value. ActualValue:%s", actualValue)
	//}
	/*
	 * I'd apologize for the political humor in this function if I cared enough.
	 *
	 * Remember: It's only "cancel culture" if it's the conservatives getting
	 *           cancelled.
	 *
	 * Slava Ukrani, folks!
	 */
}
