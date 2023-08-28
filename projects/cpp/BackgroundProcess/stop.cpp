/*
 * BackgroundProcess Class
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */
#include "BackgroundProcess.h"
void BackgroundProcess::stop() {
    if (running) {
        running = false;
        if (processThread.joinable()) {
            processThread.join();
        }
    }
}