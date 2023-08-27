/*
 * DataSource::File::Read
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */
#include "File.h"
/*
 * Read a byte from the given position
 */
byte_t DataSource::File::Read(uint64_t position) {
    byte_t data;
    file.seekg(position);
    file.read(reinterpret_cast<char *>(&data), sizeof(byte_t));
    return data;
}

/*
 * Read a byte from the current position
 */
byte_t DataSource::File::Read() {
    byte_t data;
    file.read(reinterpret_cast<char *>(&data), sizeof(byte_t));
    return data;
}
