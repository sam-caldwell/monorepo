/*
 *
 */
#ifndef BITREADER_H
#define BITREADER_H

#include "../BackgroundProcess/BackgroundProcess.h"
#include "../DataSource/DataSource.h"
#include <vector>

const int BitReaderQueueSize = 128;

class BitReader : public BackgroundProcess {
private:
    BackgroundProcess asynchronousReader;
    DataSource *source;
    byte_t currentByte;
    uint8_t currentBit;
    std::vector<byte_t> queue;
    bool byteAlign;

    void AsynchronousReader();

public:
    BitReader(DataSource *s, bool byteAlign = false);

    ~BitReader();

    byte_t pop();
};

#endif

