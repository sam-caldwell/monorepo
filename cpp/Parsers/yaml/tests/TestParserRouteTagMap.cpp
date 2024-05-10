/**
 * @name Parsers/tests/TestParserRouteTagMap.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/Tester/TestBase/main.h"
#include "projects/common/Parser/src/Parser.h"
#include "projects/common/exceptions/exceptions.h"

class TestParserRouteTagMap : public TestBase {
private:
    const string GoodYamlFile = "data_good.yaml";
    const string BadYamlFile = "data_bad.yaml";
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestParserRouteTagMap(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestParserRouteTagMap() {}

    bool test() {
        Parser p(GoodYamlFile, Yaml);
        p.nextLine();

        return expect(p.getType() == Header, "Expect header as first expected token type") &&
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.routeHeader(), "Expect routeHeader() ok") &&
               expect(p.getType() == Any, "Expect the expected token will shift to Any after header is detected") &&

               //key-value: myKey
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&

               //secA
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.routeTagMap(), "Expect routeTagMap() ok") &&
               expect(p.getType() == Any, "Expect the expected token will shift to Any after header is detected") &&

               //secA...items
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&

               //secB
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.routeTagMap(), "Expect routeTagMap() ok") &&
               expect(p.getType() == Any, "Expect the expected token will shift to Any after header is detected") &&

               //secB...key-value items
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&

               //secC
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.routeTagMap(), "Expect routeTagMap() ok") &&
               expect(p.getType() == Any, "Expect the expected token will shift to Any after header is detected") &&

               //secC.obj1
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.routeTagMap(), "Expect routeTagMap() ok") &&
               expect(p.getType() == Any, "Expect the expected token will shift to Any after header is detected") &&

               //secC.obj1.propA
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.routeTagMap(), "Expect routeTagMap() ok") &&
               expect(p.getType() == Any, "Expect the expected token will shift to Any after header is detected") &&

               //secC.obj1.propB
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.routeTagMap(), "Expect routeTagMap() ok") &&
               expect(p.getType() == Any, "Expect the expected token will shift to Any after header is detected") &&

               //secC.obj1.propC
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.routeTagMap(), "Expect routeTagMap() ok") &&
               expect(p.getType() == Any, "Expect the expected token will shift to Any after header is detected") &&

               //secC.obj2.propD
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.routeTagMap(), "Expect routeTagMap() ok") &&
               expect(p.getType() == Any, "Expect the expected token will shift to Any after header is detected") &&

               //secC.obj2.propE
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.routeTagMap(), "Expect routeTagMap() ok") &&
               expect(p.getType() == Any, "Expect the expected token will shift to Any after header is detected") &&

               //secC.obj2.propF
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.routeTagMap(), "Expect routeTagMap() ok") &&
               expect(p.getType() == Any, "Expect the expected token will shift to Any after header is detected") &&

               //secD
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.routeTagMap(), "Expect routeTagMap() ok") &&
               expect(p.getType() == Any, "Expect the expected token will shift to Any after header is detected") &&

               //secD.objD1
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.routeTagMap(), "Expect routeTagMap() ok") &&
               expect(p.getType() == Any, "Expect the expected token will shift to Any after header is detected") &&

               //secD.objD11
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.routeTagMap(), "Expect routeTagMap() ok") &&
               expect(p.getType() == Any, "Expect the expected token will shift to Any after header is detected") &&

               //secD.objD111
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.routeTagMap(), "Expect routeTagMap() ok") &&
               expect(p.getType() == Any, "Expect the expected token will shift to Any after header is detected") &&

               //secD.objD111.propG
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.routeTagMap(), "Expect routeTagMap() ok") &&
               expect(p.getType() == Any, "Expect the expected token will shift to Any after header is detected") &&

               //secD.objD111.propH
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.routeTagMap(), "Expect routeTagMap() ok") &&
               expect(p.getType() == Any, "Expect the expected token will shift to Any after header is detected") &&

               //secD.objD112
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.routeTagMap(), "Expect routeTagMap() ok") &&
               expect(p.getType() == Any, "Expect the expected token will shift to Any after header is detected") &&

               //secD.objD112.I
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.routeTagMap(), "Expect routeTagMap() ok") &&
               expect(p.getType() == Any, "Expect the expected token will shift to Any after header is detected") &&

               //secD.objD112.J
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.routeTagMap(), "Expect routeTagMap() ok") &&
               expect(p.getType() == Any, "Expect the expected token will shift to Any after header is detected") &&

               //secE
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.routeTagMap(), "Expect routeTagMap() ok") &&
               expect(p.getType() == Any, "Expect the expected token will shift to Any after header is detected");
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