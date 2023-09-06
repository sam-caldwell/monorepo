/*
 * BitReader Class Constructor
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */
#include "BitReader.h"
BitReader::BitReader(DataSource *s, bool byteAligned) {
    BackgroundProcess(AsynchronousReader);
    byteAlign = byteAligned;
    currentBit = 0;
    currentByte = 0x00;
    queue.reserve(BitReaderQueueSize);
    source = s;
};