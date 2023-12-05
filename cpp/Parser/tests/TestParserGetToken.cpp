/**
 * @name Parser/tests/TestParserGetToken.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/Tester/TestBase/main.h"
#include "projects/common/Parser/src/Parser.h"
#include "projects/common/exceptions/exceptions.h"

class TestParserGetToken : public TestBase {
private:
    const string GoodYamlFile = "data_good.yaml";
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestParserGetToken(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestParserGetToken() {}

    bool test() {
        debug(name + "::test()");
        Parser p(GoodYamlFile);
        return expect(p.getType() == Header, "00. Expect header as first expected token type") &&
               expect(p.fetchNextLine(), "01. Expect fetchNextLine() ok") && // line 4
               expect(p.fetchNextLine(), "02. Expect fetchNextLine() ok") && // line 5
               expect(p.getLine() == 5, "03. Expect line position verified (line 5)") &&
               expect(p.routeKeyValueTag(), "04. Expect routeKeyValueTag() ok") &&
               expect(p.getToken("myKey").compare("value") == 0, "05. Expect getToken('myKey') returns 'value'") &&
               expect(p.fetchNextLine(), "06. Expect fetchNextLine() ok") && // line 6;
               expect(p.getLine() == 6, "07. Expect line position verified (line 6)");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test(), "Expect test() ok");
    }
};
