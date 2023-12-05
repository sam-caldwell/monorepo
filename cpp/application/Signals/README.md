Signals
=======

(c) 2022 Sam Caldwell. All Rights Reserved.

---

## Overview

The `Signals` project is part of the larger `Application` project, and this project provides the `Application` class
an extensible mechanism for mapping application-specific signal handler functions to the underlying signal sources in
a safe and predictable way. As with much of this monorepo, `Signals` is an opinionated approach to signal handling,
and it is intended to pre-bake solutions to many common problems in a scalable way for all projects in the repo.

This project provides several features:

1. A SignalHandlers vector table for mapping a signal interrupt to a given signal handler class.
2. A base class for building signal handler classes.

---

## Usage

## Supported Signals

The following signals are tested with this library:

| Signal    | int | description                                                             |
|-----------|-----|-------------------------------------------------------------------------|
| SIGHUP    | 1   | Hangup detected on controlling terminal or death of controlling process |
| SIGINT    | 2   | Interrupt from keyboard                                                 |
| SIGQUIT   | 3   | Quit from keyboard                                                      |
| SIGILL    | 4   | Illegal Instruction                                                     |
| SIGTRAP   | 5   | Trace/breakpoint trap                                                   |
| SIGABRT   | 6   | Abort signal from abort                                                 |
| SIGFPE    | 8   | Floating-point exception                                                |
| SIGSEGV   | 11  | Invalid memory reference                                                |
| SIGPIPE   | 13  | Broken pipe: write to pipe with no readers; see pipe(7)                 |
| SIGALRM   | 14  | Timer signal from alarm                                                 |
| SIGTERM   | 15  | Termination signal                                                      |
| SIGTTIN   | 21  | Terminal input for background process                                   |
| SIGTTOU   | 22  | Terminal output for background process                                  |
| SIGXCPU   | 24  | CPU time limit exceeded (4.2BSD)                                        |
| SIGXFSZ   | 25  | File size limit exceeded (4.2BSD)                                       |
| SIGVTALRM | 26  | Virtual alarm clock (4.2BSD)                                            |
| SIGPROF   | 27  | Profiling timer expired                                                 |
| SIGWINCH  | 28  | Window resize signal (4.3BSD, Sun)                                      |

> NOTE:
>
> See https://man7.org/linux/man-pages/man7/signal.7.html
> There are several signals in the above page which cannot be captured in this library. This
> also presupposes a non-Windows, x86 or ARM CPU.

