/**
 * @name Channel/test.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */

//This will turn on debug logging and other related tooling...
#define CHANNEL_DEBUGGING
#include "Channel.h"
#include <iostream>
#include <thread>
#include <chrono>
#include <cassert>

// Function to be used as a producer thread
void producer(Channel<int> &ch) {
    for (int i = 0; i < 5; ++i) {
        ch << i;
        std::this_thread::sleep_for(std::chrono::milliseconds(100));
    }
    ch.close();
}

// Function to be used as a consumer thread
void consumer(Channel<int> &ch, std::vector<int> &result) {
    int value;
    while (ch >> value) {
        result.push_back(value);
        std::this_thread::sleep_for(std::chrono::milliseconds(50));
    }
}

int main() {
    Channel<int> ch(3); // Bounded channel with size limit 3
    std::vector<int> result;

    // Create threads for producer and consumer
    std::thread producerThread(producer, std::ref(ch));
    std::thread consumerThread(consumer, std::ref(ch), std::ref(result));

    // Wait for threads to finish
    producerThread.join();
    consumerThread.join();

    // Check if the consumer received all expected values
    std::vector<int> expected = {0, 1, 2, 3, 4};
    assert(result == expected);

    // Test closing and reading from a closed channel
    Channel<int> closedCh;
    closedCh.close();
    int value;
    bool readSuccessful = closedCh >> value;
    assert(!readSuccessful);

    std::cout << "All tests passed successfully!" << std::endl;

    return 0;
}
