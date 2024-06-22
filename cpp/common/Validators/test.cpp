/**
 * @name Validators/test.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../networking/TcpCommon/src/PipeResponse/main.h"


//

/**
 * Base tests
 */
#include "tests/TestValidateDirectoryTraversal.cpp"
#include "tests/TestValidateTags.cpp"
#include "tests/TestValidateDestination.cpp"


using namespace std;

int main(int argc, char *argv[]) {
    Tester t(argc, argv);
    t.enable_debug();

    /**
     * Base tests
     */
    t.add(new TestValidateDirectoryTraversal("TestValidateDirectoryTraversal"), RUN_TEST);
    t.add(new TestValidateTags("TestValidateTags"), RUN_TEST);
    t.add(new TestValidateDestination("TestValidateDestination"), RUN_TEST);

    /**
     *  END OF TESTS
     */
    return t.run();
}
