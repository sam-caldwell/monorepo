/*
 * DataSource::File
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */
#ifndef DATASOURCE_FILE_H
#define DATASOURCE_FILE_H

#include <iostream>
#include <fstream>
#include <cstring>
#include <stdexcept>
#include <unistd.h>
#include "../../types/byte_t.h"
#include "../DataSource.h"
/*
 * Specialization for file-based data source
 */
class DataSourceFile: public DataSource {
private:
    std::fstream file;

public:
    DataSourceFile(const std::string &filename);

    ~DataSourceFile();

    byte_t Read();

    byte_t Read(uint64_t position);

    byte_t Write(byte_t content);

    byte_t Write(uint64_t position, byte_t content);

    uint64_t Size();
};


#include "File.cpp"

#endif