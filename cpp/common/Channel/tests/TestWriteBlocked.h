/**
 * @name Channel/tests/TestWriteBlocked.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include <thread>
#include "../src/Channel.h"

using namespace std;

class TestWriteBlocked : public TestBase {
    bool status;
    bool write_complete;
    Channel<int> *data;
private:
public:
    TestWriteBlocked(string n) {
        name = n;
        data = new Channel<int>(4);
    }

    ~TestWriteBlocked() {
        delete data;
    }

    bool test_setup() {
        debug(name + "::test_setup() starting...");

        this->status = true;
        this->write_complete = false;
        debug(name + "::test_setup() starting thread...");
        thread sender([this]() {
            this->debug(name + "::test_setup() sender thread started.");
            this->data->push(1);
            this->debug(name + "::test_setup() 1 pushed");
            this->data->push(2);
            this->debug(name + "::test_setup() 2 pushed");
            this->data->push(3);
            this->debug(name + "::test_setup() 3 pushed");
            this->data->push(4);
            this->debug(name + "::test_setup() 4 pushed");
            this->debug(name + "::test_setup() pushing the blocking message.");
            this->data->push(5);//Blocking push.
            this->debug(name + "::test_setup() 5 pushed");
            this->write_complete = true; //ok.
        });
        sender.detach();
        this->debug(name + "::test_setup() completed");
        return !this->write_complete; //Expect false.
    }

    bool test_reads() {
        debug(name + "::test_setup() starting...");
        this->data->pop(); //free the block.
        this_thread::sleep_for(100ms); //Ensure sender has time to respond.
        return write_complete &&
               expect(this->data->size() == 4, "expect size 4 after pop and block freed.") &&
               expect(this->data->pop() == 2, "Expect next node pops as expected") &&
               expect(this->data->size() == 3, "expect size 3 (no new nodes were pushed");
    }


    bool main() {
        debug(name + "::main() starting...");
        return expect(test_setup(), "Expect we can setup channel with fix limit and go to block state.") &&
               expect(test_reads(), "Expect we can read, free the block and process as expected.");
    }
};