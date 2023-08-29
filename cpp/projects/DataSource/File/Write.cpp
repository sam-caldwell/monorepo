/*
 * DataSource::File::Write
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */
#include "File.h"
/*
 * Write a given byte at the specified position
 */
byte_t DataSource::File::Write(uint64_t position, byte_t content) {
    file.seekg(position);
    file.write(reinterpret_cast<char *>(&content), sizeof(byte_t));
    return content;
}
/*
 * Write a byte to the current position.
 */
byte_t DataSource::File::Write(byte_t content) {
    file.write(reinterpret_cast<char *>(&content), sizeof(byte_t));
    return content;
}