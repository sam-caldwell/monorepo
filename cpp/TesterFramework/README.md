Tester
======

(c) 2022 Sam Caldwell. All Rights Reserved.

---

## Overview

The [`Tester`]() is a test suite built solely because I felt like doing it. Sure, I could use something off 
the shelf. But as a personal project, I built this just to do it and to reinforce my skills in doing so.

## Goals

- Keep it simple. A person should be able to use it within a couple of hours.
- Keep it extensible. As I need more functionality, I should be able to add it.
- Ensure it can be monitored through in-test metrics emissions.

## ToDo

- Implement a better stats receiver first.
- Add metrics emitters for start time, stop time and errors.
- Avoid logs in favor of `key[tag1,tag2,tag3]:value` events.

## Usage

There are three (3) parts of `Tester`:

| Component      | Type  | Description                                                                                | 
|----------------|-------|--------------------------------------------------------------------------------------------|
| `Tester`       | Class | The `Tester` class used to create a test object (located in test.cpp) within each project. |
| `TestBase`     | Class | Provides a base class for building test classes to be executed by `Tester`.                |
| `Mock::Stdout` | Class | This class intercepts `stdout` to inspect content written during tests.                    |
| `Mock::Stderr` | Class | This class intercepts `stderr` to inspect content written during tests.                    |
| `Mock::File`   | Class | This class intercepts `file` to inspect content written during tests.                      |
| `Mock::Tcp`    | Class | This class intercepts TCP network traffic to inspect content during tests.                 |

