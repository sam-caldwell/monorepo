/*
 * BackgroundProcess Class
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */
#ifndef BACKGROUND_PROCESS_H
#define BACKGROUND_PROCESS_H

#include <iostream>
#include <thread>
#include <atomic>

class BackgroundProcess {
private:
    void (*taskFunc)(); // Function pointer for the task
    std::thread processThread;
    std::atomic<bool> running;

public:
    BackgroundProcess(void (*fp)());

    ~BackgroundProcess();

    void start();

    void stop();

    bool isRunning() const;
};

#endif
