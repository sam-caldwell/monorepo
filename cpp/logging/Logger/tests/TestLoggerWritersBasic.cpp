/**
 * @name Logger/tests/TestLoggerWritersBasic.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */


#include <iostream>
#include "projects/logging/Logger/src/writer.h"

class TestLoggerWritersBasic : public TestBase {
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestLoggerWritersBasic(string n) {
        name = n;
    }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestLoggerWritersBasic() {}

    /**
     * @name test_writers_id
     * @brief test that Writer::Id enum matches values in RFC5424 and binary size is 1 byte.
     *
     * @return bool
     */
    bool test_writers_id() {
        Logging::Writer::Id id;
        return expect((int) (Logging::Writer::Id::null) == 0, "Writer::Id->Expect Id null is 0") &&
               expect((int) (Logging::Writer::Id::stdout) == 1, "Writer::Id->Expect Id stdout is 1") &&
               expect((int) (Logging::Writer::Id::stderr) == 2, "Writer::Id->Expect Id stderr is 2") &&
               expect((int) (Logging::Writer::Id::file) == 3, "Writer::Id->Expect Id file is 3") &&
               expect((int) (Logging::Writer::Id::syslog) == 4, "Writer::Id->Expect Id syslog is 4") &&
               expect((int) (Logging::Writer::Id::https) == 5, "Writer::Id->Expect Id https is 5") &
               expect(sizeof(id) == 1, "Writer::Id->Expect Id is 1 byte in size.  Actual:" + to_string(sizeof(id)));
    }

    /**
     * @name test_writers_strings_table
     * @brief Test reverse look up of string values for Writer::Id
     *
     * @return bool
     */
    bool test_writers_strings_table() {
        using namespace Logging;
        using namespace Writer;
        return expect(Logging::Writer::StringsTable.at("null") == Logging::Writer::Id::null, "Expect null ok.") &&
               expect(Logging::Writer::StringsTable.at("stdout") == Logging::Writer::Id::stdout, "Expect stdout ok.") &&
               expect(Logging::Writer::StringsTable.at("stderr") == Logging::Writer::Id::stderr, "Expect stderr ok.") &&
               expect(Logging::Writer::StringsTable.at("file") == Logging::Writer::Id::file, "Expect file ok.") &&
               expect(Logging::Writer::StringsTable.at("syslog") == Logging::Writer::Id::syslog, "Expect syslog ok.") &&
               expect(Logging::Writer::StringsTable.at("https") == Logging::Writer::Id::https, "Expect https ok.");
    }

    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");
        return expect(test_writers_id(), "test_writers_id() ok") &&
               expect(test_writers_strings_table(), "test_writers_strings_table() ok");
    }

};