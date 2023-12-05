LogServer
=========

(c) 2022 Sam Caldwell. All Rights Reserved.

---

## Overview

The [`LogServer`]() project is a standalone server application which is designed to work with the client
[`Logger`](../Logger) project. This server will use a proprietary network protocol to stream logs from the client to
this server and from there to disk.

## Usage

From a commandline, we launch the server using...

```bash
./logserver --config <filename>
```

-- or --

```bash
./logserver --bind <ipAddr> --port <port> --target <directory>
```

## Configuration parameters

The following configuration parameters are reserved for `LogServer`:

| Parameter            | Type   | Required? | Description                             |
|----------------------|--------|-----------|-----------------------------------------|
| `config.file`        | string | No        | The config file                         |
| `logServer.net.ip`   | string | No        | Host/IP address for receiving logs      |
| `logServer.net.port` | uint   | No        | Network port for receiving logs         |
| `logServer.target`   | string | No        | Target directory for logs to be written |
| `log.level`          | string | No        | Log severity                            |
| `log.output`         | string | No        | Log writer                              | 
| `log.destination`    | string | No        | Log destination (for syslog,file,https) |



