/*
 * BackgroundProcess Class
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */
#include "BackgroundProcess.h"
BackgroundProcess::BackgroundProcess(void (*fp)()) {
    taskFunc = fp;
    running = false;
}