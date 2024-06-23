/**
 * @name Channel/tests/TestConcurrentUse.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include <thread>
#include "../src/Channel.h"

using namespace std;

class TestConcurrentUse : public TestBase {
    bool status;
    int sz;
    int records_sent;
    int records_recv;
    Channel<int> data;
private:
public:
    TestConcurrentUse(string n) { name = n; }

    ~TestConcurrentUse() {}

    bool test_setup() {
        debug(name + "::test_setup() starting...");
        this->status = true;
        for (sz = 1; sz < 65537; sz *= 2) {

            this->records_sent = 0;
            this->records_recv = 0;

            thread receiver([this]() {
                this_thread::sleep_for(100ms);
                for (int p = 1; p < this->sz; p++) {
                    int n = this->data.pop();
                    this->records_recv++;
                    //this->debug("recv p=" + to_string(p) +
                    //            " n=" + to_string(n) +
                    //            " sz=" + to_string(this->sz));
                    if (n == 0) {
                        this->debug("FATAL ERROR: zero value should not be in queue.  p=" + to_string(p));
                        this->status = false;
                        exit(1);
                    }
                }
            });

            //debug(name + "::test_concurrency() receiver running.");

            thread sender([this]() {
                for (int i = 1; i < this->sz; i++) {
                    if (i == 0) {//Make sure we don't send a zero.
                        this->debug("BAD INPUT.  Zero not allowed in test.");
                        this->status = false;
                        exit(1);
                    } else {
                        this->data.push(i);
                        this->records_sent++;
                        //this->debug("push " + to_string(i) + " to channel");
                    }
                }
            });

            //debug(name + "::test_concurrency() sender running.");
            //debug(name + "::test_concurrency() waiting on thread completion");
            sender.join();
            receiver.join();
            //debug(name + "::test_concurrency() threads are done");
            if (!expect(this->status, "Expect concurrency test will work with " + to_string(sz) + " elements"))
                return false;
        }
        debug(name + "::test_concurrency() all iterations are complete.");
        this->data.close();
        return status &&
               expect(this->data.closed(), "Expect the channel is closed now.") &&
               expect(records_sent == records_recv,
                      "Expect number records sent (" + to_string(records_sent) + ") matches records recd (" +
                      to_string(records_recv) + ")");
        return status;
    }

    bool main() {
        debug(name + "::main() starting...");
        return expect(test_setup(), "Expect two synchronous threads can share a channel and send/recv data.");
    }
};