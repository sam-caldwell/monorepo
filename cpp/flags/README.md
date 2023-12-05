preprocessor flags
==================

(c) 2022 Sam Caldwell. All Rights Reserved.

This project contains some global, generally useful preprocessor directives, such as TRUE/FALSE,
OPSYS and CPU_ARCH. We include these into the build process to simplify things later.

| FLAG     | VALUE        | NUMERIC VALUE | DESCRIPTION                               |
|----------|--------------|---------------|-------------------------------------------|
| OPSYS    | OPSYS_WINDOW | 0             | OPSYS defines the operating system family |
|          | OPSYS_LINUX  | 1             |                                           |
|          | OPSYS_MACOS  | 2             |                                           |
| CPU_ARCH | ARCH_X32     | 0             | Intel/AMD 32-bit architecture             |
|          | ARCH_X64     | 1             | Intel/AMD 64-bit architecture             |
|          | ARCH_ARM32   | 2             | ARM 32-bit architecture                   |
|          | ARCH_ARM64   | 3             | ARM 64-bit architecture                   |
| TRUE     | ---          | 1             | Boolean true                              |
| FALSE    | ---          | 0             | Boolean false                             |

