#include <iostream>
#include <cstdlib>

using namespace std;

const int hashSize = 256;
const int headerSize = 160;
const int blockIdSz = 64;

uint64_t calculateMinimum(uint64_t start, uint64_t stop, int crsceAlgorithm) {
    // Minimum signal size
    uint64_t Ms = 0;
    // Hash size (based on CRSCE variant)
    uint64_t h = hashSize;
    if (crsceAlgorithm == 2) {
        h *= 2; // increase h to 2 x hashSize (LHASH+VHASH) assuming CRSCE-2
    }
    for (uint64_t N = start; N < stop; N++) {

        // Cross Sum Size
        uint64_t n =static_cast<int64_t>( std::ceil( std::sqrt(8 * N) ) );

        // Cross Sum Width
        uint64_t b = static_cast<int64_t>( std::ceil( std::log2(n) ) );

        // Compressed Signal Size
        uint64_t numerator = 2 * b * n + h * n + blockIdSz + headerSize;

        // Csm Array Size.
        uint64_t denominator = n * n;

        // Padding size
        uint64_t P = ( n * n - 8 * N);

        if ( (P < n) && (numerator < denominator) ) {
            Ms = N;
            return Ms;
        }
    }
    throw std::runtime_error("ms cannot be determined");
}

int main(int argc, char *argv[]) {
    if (argc != 2) {
        std::cerr << "Usage: " << argv[0] << " CRSCE_ALGORITHM (1|2)" << std::endl;
        return 1;
    }

    int crsceAlgorithm = std::stoi(argv[1]);
    if (crsceAlgorithm != 1 && crsceAlgorithm != 2) {
        std::cerr << "Invalid input. Expects 1 or 2" << std::endl;
        return 1;
    }

    uint64_t startSz = 1024;        // Upper search bounds
    uint64_t stopSize = 1048576; // Lower search bounds

    try {
        uint64_t Ms = calculateMinimum(startSz, stopSize, crsceAlgorithm);
        uint64_t minimumSignalSizeInBits = Ms * 8;
        uint64_t minimumSignalSizeInBytes = Ms;
        std::cout << "  minimum signal size(bytes) (Ms): " << minimumSignalSizeInBytes << std::endl;
        std::cout << "        minimum signal size(bits): " << minimumSignalSizeInBits << std::endl;
    } catch (const std::runtime_error &e) {
        std::cerr << "Error: " << e.what() << std::endl;
    }

    return 0;
}
