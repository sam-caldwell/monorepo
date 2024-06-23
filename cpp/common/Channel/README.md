Channel
=======

## Objective
Cross-Thread Messaging Facility. The goal is to create a go-like Channel mechanism
for the C++ projects in the repository. This ***MAY*** in the future be extended for
Python projects, etc., whenever I feel bored enough to deal with cross-language type
issues, etc.

## Overview
The Channel class template manages a queue of elements of type T, allowing threads to 
send (`<<`) and receive (`>>`) values safely across threads. It uses mutexes and condition 
variables for synchronization.

### Constructor
Description: Constructs a new Channel instance.
```c++
Channel();
```

### Destructor
Description: Destroys the Channel instance, ensuring proper cleanup and closing of the channel.
```c++
~Channel();
```

### Member Functions

#### Operator `<<`
Description: Sends a value of type T to the channel.
```c++
void operator<<(const T& value);
```

#### Operator `>>`
Description: Receives a value from the channel. Returns true if successful, false if the channel is 
closed or empty.
```c++
bool operator>>(T& value);
```

#### Method: `close()`
Description: Closes the channel, preventing further sends and allowing remaining values to be received.
```c++
void close();
```

#### Method: `is_closed()`
Description: Checks if the channel is closed.
```c++
bool is_closed() const;
```

## Example Usage
```c++
#include "Channel.h" // Include the header file where Channel class is defined
#include <iostream>
#include <thread>
#include <chrono>

int main() {
Channel<int> ch;

    // Producer thread
    std::thread producer([&ch]() {
        for (int i = 0; i < 5; ++i) {
            ch << i; // Send using <<
            std::this_thread::sleep_for(std::chrono::milliseconds(500));
        }
        ch.close();
    });

    // Consumer thread
    std::thread consumer([&ch]() {
        int value;
        while (ch >> value) { // Receive using >>
            std::cout << "Received: " << value << std::endl;
        }
    });

    producer.join();
    consumer.join();

    return 0;
}
```
### Explanation of Example
#### `main()` Function:
* Creates a Channel<int> instance ch for communication between producer and consumer threads.
* Spawns a producer thread that sends integers (0 to 4) to ch using `<<`.
* Spawns a consumer thread that receives integers from ch using `>>`.

#### `producer` Thread:
* Uses operator `<<` to send integers to the channel.
* Closes the channel (ch.close()) after sending all values.

#### `consumer` Thread:
* Uses operator `>>` to receive integers from the channel.
* Continues receiving until the channel is closed.

#### Notes
* The Channel class template ensures thread-safe communication and synchronization using mutexes
  and condition variables.
* Proper cleanup is handled by the destructor (~Channel()), ensuring resources are released when
  the Channel object goes out of scope.
* Use ch.is_closed() to check if the channel is closed before attempting to send or receive data.
* This documentation provides an overview of the Channel class template, its usage, and example 
  scenarios demonstrating how to use it effectively for inter-thread communication in C++.
