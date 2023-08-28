/*
 * BitReader Class Destructor
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */
#include "BitReader.h"
BitReader::~BitReader(){
    delete source;
    queue.clear();
}