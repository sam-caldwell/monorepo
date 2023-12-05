types
=====
(c) 2022 Sam Caldwell.  All Rights Reserved.

This is a generic library of various simple type
definitions.  Some are easy typedefs and enums.
Others are a bit more complex types.

## Usage
This project creates top-level C++ header files (.h)
which present the type to the rest of the monorepo
in an easier path.

For example, rather than this...
```c++
#include "types/src/SortOrder.h"
```
The consumer only needs to implement
```c++
#include "types/SortOrder.h"
```
and this `SortOrder.h` includes the various C++ files
which create our SortOrder enum.


## Types

| Type                | Category | Description                                                                             |
|---------------------|----------|-----------------------------------------------------------------------------------------|
| SortOrder           | Enum     | Defines a sort order as (ascending, descending or undefined)                            |
| ValueTypes          | Enum     | Defines a value type for arbitraryValue* projects to store strongly typed runtime data. |
| uint                | Typedef  | Defines a simple unsigned int standard for the monorepo                                 | 
| ConfigStateMap | Class    | Defines a map between two strings and a ValueTypes value                                |
