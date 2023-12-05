Application
===========

(c) 2022 Sam Caldwell. All Rights Reserved.

---

## Overview

The `Application` project is a foundation class for building software applications within this monorepo, and
it can support anything from a simple command-line tool to a larger service application. The goal is to increase
developer velocity by providing an opinionated C++ framework for starting application development within the
monorepo.

---

## Usage

This project is intended to be used as a base class for your specific Application project. The `Application`
class will require a `ConfigStateMap` object (passed by reference) and an optional `SignalHandlers` table to
get started. The `ConfigStateMap` will define the `Configuration` parameters the application will load from
environment variables, command-line parameters or a configuration file into an in-memory structure that can be
used throughout the lifecycle of the software.

## Features: What does `Application` provide?

The `Application` class provides the following standardized features to your application software:

* Configuration services
    * Environment Variable Validation/Loading
    * Command-Line Argument Parsing and Validation
    * Configuration File Parsing and Validation
    * Centralized configuration storage.

* Signal Processing / Handling
    * Facilities to register signal handlers.

* Observability services
    * Logging
    * Metrics and event emissions

These are discussed in more detail, below:

### Configuration Services

#### Data Sources and Automatic Loading/Mapping

Given a `ConfigStateMap` object, the `Application` will load the current configuration from the following three (3)
sources in the following order of precedence:

* Environment Variables (`EnvironmentVariableParser`)
* Command-Line Arguments (`CommandLineParser`)
* Config File (`ConfigFileParser`)

Thus, where an environment variable is loaded and maps to the same internal name as a command-line argument, the
command-line argument value will overwrite the value loaded from the environment variable. Likewise, a configuration
file value will overwrite a command-line argument.

> ***Todo***:
>
> We intend to integrate with 1Password and other secrets managers (e.g. GPG-encrypted files) in the future to load
> secrets directly from the secrets managers as a fourth configuration source. This would ultimately overwrite the
> value from any of the other three sources.

#### In-Memory Configuration Store

Once the Application is initialized and the configuration is loaded from the supported data sources, a software
application can read/write data to/from this store during the application lifecycle. The data stored in the
`Configuration` object is strongly typed using the `ArbitraryValue` project and heterogeneous data types are supported.
Thus, a diverse collection of data objects can be maintained using `Bool`, `Double`, `Float`, `Int`, `String` or `Uint`
data types.

----

### Signal Processing / Handling

The `Application` class provides an easy-to-use vector table for processing operating system signals. The user
only needs to register a signal handler function with the `Application` signal handler facilities and any signals
will trigger the associated handler function(s).


----

### Observability Services

Modern applications require the means to observe their performance. The `Application` class provides this through
its `Logger` and `Events` facilities. These opinionated observability foundations allow the application developer to
generate structured text logs or high-cardinality metrics at any point during program execution.