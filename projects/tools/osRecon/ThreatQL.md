ThreatQL
========

## Objectives:

* Define a Threat Intelligence Query Language (ThreatQL).

## Query Pattern (Theory)

The theoretical pattern for a ThreatQL query is:

```text
    <table>.<columnA>[,<columnB>,<columnC>,...,<columnZ>]|<column><operator><value>
```

For example:

| Query                                       | Description                                                         |
|---------------------------------------------|---------------------------------------------------------------------|
| `cpu.loadavg1`                              | Query returns 1m load average.                                      |
| `cpu.loadavg1,loadavg5,loadavg15`           | The following query will return the 1m, 5m and 15m load average     |
| `cpu.user`                                  | The following query will show %user cpu load                        |
| `cpu.wait`                                  | return wait time                                                    |
| `cpu.idle`                                  | return idle time                                                    |
| `cpu.cores`                                 | return number cores                                                 |
| `mem.total`                                 | total memory (KB)                                                   |
| `mem.free`                                  | free memory (KB)                                                    |
| `mem.swap`                                  | virtual memory (KB)                                                 |
| `process.pid,user,command\| pid==3312`      | return the process Id, user and command of a process with pid 3312  |
| `file.name,size,hash\|name="/usr/bin/foo*"` | return the file name, size and hash for a file with a given pattern |

1. Note that the query only selects fields and filters the results with a 'where' clause (represented to the 
   right of a pipe '|' character).
2. The query language does not sort, aggregate or paginate query results.  Instead, the client is allowed to paginate
   results before emitting them to the server.
