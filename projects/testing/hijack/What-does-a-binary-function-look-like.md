What does a Binary Function Look Like In Golang
===============================================

## Description
This document is a working document while I try to figure out why the hijack operation is not working.
I suspect I'm writing to the wrong place in memory.

## What format?
We are on a macOS machine, so we are using Mach-O format.

## Mach-O Format
```text
-------------------------------------------
| Mach-O Header                          |
-------------------------------------------
| Load Commands                          |
|   * Segment Load Command               |
|     * Section                          |
|   * Dylib Load Command                 |
|   * Dyld Info Load Command             |
|   * Symtab Load Command                |
|   * DySymtab Load Command              |
|   * ...                                |
-------------------------------------------
| Segments and their contained Sections  |
|   * __TEXT                             |
|     * __text                           |
|     * __stubs                          |
|     * __stub_helper                    |
|     * __const                          |
|     * __cstring                        |
|     * ...                              |
|   * __DATA                             |
|     * __data                           |
|     * __bss                            |
|     * ...                              |
|   * __LINKEDIT                         |
-------------------------------------------
| Raw Segment Data                       |
-------------------------------------------
| Link Edit Information                  |
|   * Dynamic Symbol Table               |
|   * Function Starts                    |
|   * Data in Code                       |
|   * Code Signature                     |
-------------------------------------------
```
### Details:
* ***Mach-O Header***: Describes the Mach-O file (e.g., whether it's for 32-bit or 64-bit, little or big endian, etc.)
* ***Load Commands***: Metadata about how to load the file into memory, includes segment commands, dylib commands, etc.
* ***Segments***: Contiguous regions of memory. Each segment contains zero or more sections.
* ***Sections***: Smaller divisions of a segment that contain code or data of some kind.
* ***Raw Segment Data***: The actual bytes of code or data to be loaded into memory.
* ***Link Edit Information***: Information for the dynamic linker, including the symbol table and other linking metadata.

## Analysis of the Mach-O file using hexdump
```text
00000000  cf fa ed fe 07 00 00 01  03 00 00 00 02 00 00 00  |................|
00000010  0c 00 00 00 70 09 00 00  00 00 00 00 00 00 00 00  |....p...........|
00000020  19 00 00 00 48 00 00 00  5f 5f 50 41 47 45 5a 45  |....H...__PAGEZE|
00000030  52 4f 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |RO..............|
00000040  00 00 00 01 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000050  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000060  00 00 00 00 00 00 00 00  19 00 00 00 78 02 00 00  |............x...|
00000070  5f 5f 54 45 58 54 00 00  00 00 00 00 00 00 00 00  |__TEXT..........|
00000080  00 00 00 01 00 00 00 00  00 f0 13 00 00 00 00 00  |................|
00000090  00 00 00 00 00 00 00 00  00 f0 13 00 00 00 00 00  |................|
000000a0  07 00 00 00 05 00 00 00  07 00 00 00 00 00 00 00  |................|
000000b0  5f 5f 74 65 78 74 00 00  00 00 00 00 00 00 00 00  |__text..........|
000000c0  5f 5f 54 45 58 54 00 00  00 00 00 00 00 00 00 00  |__TEXT..........|
000000d0  00 10 00 01 00 00 00 00  d9 72 09 00 00 00 00 00  |.........r......|
000000e0  00 10 00 00 05 00 00 00  00 00 00 00 00 00 00 00  |................|
000000f0  00 04 00 80 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000100  5f 5f 73 79 6d 62 6f 6c  5f 73 74 75 62 31 00 00  |__symbol_stub1..|
00000110  5f 5f 54 45 58 54 00 00  00 00 00 00 00 00 00 00  |__TEXT..........|
00000120  e0 82 09 01 00 00 00 00  26 01 00 00 00 00 00 00  |........&.......|
00000130  e0 82 09 00 05 00 00 00  00 00 00 00 00 00 00 00  |................|
00000140  08 04 00 80 00 00 00 00  06 00 00 00 00 00 00 00  |................|
00000150  5f 5f 72 6f 64 61 74 61  00 00 00 00 00 00 00 00  |__rodata........|
00000160  5f 5f 54 45 58 54 00 00  00 00 00 00 00 00 00 00  |__TEXT..........|
00000170  20 84 09 01 00 00 00 00  d3 d3 03 00 00 00 00 00  | ...............|
00000180  20 84 09 00 05 00 00 00  00 00 00 00 00 00 00 00  | ...............|
00000190  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
000001a0  5f 5f 74 79 70 65 6c 69  6e 6b 00 00 00 00 00 00  |__typelink......|
000001b0  5f 5f 54 45 58 54 00 00  00 00 00 00 00 00 00 00  |__TEXT..........|
000001c0  00 58 0d 01 00 00 00 00  9c 05 00 00 00 00 00 00  |.X..............|
000001d0  00 58 0d 00 05 00 00 00  00 00 00 00 00 00 00 00  |.X..............|
000001e0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
000001f0  5f 5f 69 74 61 62 6c 69  6e 6b 00 00 00 00 00 00  |__itablink......|
00000200  5f 5f 54 45 58 54 00 00  00 00 00 00 00 00 00 00  |__TEXT..........|
00000210  a0 5d 0d 01 00 00 00 00  88 00 00 00 00 00 00 00  |.]..............|
00000220  a0 5d 0d 00 05 00 00 00  00 00 00 00 00 00 00 00  |.]..............|
00000230  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000240  5f 5f 67 6f 73 79 6d 74  61 62 00 00 00 00 00 00  |__gosymtab......|
00000250  5f 5f 54 45 58 54 00 00  00 00 00 00 00 00 00 00  |__TEXT..........|
00000260  28 5e 0d 01 00 00 00 00  00 00 00 00 00 00 00 00  |(^..............|
00000270  28 5e 0d 00 00 00 00 00  00 00 00 00 00 00 00 00  |(^..............|
00000280  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000290  5f 5f 67 6f 70 63 6c 6e  74 61 62 00 00 00 00 00  |__gopclntab.....|
000002a0  5f 5f 54 45 58 54 00 00  00 00 00 00 00 00 00 00  |__TEXT..........|
000002b0  40 5e 0d 01 00 00 00 00  c0 84 06 00 00 00 00 00  |@^..............|
000002c0  40 5e 0d 00 05 00 00 00  00 00 00 00 00 00 00 00  |@^..............|
000002d0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
000002e0  19 00 00 00 28 02 00 00  5f 5f 44 41 54 41 00 00  |....(...__DATA..|
000002f0  00 00 00 00 00 00 00 00  00 f0 13 01 00 00 00 00  |................|
00000300  90 97 04 00 00 00 00 00  00 f0 13 00 00 00 00 00  |................|
00000310  00 7f 01 00 00 00 00 00  03 00 00 00 03 00 00 00  |................|
00000320  06 00 00 00 00 00 00 00  5f 5f 67 6f 5f 62 75 69  |........__go_bui|
00000330  6c 64 69 6e 66 6f 00 00  5f 5f 44 41 54 41 00 00  |ldinfo..__DATA..|
00000340  00 00 00 00 00 00 00 00  00 f0 13 01 00 00 00 00  |................|
00000350  60 01 00 00 00 00 00 00  00 f0 13 00 04 00 00 00  |`...............|
00000360  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000370  00 00 00 00 00 00 00 00  5f 5f 6e 6c 5f 73 79 6d  |........__nl_sym|
00000380  62 6f 6c 5f 70 74 72 00  5f 5f 44 41 54 41 00 00  |bol_ptr.__DATA..|
00000390  00 00 00 00 00 00 00 00  60 f1 13 01 00 00 00 00  |........`.......|
000003a0  88 01 00 00 00 00 00 00  60 f1 13 00 02 00 00 00  |........`.......|
000003b0  00 00 00 00 00 00 00 00  06 00 00 00 31 00 00 00  |............1...|
000003c0  00 00 00 00 00 00 00 00  5f 5f 6e 6f 70 74 72 64  |........__noptrd|
000003d0  61 74 61 00 00 00 00 00  5f 5f 44 41 54 41 00 00  |ata.....__DATA..|
000003e0  00 00 00 00 00 00 00 00  00 f3 13 01 00 00 00 00  |................|
000003f0  c0 07 01 00 00 00 00 00  00 f3 13 00 05 00 00 00  |................|
00000400  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000410  00 00 00 00 00 00 00 00  5f 5f 64 61 74 61 00 00  |........__data..|
00000420  00 00 00 00 00 00 00 00  5f 5f 44 41 54 41 00 00  |........__DATA..|
00000430  00 00 00 00 00 00 00 00  c0 fa 14 01 00 00 00 00  |................|
00000440  30 74 00 00 00 00 00 00  c0 fa 14 00 05 00 00 00  |0t..............|
00000450  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000460  00 00 00 00 00 00 00 00  5f 5f 62 73 73 00 00 00  |........__bss...|
00000470  00 00 00 00 00 00 00 00  5f 5f 44 41 54 41 00 00  |........__DATA..|
00000480  00 00 00 00 00 00 00 00  00 6f 15 01 00 00 00 00  |.........o......|
00000490  a0 e0 02 00 00 00 00 00  00 00 00 00 05 00 00 00  |................|
000004a0  00 00 00 00 00 00 00 00  01 00 00 00 00 00 00 00  |................|
000004b0  00 00 00 00 00 00 00 00  5f 5f 6e 6f 70 74 72 62  |........__noptrb|
000004c0  73 73 00 00 00 00 00 00  5f 5f 44 41 54 41 00 00  |ss......__DATA..|
000004d0  00 00 00 00 00 00 00 00  a0 4f 18 01 00 00 00 00  |.........O......|
000004e0  f0 37 00 00 00 00 00 00  00 00 00 00 05 00 00 00  |.7..............|
000004f0  00 00 00 00 00 00 00 00  01 00 00 00 00 00 00 00  |................|
00000500  00 00 00 00 00 00 00 00  19 00 00 00 78 02 00 00  |............x...|
00000510  5f 5f 44 57 41 52 46 00  00 00 00 00 00 00 00 00  |__DWARF.........|
00000520  00 90 18 01 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000530  00 70 15 00 00 00 00 00  61 c3 08 00 00 00 00 00  |.p......a.......|
00000540  00 00 00 00 00 00 00 00  07 00 00 00 00 00 00 00  |................|
00000550  5f 5f 7a 64 65 62 75 67  5f 61 62 62 72 65 76 00  |__zdebug_abbrev.|
00000560  5f 5f 44 57 41 52 46 00  00 00 00 00 00 00 00 00  |__DWARF.........|
00000570  00 90 18 01 00 00 00 00  29 01 00 00 00 00 00 00  |........).......|
00000580  00 70 15 00 00 00 00 00  00 00 00 00 00 00 00 00  |.p..............|
00000590  00 00 00 02 00 00 00 00  00 00 00 00 00 00 00 00  |................|
000005a0  5f 5f 7a 64 65 62 75 67  5f 6c 69 6e 65 00 00 00  |__zdebug_line...|
000005b0  5f 5f 44 57 41 52 46 00  00 00 00 00 00 00 00 00  |__DWARF.........|
000005c0  29 91 18 01 00 00 00 00  f6 13 02 00 00 00 00 00  |)...............|
000005d0  29 71 15 00 00 00 00 00  00 00 00 00 00 00 00 00  |)q..............|
000005e0  00 00 00 02 00 00 00 00  00 00 00 00 00 00 00 00  |................|
000005f0  5f 5f 7a 64 65 62 75 67  5f 66 72 61 6d 65 00 00  |__zdebug_frame..|
00000600  5f 5f 44 57 41 52 46 00  00 00 00 00 00 00 00 00  |__DWARF.........|
00000610  1f a5 1a 01 00 00 00 00  a2 64 00 00 00 00 00 00  |.........d......|
00000620  1f 85 17 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000630  00 00 00 02 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000640  5f 5f 64 65 62 75 67 5f  67 64 62 5f 73 63 72 69  |__debug_gdb_scri|
00000650  5f 5f 44 57 41 52 46 00  00 00 00 00 00 00 00 00  |__DWARF.........|
00000660  c1 09 1b 01 00 00 00 00  3b 00 00 00 00 00 00 00  |........;.......|
00000670  c1 e9 17 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000680  00 00 00 02 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000690  5f 5f 7a 64 65 62 75 67  5f 69 6e 66 6f 00 00 00  |__zdebug_info...|
000006a0  5f 5f 44 57 41 52 46 00  00 00 00 00 00 00 00 00  |__DWARF.........|
000006b0  fc 09 1b 01 00 00 00 00  90 d9 03 00 00 00 00 00  |................|
000006c0  fc e9 17 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
000006d0  00 00 00 02 00 00 00 00  00 00 00 00 00 00 00 00  |................|
000006e0  5f 5f 7a 64 65 62 75 67  5f 6c 6f 63 00 00 00 00  |__zdebug_loc....|
000006f0  5f 5f 44 57 41 52 46 00  00 00 00 00 00 00 00 00  |__DWARF.........|
00000700  8c e3 1e 01 00 00 00 00  e2 dc 01 00 00 00 00 00  |................|
00000710  8c c3 1b 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000720  00 00 00 02 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000730  5f 5f 7a 64 65 62 75 67  5f 72 61 6e 67 65 73 00  |__zdebug_ranges.|
00000740  5f 5f 44 57 41 52 46 00  00 00 00 00 00 00 00 00  |__DWARF.........|
00000750  6e c0 20 01 00 00 00 00  f3 92 00 00 00 00 00 00  |n. .............|
00000760  6e a0 1d 00 00 00 00 00  00 00 00 00 00 00 00 00  |n...............|
00000770  00 00 00 02 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000780  19 00 00 00 48 00 00 00  5f 5f 4c 49 4e 4b 45 44  |....H...__LINKED|
00000790  49 54 00 00 00 00 00 00  00 90 18 01 00 00 00 00  |IT..............|
000007a0  30 75 01 00 00 00 00 00  00 40 1e 00 00 00 00 00  |0u.......@......|
000007b0  30 75 01 00 00 00 00 00  01 00 00 00 01 00 00 00  |0u..............|
000007c0  00 00 00 00 00 00 00 00  32 00 00 00 18 00 00 00  |........2.......|
000007d0  01 00 00 00 00 0d 0a 00  00 0d 0a 00 00 00 00 00  |................|
000007e0  05 00 00 00 b8 00 00 00  04 00 00 00 2a 00 00 00  |............*...|
000007f0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000870  60 f2 05 01 00 00 00 00  00 00 00 00 00 00 00 00  |`...............|
00000880  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000890  00 00 00 00 00 00 00 00  02 00 00 00 18 00 00 00  |................|
000008a0  00 40 1e 00 7c 09 00 00  48 d9 1e 00 e8 db 00 00  |.@..|...H.......|
000008b0  0b 00 00 00 50 00 00 00  00 00 00 00 4b 09 00 00  |....P.......K...|
000008c0  4b 09 00 00 00 00 00 00  4b 09 00 00 31 00 00 00  |K.......K...1...|
000008d0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
000008e0  00 00 00 00 00 00 00 00  c0 d7 1e 00 62 00 00 00  |............b...|
000008f0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000900  0e 00 00 00 20 00 00 00  0c 00 00 00 2f 75 73 72  |.... ......./usr|
00000910  2f 6c 69 62 2f 64 79 6c  64 00 00 00 00 00 00 00  |/lib/dyld.......|
00000920  0c 00 00 00 38 00 00 00  18 00 00 00 00 00 00 00  |....8...........|
00000930  00 00 00 00 00 00 00 00  2f 75 73 72 2f 6c 69 62  |......../usr/lib|
00000940  2f 6c 69 62 53 79 73 74  65 6d 2e 42 2e 64 79 6c  |/libSystem.B.dyl|
00000950  69 62 00 00 00 00 00 00  0c 00 00 00 38 00 00 00  |ib..........8...|
00000960  18 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000970  2f 75 73 72 2f 6c 69 62  2f 6c 69 62 72 65 73 6f  |/usr/lib/libreso|
00000980  6c 76 2e 39 2e 64 79 6c  69 62 00 00 00 00 00 00  |lv.9.dylib......|
00000990  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00001000  ff 20 47 6f 20 62 75 69  6c 64 20 49 44 3a 20 22  |. Go build ID: "|
00001010  6e 37 61 37 4d 63 4d 33  2d 5f 30 48 66 44 64 58  |n7a7McM3-_0HfDdX|
00001020  4b 59 4a 79 2f 39 38 35  72 46 53 31 6d 68 51 2d  |KYJy/985rFS1mhQ-|
00001030  52 52 69 52 49 52 30 4e  4c 2f 4a 5a 34 42 4a 54  |RRiRIR0NL/JZ4BJT|
00001040  38 6c 6b 43 6b 5a 44 6b  6b 55 7a 4e 6a 50 2f 78  |8lkCkZDkkUzNjP/x|
00001050  4d 48 43 77 43 44 49 49  72 46 55 36 4a 65 41 57  |MHCwCDIIrFU6JeAW|
00001060  39 61 46 22 0a 20 ff cc  cc cc cc cc cc cc cc cc  |9aF". ..........|
```
* At address 00000000 we confirm the file is Mach-O from the magic number: `cf fa ed fe`.
* Following the magic number we identify CPU type and subtype: `07 00 00 01`
* We then see the "Load Commands" for our file from 00000008 until 00001000
  * _PAGEZERO: This segment occupies the virtual memory region from address 0 to a higher address,
        ensuring that null pointers in the program do not point to valid memory.
  * __TEXT and __text: This is where the compiled machine code resides.
  * __DATA: This is typically where global variables and static variables that are initialized in the code are stored.
  * __symbol_stub1: These are stubs for system functions.
  * __rodata: This section contains read-only data.
  * __typelink: 
  * __itablink: 
  * __gosymtab: symbols table
  * __gopclntab: 
  * __LINKEDIT: linker information
  * /usr/lib/dyld: Dynamic Linker
  * /usr/lib/libSystem.B.dylib: system's standard C library (libc), the POSIX thread library (libpthread),
                                and other UNIX system calls and standard utility functions.
  * /usr/lib/libresolv.9.dylib: system library used for resolver routines such as those used for DNS lookups.
