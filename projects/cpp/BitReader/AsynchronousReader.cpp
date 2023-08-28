/*
 * BitReader Asynchronous Reader
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */
#include "BitReader.h"
void BitReader::AsynchronousReader() {
    while(!source->eof() && ( queue.size() < queue.capacity() ) ) {
        queue.push_back(source->Read());
    }
}