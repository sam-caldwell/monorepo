/**
 * @name Parsers/tests/TestParserTokenCount.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "../../../../Tester/Tester/main.h"
#include "../src/Parser.h"
#include "../../../exceptions/exceptions.h"

class TestParserTokenCount : public TestBase {
private:
    const string GoodYaml = "data_good.yaml";
    Parser *p;
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestParserTokenCount(string n) {
        name = n;
        p = new Parser(GoodYaml, Yaml);
    }

    bool test_parse() {
        p->parse();
        return true;
    }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestParserTokenCount() {
        if (p) delete p;
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(p->count() == 0, "Expect count()==0 ok.  count:" + to_string(p->count())) &&
               expect(test_parse(), "expect parse() ok.") &&
               expect(p->count() == 20, "Expect count()==20 ok.  count:" + to_string(p->count()));
    }
};
