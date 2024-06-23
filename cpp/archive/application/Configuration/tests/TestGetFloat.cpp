/**
 * @name Configuration/tests/TestGetFloat.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../../../../types/src/ConfigStateMap/ConfigStateMap.h"
#include "../src/Configuration.h"

class TestGetFloat : public TestBase {
private:
    Configuration *cfg;
    ConfigStateMap *map;
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestGetFloat(string n) {
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
    ~TestGetFloat() {
        if (cfg) delete cfg;
    }

    /**
     * @name test_get_float
     * @brief expect that given a known ConfigStateMap record the getFloat() method will return the corresponding
     *        ValueTypes value.
     *
     * @return bool (pass/fail)
     */
    bool test_get_float() {
        return expect(cfg->getFloat("internal4") == (float)(3.14), "Expect .getFloat(internal4) will return 3.14.");
    }

    /**
     * @name test_get_float_missing
     * @brief expect that an out_of_range exception will be thrown if a non-existent key name is passed to getFloat()
     *
     * @return bool (pass/fail)
     */
    bool test_get_float_missing() {
        try {
            return expect(cfg->getFloat("BadToken") == 3.14,
                          "Expect getFloat() exception on non-existing token. Oops");
        } catch (out_of_range &e) {
            return expect(true, "Expect getFloat() exception on non-existing token. ok");
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
               expect(test_get_float(), "Expect test_get_float() OK") &&
               expect(test_get_float_missing(), "Expect test_get_float_missing() OK");
    }
};/*end of class*/