* At 00001000, we see the Go BuildID: "n7a7McM3-_0HfDdXKYJy/985rFS1mhQ-RRiRIR0NL/JZ4BJT8lkCkZDkkUzNjP/xMHCwCDIIrFU6JeAW9aF"
*  
## We disassemble our target (`hijack-test`)
```shell
go tool objdump -S hijack-test > hijack-test.asm.detailed
```
* Search the `hijack-test.asm.detailed` file for `maing.go` to find our program entrypoint (comments added 
  during analysis):

### main() function:

<pre style="color:grey">
TEXT main.main(SB) /Users/samcaldwell/git/go/cmd/debugging/main.go
func main() {
</pre>
<pre>
  0x10973a0             4c8da42498fcffff        LEAQ 0xfffffc98(SP), R12
        ; Calculate effective address from stack pointer 

  0x10973a8             4d3b6610                CMPQ 0x10(R14), R12
        ; Compare 64-bit int at address offset 0x10 from R14 with the value in R12.      

  0x10973ac             0f865a150000            JBE 0x109890c
        ; Jump to address if zero or carry flag are set.      

  0x10973b2             4881ece8030000          SUBQ $0x3e8, SP
        ; Subtract 0x3e8 from SP, allocating space on the stack for local variables and saved registers.

  0x10973b9             4889ac24e0030000        MOVQ BP, 0x3e0(SP)
        ; Move base pointer (BP) into allocated stack space at offset 0x3e0.

  0x10973c1             488dac24e0030000        LEAQ 0x3e0(SP), BP
        ; Update base pointer (BP) to point to new base of the stack after the previously made adjustment.
</pre>
<pre style="color:grey">
    var actualValue = words.EmptyString
</pre>
<pre>
  0x10973c9             488d0510980000          LEAQ runtime.rodata+32832(SB), AX
        ; Load effective address of runtime.rodata+32832 into the AX register. runtime.rodata is a
        ; section in the Go binary that contains read-only data.   

  0x10973d0             e80b4af7ff              CALL runtime.newobject(SB)
        ; Call Go runtime function newobject, which is responsible for allocating new objects. The
        ; argument to newobject is the type of the object to be allocated, which is supplied in the AX
        ; register.

  0x10973d5             48898424c8000000        MOVQ AX, 0xc8(SP)   
        ; Moving the value of the AX register (which now holds the pointer to the newly
        ; allocated object) to the memory address at an offset of 0xc8 from the current stack pointer
        ; SP. In essence, it's storing the address of the new object on the stack.

  0x10973dd             48c70000000000          MOVQ $0x0, 0(AX)
        ; Finally, this line is writing a 0 (or NULL) to the memory pointed to by the AX register,
        ; effectively initializing the new object's memory to 0. This is often done to ensure that the new
        ; object starts in a known state.
</pre>
<pre style="color:grey">
        targetFunc := func() {
</pre>
<pre>
  0x10973e4             488d05b5e90000          LEAQ runtime.rodata+53760(SB), AX       
  0x10973eb             e8f049f7ff              CALL runtime.newobject(SB)              
  0x10973f0             48898424c0000000        MOVQ AX, 0xc0(SP)                       
  0x10973f8             488d0dc1150000          LEAQ main.main.func1(SB), CX            
  0x10973ff             488908                  MOVQ CX, 0(AX)                          
  0x1097402             833d17ee0e0000          CMPL $0x0, runtime.writeBarrier(SB)     
  0x1097409             750e                    JNE 0x1097419                           
  0x109740b             488b8c24c8000000        MOVQ 0xc8(SP), CX                       
  0x1097413             48894808                MOVQ CX, 0x8(AX)                        
  0x1097417             eb11                    JMP 0x109742a                           
  0x1097419             488d7808                LEAQ 0x8(AX), DI                        
  0x109741d             488b8c24c8000000        MOVQ 0xc8(SP), CX                       
  0x1097425             e89668fcff              CALL runtime.gcWriteBarrierCX(SB)
</pre>
<pre style="color:grey">
        imposterFunc := func() {
</pre>
<pre>
  0x109742a             488d056fe90000          LEAQ runtime.rodata+53760(SB), AX       
  0x1097431             e8aa49f7ff              CALL runtime.newobject(SB)              
  0x1097436             488d0de3140000          LEAQ main.main.func2(SB), CX            
  0x109743d             488908                  MOVQ CX, 0(AX)                          
  0x1097440             833dd9ed0e0000          CMPL $0x0, runtime.writeBarrier(SB)     
  0x1097447             750e                    JNE 0x1097457                           
  0x1097449             488bbc24c8000000        MOVQ 0xc8(SP), DI                       
  0x1097451             48897808                MOVQ DI, 0x8(AX)                        
  0x1097455             eb14                    JMP 0x109746b                           
  0x1097457             488d7808                LEAQ 0x8(AX), DI                        
  0x109745b             488b9424c8000000        MOVQ 0xc8(SP), DX                       
  0x1097463             e87868fcff              CALL runtime.gcWriteBarrierDX(SB)     
</pre>
<pre style="color:grey">
        if actualValue != words.EmptyString {
</pre>
<pre>
  0x1097468             4889d7                  MOVQ DX, DI             
</pre>

### hackFunction() function:
<pre style="color:grey">
TEXT github.com/sam-caldwell/go/v2/projects/testing/hijack.hijackFunction(SB) /Users/samcaldwell/git/go/projects/testing/hijack/hijackFunction.go
func hijackFunction(targetFuncPtr, imposterFuncPtr uintptr) (originalMemory []byte, err error) {
</pre>
<pre>
  0x1096380             4c8da424c0fdffff        LEAQ 0xfffffdc0(SP), R12        
  0x1096388             4d3b6610                CMPQ 0x10(R14), R12             
  0x109638c             0f86160c0000            JBE 0x1096fa8                   
  0x1096392             4881ecc0020000          SUBQ $0x2c0, SP                 
  0x1096399             4889ac24b8020000        MOVQ BP, 0x2b8(SP)              
  0x10963a1             488dac24b8020000        LEAQ 0x2b8(SP), BP              
  0x10963a9             48895c2470              MOVQ BX, 0x70(SP)               
  0x10963ae             4889442468              MOVQ AX, 0x68(SP)               
        ansi.Blue().Println("hijackFunction() start").
  0x10963b3             e808f9ffff              CALL github.com/sam-caldwell/go/v2/projects/ansi.Blue(SB)       
        fmt.Println(message)
</pre>

