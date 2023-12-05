TcpServer
=========

(c) 2022 Sam Caldwell. All Rights Reserved.

---

## Overview

The [`TcpServer`](../TcpServer) project creates a re-usable server for listening on a given IP address and port for
TCP network communications and routing received messages to provided handler functions.  This project provides the
low-level network setup and operations framework only.  Any higher-level routing is the responsibility of the consuming
functions.

## Usage

```c++
TcpServer server(&config, &handlerClass);
```

## Theory of Operation

### Initialization
1. Class constructor should validate the IPAddress and port inputs
2. Class constructor should verify the handler class reference is not null.

### 