#include "File.h"
/*
 *
 */
DataSource::File::File(const std::string &filename) {
    file.open(filename, std::ios::binary | std::ios::out | std::ios::in);
    if (!file.is_open()) {
        throw std::runtime_error("File not found or couldn't be opened");
    }
}