Logger
======

(c) 2022 Sam Caldwell. All Rights Reserved.

---

## Overview

The `Logger` project is a very opinionated but versatile and extensible logging client which can be configured to
log to `null` (nowhere), `stdout`, `stderr`, `file`, `syslog` or `https`. The consumer makes configuration decisions
during `Logger` class instantiation and thereafter the logger can operate. At times during operation of the logger,
the user may wish to temporarily change the logging behavior. This can be done using the `Context` functionality
where a given state can be altered in a new context and then the changes can be reverted in a single call to the
prior state. Context is useful for many things, including signaling to the logs that a new "context" is being
logged in the application. For example, when the logger is executed in `main()` and control passes to a function
named `foo()`, the logger could use `Context` to signal this transition.

Each context has its own unique (serial) `contextId` which may indicate "logging depth" within an application.

The `Logger` intends to comply with RFC5424 and any deviation should be considered a bug.

---

## Usage

There are several ways to configure the `Logger` (e.g. manually or through `Configuration`):

---

### Option 1: Use `Configuration`:

The easiest way to set up `Logger` is to pass a reference to a `Configuration` object which has
the required parameters (discussed later):

```c++
Configuration cfg;
// ...Configuration setup stuff...
Logger log(&cfg);
```

The `Logger` class is configured by multiple variables. Within the monorepo, a set of configuration string patterns
are reserved for use in the `Configuration` project for configuring the `Logger` class but these values must be
passed specifically to the constructor:

| Configuration object | Type   | Description                                                                              |
|----------------------|--------|------------------------------------------------------------------------------------------|
| `log.level`          | String | set a log level consistent with Page7, [RFC5427](https://www.rfc-editor.org/rfc/rfc5427) |
| `log.output`         | String | set the log output (e.g. `stdout`, `stderr`,`syslog`,`https`                             |
| `log.destination`    | String | protocol and destination file/network address (e.g. syslog://addr:port)                  | 
| `log.tags`           | String | a comma-delimited list of tags to be appended to the log messages (not stdout)           |

---

### Option 2: Manual Configuration

There are several ways to manually configure `Logger`:

* Standard Default Logger
  ```c++
    Logger log() //default log constructor (log severity: info, output: stdout)
  ```
    * Log severity: Info
    * Output (writer): stdout

* Specified Severity and Writer
  ```c++
    Logger log(Logging::Severity::debug, Logging::Writer::Id::stderr);
  ```
    * Log severity is configured using the enumerated type `Logging::Severity` (see below)
    * Output (writer): is configured using the enumerated type `Logging::Writer::Id` (see below)

* Specified Tags
  ```c++
    Logger log(Logging::Severity::debug, 
               Logging::Writer::Id::stderr,
               "\"tag1\",\"tag2\",\"tag3\"");
  ```
    * Log severity: see above
    * Output (writer): see above
    * Tags: A quoted, comma-delimited string of `key:value` or `simple` tags (e.g. '"tag1","myTag:myValue"')
      each tag set is appended to each log line written by `Logger`.

* Specified Destination (required for `file`, `syslog` and `https` writers)
  ```c++
    Logger log(Logging::Severity::debug, 
               Logging::Writer::Id::stderr,
               "\"tag1\",\"tag2\",\"tag3\"",
               "/var/log/myapp/myapp.log);
  ```
    * Log severity: see above
    * Output (writer): see above
    * Tags: see above
    * Destination: a path/file or URL string used to indicate the log destination for certain log writers.
      For example: `file`, `syslog`, or `https` require a destination to be specified.

---

### Logging Context

When a user implements `Logger` there may be cases where the configuration state of the `Logger` may need to be
altered temporarily. There may also be cases where the user wishes to signal via the logs that the `Logger` is
recording events from a child scope. To facilitate this, the `Logger` supports a "context" feature. The `Logger`
context is a configuration state. By default, `Logger` has a "root context" (ContextId 0). But by calling the
`Logger::newContext()` method, the user may create a child context with a new severity level, writer, tags, etc.
However, unlike simply changing the configuration state of the root context, when the user is finished with the
context, a call to `Logger::popContext()` will revert the `Logger` object to its prior state in one easy function call.

Things to remember:

* The system can support up to 256 contexts (including the "root context").
* Each `Logger` instance will have its own concept of "context," independent of one another.
* Only one context is active at any given time (the last created).
* When a context is "popped" from the stack via `Logger::popContext()` the context state is lost.

#### Create a new context

To create a new context...

Assuming you have a `Logger` object named `log` already, you have a two options...

##### Option 1: Copy-create from the current context...

```c++
log.newContext();
```

In this case, the new context will take the parameters of the current context as the basis for the
new context. Only the `contextId` will increment. This could be useful where a user may want to
increment `contextId` as a simple way of asserting that subsequent logs are written in some child scope.

##### Option 2: Define the context parameters....

```c++
log.newContext(severity, writerId, *tags, *destination);
```

In the above case, the `newContext()` method expects a `Logging::Severity`, `Logging::Writer::Id` value as well
as pointers to the `tags` and `destination` strings.

##### Potential errors:

1. The `Logger::newContext()` method will validate inputs, and in the event an invalid input is encountered, an
   `out_of_range` exception will be thrown.
2. If a user attempts to create more than 256 contexts for a single `Logger` instance, an `out_of_range` exception
   will be thrown.

#### Remove/Revert a Context

When a user has finished with a logging context, the `Logger::popContext()` method may be invoked which will
pop the current context off the `Logger` context stack and destroy the `Context` object. Thereafter, all log
activity will be performed under the preceding (now current) context.

---

### Log Severity

This project is intended to comply with [RFC5424](https://www.rfc-editor.org/rfc/rfc5424).
Accordingly, the log severity values are defined in [severity.h](src/severity.h).

### Log Writers

There are currently four (4) working log writers and two (2) log writers which only require final implementation
and tests to be written. The log writers are identified by a single byte (`Logging::Writer::Id`) in the
[writer.h](src/writer.h) file.

Each writer is a class with standard interfaces used by `Context::write()` to write messages to a specific endpoint,
e.g. null, stdout, stderr, file, syslog, https, etc.

### Log Formatting

Currently, the log format is hard coded with plans to extend this in the future to support custom log formatting.
At this time, Log format is handled by `Context::write()` found in [Context/write.cpp](src/Context/write.cpp).

### Logging Messages

The core function of `Logger` is to accept log messages and to render them in a standard format to
some specified output. Once a `Logger` exists and the proper context is set, the user may write logs using--

* `Logger::emergency(messageString);`  // = 0 -> emergency; system is unusable
* `Logger::alert(messageString);`      // = 1 -> action must be taken immediately
* `Logger::critical(messageString);`   // = 2 -> critical condition
* `Logger::error(messageString);`      // = 3 -> error condition
* `Logger::warning(messageString);`    // = 4 -> warning condition
* `Logger::notice(messageString);`     // = 5 -> normal but significant condition
* `Logger::info(messageString);`       // = 6 -> informational message
* `Logger::debug(messageString);`      // = 7 -> debug-level messages
