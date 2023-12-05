/**
 * @name types/tests/TestMapper.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/common/types/Mapper.h"

class TestMapper : public TestBase {
public:
    TestMapper(string n) {
        name = n;
    }

    ~TestMapper() {/*noop*/}

    bool test_initial_state() {
        Mapper<string, int> m;
        return true;
    }

    bool test_push() {
        Mapper<string, int> m;
        return expect(m.push("A", 1), "push() ok");
    }

    bool test_count() {
        Mapper<string, int> m;
        return expect(m.count() == 0, "count() ok") &&
               expect(m.push("A", 1), "push() ok") &&
               expect(m.count() == 1, "count() ok") &&
               expect(m.push("B", 2), "push() ok") &&
               expect(m.count() == 2, "count() ok");
    }

    bool test_currentKey() {

        Mapper<string, int> m;
        return expect(m.count() == 0, "count() ok") &&
               expect(m.push("A", 1), "push() ok") &&
               expect(m.count() == 1, "count() ok") &&
               expect(m.currentKey() == "A", "expect current A") &&
               expect(m.push("B", 2), "push() ok") &&
               expect(m.currentKey() == "B", "expect current B") &&
               expect(m.count() == 2, "count() ok");
    }

    bool test_current() {

        Mapper<string, int> m;
        return expect(m.count() == 0, "count() ok") &&
               expect(m.push("A", 1), "push() ok") &&
               expect(m.count() == 1, "count() ok") &&
               expect(m.current() == 1, "expect current A") &&
               expect(m.push("B", 2), "push() ok") &&
               expect(m.current() == 2, "expect current B") &&
               expect(m.count() == 2, "count() ok");
    }

    bool test_pop() {
        Mapper<string, int> m;
        return expect(m.count() == 0, "count() ok") &&
               expect(m.push("A", 1), "push() ok") &&
               expect(m.count() == 1, "count() ok") &&
               expect(m.push("B", 1), "push() ok") &&
               expect(m.count() == 2, "count() ok") &&
               expect(m.pop("A"), "pop() ok") &&
               expect(m.count() == 1, "count() ok");
    }

    bool test_empty() {
        Mapper<string, int> m;
        return expect(m.count() == 0, "count() ok") &&
               expect(m.empty(), "empty() ok") &&
               expect(m.push("A", 1), "push() ok") &&
               expect(m.count() == 1, "count() ok") &&
               expect(m.push("B", 2), "push() ok") &&
               expect(m.count() == 2, "count() ok") &&
               expect(!m.empty(), "empty() ok");
    }

    bool test_get() {
        Mapper<string, int> m;
        return expect(m.count() == 0, "count() ok") &&
               expect(m.empty(), "empty() ok") &&
               expect(m.push("A", 1), "push() ok") &&
               expect(m.count() == 1, "count() ok") &&
               expect(m.push("B", 2), "push() ok") &&
               expect(m.count() == 2, "count() ok") &&
               expect(!m.empty(), "empty() ok") &&
               expect(m.get("A") == 1, "get() ok") &&
               expect(m.get("B") == 2, "get() ok");
    }

    bool test_get_error() {
        Mapper<string, int> m;
        expect(m.count() == 0, "count() ok") &&
        expect(m.empty(), "empty() ok") &&
        expect(m.push("A", 1), "push() ok");
        try {
            expect(!m.get("foo"), "empty() failed to throw exception");
            return false;
        } catch (out_of_range) {
            return expect(true, "Expect exception.  ok");
        }

    }

    bool test_push_duplicate() {
        Mapper<string, int> m;
        return expect(m.push("A", 1), "push() ok") &&
               expect(m.push("A", 2), "push() ok (update)");
    }

    bool main() {
        debug(name + "::main()");
        return expect(test_initial_state(), "Expect test_initial_state() OK") &&
               expect(test_push(), "Expect test_push() ok") &&
               expect(test_count(), "Expect test_count() ok") &&
               expect(test_currentKey(), "Expect test_currentKey() ok") &&
               expect(test_current(), "Expect test_current() ok") &&
               expect(test_pop(), "Expect test_pop() ok") &&
               expect(test_empty(), "Expect test_empty() ok") &&
               expect(test_get_error(), "Expect test_get_error() ok") &&
               expect(test_push_duplicate(), "Expect test_push_duplicate() ok");
    }

};/*end of class*/