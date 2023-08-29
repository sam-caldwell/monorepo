#include "File.h"
/*
 *
 */
DataSource::File::~File() {
    file.close();
}