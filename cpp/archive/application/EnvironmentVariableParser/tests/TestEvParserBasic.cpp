/**
 * @name EnvironmentVariableParser/tests/TestEvParserBasic.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../../Configuration/src/Configuration.h"
#include "../src/EnvironmentVariableParser.h"

class TestEvParserBasic : public TestBase {
public:
    TestEvParserBasic(string n) {
        name = n;
    }

    ~TestEvParserBasic() {/*noop*/}

    bool test_map_ok() {
        debug(name + "::test_map_ok()");
        ConfigStateMap map;
        return expect(
                map.add("HOME", "homedir", String, false, false) &&
                map.add("USER", "username", String, false, false) &&
                map.add("DELETE_ME", "to.delete", String, false, false), "Expect we can add three records"
        ) &&
               expect(map.count() == 3, "expect count is 3") &&
               expect(map.erase("DELETE_ME"), "expect we can delete one record") &&
               expect(map.count() == 2, "expect count is 2");
    }

    bool test_parser_with_map_by_reference() {
        debug(name + "::test_parser_with_map_by_reference()");
        Configuration cfg;
        ConfigStateMap map;
        if (!expect(
                map.add("HOME", "homedir", String) &&
                map.add("USER", "username", String) &&
                map.add("DELETE_ME", "delete.me", String), "Expect we can add three records"))
            return false;
        EnvironmentVariableParser parser(&cfg, &map);
        return true;
    }

    bool test_parser_with_map_by_value() {
        debug(name + "::test_parser_with_map_by_value()");
        Configuration cfg;
        ConfigStateMap map;
        if (!expect(
                map.add("HOME", "homedir", String) &&
                map.add("USER", "username", String) &&
                map.add("DELETE_ME", "delete.me", String), "Expect we can add three records"))
            return false;
        EnvironmentVariableParser parser(&cfg, &map);
        return true;
    }

    bool main() {
        debug(name + "::main()");
        return expect(test_map_ok(), "Expect test_map_ok OK") &&
               expect(test_parser_with_map_by_reference(), "expect test_parser_with_map_by_reference OK") &&
               expect(test_parser_with_map_by_value(), "expect test_parser_with_map_by_value OK");
    }

};/*end of class*/