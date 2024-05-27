Logger
======

## Objectives
* Create a comprehensive structured logger.

## Implementation

* See [examples](./example)

* We implement logger as a generic: `Logger[TGT LogTarget.LogTarget]`
* This generic implementation allows the target to be changed easily without
  sacrificing compatibility. At most this means we will need to add some 
  configuration for the new target (e.g. a File or hostname, etc.).
