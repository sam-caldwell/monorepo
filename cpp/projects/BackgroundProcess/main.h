/*
 * Background Process Test Program
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */

#include "BackgroundProcess.h"

// Example usage
void myTask() {
    for (int i = 0; i < 5; ++i) {
        std::cout << "Task iteration: " << i << std::endl;
        std::this_thread::sleep_for(std::chrono::seconds(1));
    }
}

int main() {
    BackgroundProcess background(myTask);

    background.start();

    // Do other work...

    if (background.isRunning()) {
        std::cout << "Background process is running." << std::endl;
    }

    // Do more work...

    background.stop();

    return 0;
}
