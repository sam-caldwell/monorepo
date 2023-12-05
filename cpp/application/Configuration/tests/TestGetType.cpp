/**
 * @name Configuration/tests/TestGetType.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */


#include <iostream>
#include "projects/common/types/ConfigStateMap.h"
#include "projects/application/Configuration/src/Configuration.h"

class TestGetType : public TestBase {
private:
    ConfigStateMap *map;
    Configuration *cfg;
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestGetType(string n) {
        name = n;

        map = new ConfigStateMap;
        map->add("PropertyA", "internal1", Bool);
        map->add("PropertyB", "internal2", Bool);
        map->add("PropertyC", "internal3", Double);
        map->add("PropertyD", "internal4", Float);
        map->add("PropertyE", "internal5", Int);
        map->add("PropertyF", "internal6", String);
        map->add("PropertyG", "internal7", Uint);

        cfg = new Configuration;
        cfg->loadData(map, "PropertyA", "true");
        cfg->loadData(map, "PropertyB", "false");
        cfg->loadData(map, "PropertyC", "1.0");
        cfg->loadData(map, "PropertyD", "3.14");
        cfg->loadData(map, "PropertyE", "-10");
        cfg->loadData(map, "PropertyF", "hello");
        cfg->loadData(map, "PropertyG", "25");
    }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestGetType() {
        if (cfg) delete cfg;
    }

    /**
     * @name test_get_type
     * @brief expect that given a known ConfigStateMap record the getType() method will return the corresponding
     *        ValueTypes value.
     *
     * @return bool (pass/fail)
     */
    bool test_get_type() {
        return expect(cfg->getType("internal1") == Bool, "Expect .getType(internal1) will return Bool. Got:"+ ValueTypeToString(cfg->getType("internal1")));
               expect(cfg->getType("internal2") == Bool, "Expect .getType(internal2) will return Bool. Got:"+ValueTypeToString(cfg->getType("internal2")));
               expect(cfg->getType("internal3") == Double, "Expect .getType(internal3) will return Double. Got:"+ValueTypeToString(cfg->getType("internal3")));
               expect(cfg->getType("internal4") == Float, "Expect .getType(internal4) will return Float. Got:"+ValueTypeToString(cfg->getType("internal4")));
               expect(cfg->getType("internal5") == Int, "Expect .getType(internal5) will return Int. Got:"+ValueTypeToString(cfg->getType("internal5")));
               expect(cfg->getType("internal6") == String, "Expect .getType(internal6) will return String. Got:"+ValueTypeToString(cfg->getType("internal6")));
               expect(cfg->getType("internal7") == Uint, "Expect .getType(internal7) will return Uint. Got:"+ValueTypeToString(cfg->getType("internal7")));
    }

    /**
     * @name test_get_type_missing
     * @brief expect that an out_of_range exception will be thrown if a non-existent key name is passed to getType()
     *
     * @return bool (pass/fail)
     */
    bool test_get_type_missing() {
        try {
            return expect(cfg->getType("BadToken") == Bool, "Expect getType() exception on non-existing token. Oops");
        } catch (out_of_range &e) {
            return expect(true, "Expect getType() exception on non-existing token. ok");
        }
    }

    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");
        return expect(map, "Expect map is initialized.") &&
               expect(cfg, "Expect cfg is initialized.") &&
               expect(cfg->count() == 7, "expect record count 7 in Configuration. Got:" + to_string(cfg->count())) &&
               expect(map->count() == 7, "expect record count 7 in ConfigStateMap. Got:" + to_string(map->count())) &&
               expect(test_get_type(), "Expect test_get_type() OK") &&
               expect(test_get_type_missing(), "Expect test_get_type_missing() OK");
    }

};/*end of class*/