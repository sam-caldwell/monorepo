/**
 * @name Logger/test.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/Tester/Tester/main.h"


//

/**
 * Base tests
 */
#include "tests/TestBasics.cpp"
#include "tests/TestLoggerWritersBasic.cpp"
#include "tests/TestLoggerSeverity.cpp"
#include "tests/TestLoggerContext.cpp"
#include "tests/TestLoggerAddContextWithDefaults.cpp"
#include "tests/TestLoggerAddContextWithValues.cpp"
#include "tests/TestLoggerLevel.cpp"
#include "tests/TestLoggerWriter.cpp"
#include "tests/TestLoggerNullWriter.cpp"
#include "tests/TestLoggerStdoutWriter.cpp"
#include "tests/TestLoggerStderrWriter.cpp"
#include "tests/TestLoggerFileWriter.cpp"
#include "tests/TestLoggerSyslogWriter.cpp"
#include "tests/TestLoggerHttpsWriter.cpp"


using namespace std;

int main(int argc, char *argv[]) {
    Tester t(argc, argv);
    t.enable_debug();

    /**
     * Base tests
     */
    t.add(new TestBasics("TestBasics"), RUN_TEST);
    t.add(new TestLoggerWritersBasic("TestLoggerWritersBasic"), RUN_TEST);
    t.add(new TestLoggerSeverity("TestLoggerSeverity"), RUN_TEST);
    t.add(new TestLoggerContext("TestLoggerContext"), RUN_TEST);

    t.add(new TestLoggerAddContextWithDefaults("TestLoggerAddContextWithDefaults"), RUN_TEST);
    t.add(new TestLoggerAddContextWithValues("TestLoggerAddContextWithValues"), RUN_TEST);
    t.add(new TestLoggerLevel("TestLoggerLevel"), RUN_TEST);

    t.add(new TestLoggerWriter("TestLoggerWriter"), RUN_TEST);

    t.add(new TestLoggerNullWriter("TestLoggerNullWriter"), RUN_TEST);
    t.add(new TestLoggerStdoutWriter("TestLoggerStdoutWriter"), RUN_TEST);
    t.add(new TestLoggerStderrWriter("TestLoggerStderrWriter"), RUN_TEST);
    t.add(new TestLoggerFileWriter("TestLoggerFileWriter"), RUN_TEST);

    t.add(new TestLoggerSyslogWriter("TestLoggerSyslogWriter"), RUN_TEST);
    t.add(new TestLoggerHttpsWriter("TestLoggerHttpsWriter"), RUN_TEST);

    /**
     *  END OF TESTS
     */
    return t.run();
}
