Channel
=======

Cross-Thread Messaging Facility. The goal is to create a go-like Channel mechanism
for the C++ projects in the repository. This ***MAY*** in the future be extended for
Python projects, etc., whenever I feel bored enough to deal with cross-language type
issues, etc.

## Usage

```c++

Channel<int> data(4);  //create int channel with four(4) provisioned nodes. Will block on writes.
Channel<string> data(); //create a string channel with unbounded nodes.
```