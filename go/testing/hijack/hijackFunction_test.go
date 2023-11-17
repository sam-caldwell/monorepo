package hijack

import (
	"fmt"
	ansi2 "github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"reflect"
	"testing"
)

func TestHijackFunction(t *testing.T) {
	t.Skip("Test disabled until we can spend more time on this.")
	/*
	 * The test plan...
	 *
	 * 0. Run tests for dependencies just to be sure we aren't chasing our tail.
	 *
	 * 1. Set up the test by creating an initial state and altered state.
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
	)

	/*
	 * actualValue is in the outer test function scope, but it is visible
	 * to the targetFunc and imposterFunc, so we can alter the actualValue
	 * state and test its value as evidence of our targetFunc() and imposterFunc()
	 * operations, as we will see.
	 */
	var actualValue = words.EmptyString
	targetFunc := func() {
		actualValue = initialValue
		fmt.Println("---targetFunc---")
	}
	imposterFunc := func() {
		actualValue = alteredValue
		fmt.Println("---imposterFunc---")
	}
	/*
	 * First, let's check and make sure our actual value is what we expect.
	 * Doveryay no proveryay.
	 */
	if actualValue != words.EmptyString {
		ansi2.Red()
		defer ansi2.Reset()
		t.Fatal("expected actualValue to be empty")
	}
	/*
	 * Now run our targetFunc().  We expect it to change actualValue to its
	 * initialValue because we have not hijacked it yet.  See? It's happy
	 * and healthy...
	 */
	if targetFunc(); actualValue != initialValue {
		ansi2.Red()
		defer ansi2.Reset()
		t.Fatal("targetFunc did not maintain initial value")
	}
	/*
	 * Now we run our imposterFunc().  We expect that it will alter our
	 * actualValue to an alteredValue because it's a deviant little critter
	 * that does these things.  Quick, call Greg Abbott!  It's an ALTERED
	 * value...
	 */
	if imposterFunc(); actualValue != alteredValue {
		ansi2.Red()
		defer ansi2.Reset()
		t.Fatal("imposterFunc did not alter the value")
	}
	/*
	 * We have made it this far and our test is running as expected.
	 * Let's reset actualValue to EmptyString.  Greg "Gimmie Money" Abbott
	 * can be happy.  The Karens can relax, we're back to empty state.
	 */
	actualValue = words.EmptyString
	ansi2.Blue().Println("setup").Reset()
	/*
	 * As the Beastie Boys would say...  It's time to get Func'y
	 * Hijack this thingy.  It's time for some Sabotage!
	 */
	targetFuncPtr := reflect.ValueOf(targetFunc).Pointer()
	if targetFuncPtr != reflect.ValueOf(targetFunc).Pointer() {
		t.Fatalf("targetFuncPtr is not referncing the function")
	}
	ansi2.Yellow().Printf("targetFuncPtr:  %00x", targetFuncPtr).LF().Reset()

	imposterFuncPtr := reflect.ValueOf(imposterFunc).Pointer()
	if imposterFuncPtr != reflect.ValueOf(imposterFunc).Pointer() {
		t.Fatalf("imposterFuncPtr is not referncing the function")
	}

	ansi2.Yellow().Printf("imposterFuncPtr:%00x", imposterFuncPtr).LF().Reset()

	ansi2.Blue().Println("ready to hijack").
		Yellow().Printf("> targetFuncPtr:  %00x", targetFuncPtr).LF().
		Yellow().Printf("> imposterFuncPtr:%00x", imposterFuncPtr).LF().Reset()

	//originalMemory, err := AssemblyJmpToFunction(targetFuncPtr), error(nil)

	//originalMemory, err := hijackFunction(targetFuncPtr, imposterFuncPtr)

	ansi2.Blue().Println("hijack completed").
		Yellow().Printf("> targetFuncPtr:  %00x", targetFuncPtr).LF().
		Yellow().Printf("> imposterFuncPtr:%00x", imposterFuncPtr).LF().Reset()
	/*
	 * We should do some quick sanity checks (it's good for the code but too late for me).
	 */
	//if err != nil {
	//	ansi.Red()
	//	defer ansi.Reset()
	//	t.Fatalf("hijackFunction returned an error: %v", err)
	//}
	ansi2.Blue().Println("hijack complete with no errors").Reset()
	/*
	 * Make sure we didn't get too greedy.  Our hijack operation should only copy out
	 * a specific memory area (the jump code) so we are fast, light and sweet.
	 */
	//if len(originalMemory) != len(AssemblyJmpToFunction(imposterFuncPtr)) {
	//	ansi.Red()
	//	defer ansi.Reset()
	//	t.Fatalf("we expected the original memory to only be the size of the jump code")
	//}
	ansi2.Blue().Println("memory saved is the right size").Reset()
	/*
	 * Let's check and make sure our actual value is what we expect.
	 * Doveryay no proveryay.  Someone may drink and code some day and forget that
	 * we need the initial state to be correct.
	 */
	if actualValue != words.EmptyString {
		ansi2.Red()
		defer ansi2.Reset()
		t.Fatal("expected actualValue to be empty")
	}
	ansi2.Blue().Println("initial state verified").Reset()
	/*
	 * Run our targetFunc() and expect that it now alters the views.  Call the Karens!
	 * Someone get Greg Abbott on the phone and summon the Texas Legislature into Special Session,
	 * the pure targetFunc has been hijacked near the Texas Border, and it's now spreading its
	 * alteredValue.
	 */
	if targetFunc(); actualValue != alteredValue {
		ansi2.Red()
		defer ansi2.Reset()
		t.Fatalf("targetFunc should have altered the value. ActualValue:%s", actualValue)

	}
	ansi2.Green().Println("target is altered").Reset()
	/*
	 * And our imposterFunc() is still doing its thing.
	 */
	if imposterFunc(); actualValue != alteredValue {
		ansi2.Red()
		defer ansi2.Reset()
		t.Fatalf("imposterFunc should have altered the value. ActualValue:%s", actualValue)
	}
	/*
	 * I'd apologize for the political humor in this function if I cared enough.
	 *
	 * Remember: It's only "cancel culture" if it's the conservatives getting
	 *           cancelled.
	 *
	 * Slava Ukrani, folks!
	 */
}
