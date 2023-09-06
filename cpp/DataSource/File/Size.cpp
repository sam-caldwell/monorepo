/*
 * DataSource::File::Size
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */
#include "File.h"
/*
 *
 */
uint64_t DataSource::File::Size() {
    // Get the current position
    std::streampos currentPosition = file.tellg();

    // Move to the end of the file to determine its size
    file.seekg(0, std::ios::end);
    uint64_t fileSize = file.tellg();

    // Move back to the original position
    file.seekg(currentPosition);

    return fileSize;
}
