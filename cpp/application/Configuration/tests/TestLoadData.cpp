/**
 * @name Configuration/tests/TestLoadData.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#define PUBLIC_FOR_TESTING_ONLY() public:
#define UNDO_PUBLIC_FOR_TESTING_ONLY() protected:

#include "projects/common/types/ConfigStateMap.h"
#include "projects/application/Configuration/src/Configuration.h"





class TestLoadData : public TestBase {
private:
    Configuration *cfg;
    ConfigStateMap *map;
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestLoadData(string n) {
        name = n;
        cfg = new Configuration;
        map = new ConfigStateMap;
        if (map->count() != 0) throw runtime_error("Expect map node count 0");
        map->add("PropertyA", "internal1", Bool);
        map->add("PropertyB", "internal2", Bool);
        map->add("PropertyC", "internal3", Double);
        map->add("PropertyD", "internal4", Float);
        map->add("PropertyE", "internal5", Int);
        map->add("PropertyF", "internal6", String);
        map->add("PropertyG", "internal7", Uint);
    }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestLoadData() {
        if (cfg) delete cfg;
    }

    /**
     * @name test_loadData
     * @brief expect that we can call Configuration::loadData() and load a data record into the tree without error.
     *
     * @param property string
     * @param value value
     * @return bool
     */
    bool test_loadData(string property, string value) {
        return cfg->loadData(map, property, value);
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(map, "Expect map is initialized.") &&
               expect(cfg, "Expect cfg is initialized.") &&
               expect(cfg->count() == 0, "expect record count 0 in Configuration. Got:" + to_string(cfg->count())) &&
               expect(map->count() == 7,
                      "expect record count 7 in ConfigStateMap. Got:" + to_string(map->count())) &&
               expect(test_loadData("PropertyA", "true"), "expect loadData() bool ok") &&
               expect(test_loadData("PropertyB", "false"), "expect loadData() bool ok") &&
               expect(test_loadData("PropertyC", "1.0"), "expect loadData() double ok") &&
               expect(test_loadData("PropertyD", "3.14"), "expect loadData() float ok") &&
               expect(test_loadData("PropertyE", "-10"), "expect loadData() int ok") &&
               expect(test_loadData("PropertyF", "hello"), "expect loadData() string ok") &&
               expect(test_loadData("PropertyG", "25"), "expect loadData() uint ok");
    }
};
