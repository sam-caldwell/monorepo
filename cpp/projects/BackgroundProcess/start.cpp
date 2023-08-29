/*
 * BackgroundProcess Class
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */
#include "BackgroundProcess.h"
void BackgroundProcess::start() {
    if (!running) {
        running = true;

        auto t = [&]() {
            taskFunc();
            running = false;
        }
        processThread = std::thread([&]() {
            taskFunc();
            running = false;
        });
    }
}