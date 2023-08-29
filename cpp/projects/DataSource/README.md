DataSource
==========

## Description
A DataSource is an abstraction layer between files, network connections and other data sources.


```c++

#include "DataSource.h"

int main() {
    try {
        DataSource::Root<DataSource::File> fileDataSource("data.bin");
        DataSource::Root<DataSource::Network> networkDataSource("192.168.1.100", 8080);

        // Use the data sources for read/write operations
        byte_t dataFromFile = fileDataSource.Read(0);
        byte_t dataFromNetwork = networkDataSource.Read(0);

        // Perform write operations if needed
    } catch (const exception& e) {
        cout << "Error: " << e.what() << endl;
    }

    return 0;
}

```