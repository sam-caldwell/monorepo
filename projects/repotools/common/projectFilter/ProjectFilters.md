repotools.ProjectFilters
========================

## Description

This file discusses how ListProjectsFilter works.

## Theory of Operation

1. The ListProjectsFilter is a 32-bit unsigned integer.
2. The 32-bits provide a bit field for representing filters we can use when listing projects
3. The ListProjectsFilter represents two pieces of data:
    * General filters (or general flags)
    * Language filters (or language flags)

```text
ListProjectsFilter
------------------
|3 3 2 2 2 2 2 2 2 2 2 2 1 1 1 1 | 1 1 1 1 1 1 0 0 0 0 0 0 0 0 0 0| 
|1 0 9 8 7 6 5 4 3 2 1 0 9 8 7 6 | 5 4 3 2 1 0 9 8 7 6 5 4 3 2 1 0|
|--------------------------------|--------------------------------|
|     16-bit language specifier  |      16-bit general flags      |
|--------------------------------|--------------------------------|
```

> The above gives us support for 16 languages and if we needed more we
> could simply implement uint64 and continue adding languages with the
> most significant bits.

<style>
.table-striped tr:nth-child(even) {
    background-color: #f2f2f2;
}
</style>

| Bit |      Value |       Name       | Description                                                               |
|----:|-----------:|:----------------:|:--------------------------------------------------------------------------|
|  00 |       0x01 |   HideEnabled    | filter flag indicating enabled projects should be filtered                |
|  01 |       0x02 |   HideDisabled   | filter flag indicating disabled projects should be filtered               |
|  02 |       0x04 |     Windows      | filter flag indicating Windows support                                    |
|  03 |       0x08 |      Linux       | filter flag indicating Linux support                                      |
|  04 |       0x10 |      Darwin      | filter flag indicating MacOS support                                      | 
|  05 |       0x20 |     Command      | filter flag indicating projects which compile to executables              |
|  06 |       0x40 |     Package      | filter flag indicating reusable code projects which produce no executable |
|  07 |       0x80 |      <open>      | <unallocated>                                                             |
|  08 |      0x100 |      <open>      | <unallocated>                                                             |
|  09 |      0x200 |      <open>      | <unallocated>                                                             |
|  10 |      0x400 |      <open>      | <unallocated>                                                             |
|  11 |      0x800 |      <open>      | <unallocated>                                                             |
|  12 |     0x1000 |      <open>      | <unallocated>                                                             |
|  13 |     0x2000 |      <open>      | <unallocated>                                                             |
|  14 |     0x4000 |      <open>      | <unallocated>                                                             |
|  15 |     0x8000 |      <open>      | <unallocated>                                                             |
|     |            |                  |                                                                           |
|  16 |    0x10000 |    LanguageGo    | Indicates project is primarily Golang                                     |
|  17 |    0x20000 |    LanguageC     | Indicates project is primarily C                                          |
|  18 |    0x40000 |   LanguageCpp    | Indicates project is primarily C++                                        |
|  19 |    0x80000 | LanguageAsmAmd64 | Indicates project is primarily AMD64 (X86_64) Assembly                    |
|  20 |   0x100000 | LanguageAsmArm64 | Indicates project is primarily Arm64 Assembly                             |
|  21 |   0x200000 |   LanguageRust   | Indicates project is primarily Rust                                       |
|  22 |   0x400000 |  LanguageNodeJs  | Indicates project is primarily NodeJs / Typescript                        |
|  23 |   0x800000 |  LanguagePython  | Indicates project is primarily Python                                     |
|  24 |  0x1000000 |      <open>      | <unallocated>                                                             |
|  25 |  0x2000000 |      <open>      | <unallocated>                                                             |
|  26 |  0x4000000 |      <open>      | <unallocated>                                                             |
|  27 |  0x8000000 |      <open>      | <unallocated>                                                             |
|  28 | 0x10000000 |      <open>      | <unallocated>                                                             |
|  29 | 0x20000000 |      <open>      | <unallocated>                                                             |
|  30 | 0x40000000 |      <open>      | <unallocated>                                                             |
|  31 | 0x80000000 |      <open>      | <unallocated>                                                             |



