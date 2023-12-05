fs/file/BitFile 
===============

## Objectives
* Define a bitwise file reader/writer

## Features
1. `BitFile` structure allows for a buffered reader/writer capable of byte-level or bit-level read/write operations.
2. `BitFile.Close()` closes any open file.
3. `BitFile.Create()` will create a new text file (if not found or open an existing file).
4. `BitFile.Name()` returns the name of an open file or "" if no file is open.
5. `BitFile.Open()` will open a file if it exists.
6. `BitFile.Read()` will read a bit 
7. 