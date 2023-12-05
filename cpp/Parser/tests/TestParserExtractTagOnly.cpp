/**
 * @name Parser/tests/TestParserExtractTagOnly.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/Tester/TestBase/main.h"
#include "projects/common/formatter/formatter.h"
#include "projects/common/exceptions/exceptions.h"
#include "projects/common/Parser/src/extractTagOnly.cpp"


class TestParserExtractTagOnly : public TestBase {
private:
    const string dataFile = "data_extract_only.yaml";

public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestParserExtractTagOnly(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestParserExtractTagOnly() {}

    bool test_extract_tag(string line, string result) {
        return expect(extractTagOnly(&line).compare(result) == 0,
                      string("extractTagOnly() line'") +
                      line + "' " +
                      string("result:'") + result + string("'"));
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        string ln1 = "tag1:";
        string r1 = "tag1";
        string ln2 = "tag2: value";
        string r2 = "tag2";
        string ln3 = " - tag3";
        string r3 = "tag3";
        return expect(test_extract_tag(ln1, r1), "test_extract_tag() ok(1)") &&
               expect(test_extract_tag(ln2, r2), "test_extract_tag() ok(2)") &&
               expect(test_extract_tag(ln3, r3), "test_extract_tag() ok(3)");
    }
};
