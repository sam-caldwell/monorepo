lock
=====
(c) 2022 Sam Caldwell.  All Rights Reserved.

A simple library of locking mechanisms.

## SimpleLock
* a simple locking mechanism
* provides on(), off(), and wait().

* on() - enable lock
* off() - disable lock
* wait() - wait until lock is released or timeout expires.  Throw exception on timeout.
* check() - return true/false if locked or unlocked, respectively.