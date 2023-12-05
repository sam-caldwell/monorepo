/**
 * @name exceptions/test.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/Tester/Tester/main.h"
/**
 * Base tests
 */
#include "tests/TestBoundsCheckError.cpp"
#include "tests/TestClosedChannelException.cpp"
#include "tests/TestInvalidKeyString.cpp"
#include "tests/TestLockTimeoutExpired.cpp"
#include "tests/TestNotFound.cpp"
#include "tests/TestNullValue.cpp"
#include "tests/TestReadOnlyKeyValue.cpp"
#include "tests/TestTypeError.cpp"
#include "tests/TestUnexpectedType.cpp"
#include "tests/TestValueAssignmentError.cpp"
#include "tests/TestValueError.cpp"
#include "tests/TestValueGetStringError.cpp"
#include "tests/TestMissingInputs.cpp"
#include "tests/TestFileIoError.cpp"


using namespace std;

int main(int argc, char *argv[]) {
    cout << "test exceptions/main.cpp:main() starting" << endl;
    Tester t(argc, argv);
    cout << "test exceptions/main.cpp:main() tester loaded" << endl;
    t.enable_debug();
    /**
     * Base tests
     */
    t.add(new TestBoundsCheckError("TestBoundsCheckError"), RUN_TEST);
    t.add(new TestClosedChannelException("TestClosedChannelException"), RUN_TEST);
    t.add(new TestInvalidKeyString("TestInvalidKeyString"), RUN_TEST);
    t.add(new TestLockTimeoutExpired("TestLockTimeoutExpired"), RUN_TEST);
    t.add(new TestNotFound("TestNotFound"), RUN_TEST);
    t.add(new TestNullValue("TestNullValue"), RUN_TEST);
    t.add(new TestReadOnlyKeyValue("TestReadOnlyKeyValue"), RUN_TEST);
    t.add(new TestTypeError("TestTypeError"), RUN_TEST);
    t.add(new TestUnexpectedType("TestUnexpectedType"), RUN_TEST);
    t.add(new TestValueAssignmentError("TestValueAssignmentError"), RUN_TEST);
    t.add(new TestValueError("TestValueError"), RUN_TEST);
    t.add(new TestValueGetStringError("TestValueGetStringError"), RUN_TEST);
    t.add(new TestMissingInputs("TestMissingInputs"), RUN_TEST);
    t.add(new TestFileIoError("TestFileIoError"), RUN_TEST);

    /**
     *  END OF TESTS
     */
    return t.run();
}
