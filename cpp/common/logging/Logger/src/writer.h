/**
 * @name Logger/src/writer.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef Logging_Writer_h
#define Logging_Writer_h

#include <map>
#include <sstream>
#include <string>
#include <fstream>
#include "../../../../system/file.h"

namespace Logging {
    namespace Writer {
        enum Id : unsigned char {
            null = 0,
            stdout = 1,
            stderr = 2,
            file = 3,
            syslog = 4,
            https = 5
        };

        //table for writer lookups.
        const static map <string, Logging::Writer::Id> StringsTable = {
                {"null",   Logging::Writer::Id::null},
                {"stdout", Logging::Writer::Id::stdout},
                {"stderr", Logging::Writer::Id::stderr},
                {"file",   Logging::Writer::Id::file},
                {"syslog", Logging::Writer::Id::syslog},
                {"https",  Logging::Writer::Id::https}
        };
    }
}

/**
 * @name writerBase
 * @brief This is the base class for the writer classes used in Logger Context.
 */
class writerBase {
protected:
    string destination;
public:
    writerBase() {/*noop*/}

    ~writerBase() {/*noop*/}

    virtual void write(stringstream *msg) {/*noop*/}

    virtual const string *encode(const string *msg) { return msg; }

    string getDestination() { return destination; }
};

/**
 *
 */
class writerFile : public writerBase {
public:
    writerFile() = delete;

    writerFile(string *d) {
        destination = *d;
    }

    ~writerFile() {
        //handle.close();
        cout << "writerFile destructor" << endl;
    }

    void write(stringstream *msg) {
        std::fstream handle;
        handle.open(destination, std::fstream::out | std::fstream::app);
        if (handle.is_open())
            handle << msg->str() << endl;
        else
            throw runtime_error("failed to open destination file: " + destination);
        handle.flush();
        handle.close();
    }

    string getDestination() { return destination; }
};


/**
 *
 */
class writerHttps : public writerBase {
public:
    writerHttps() = delete;

    writerHttps(string *d) {
        destination = *d;
        //ToDo: connect to the server
    }

    ~writerHttps() {
        //handle.close();
    }

    void write(stringstream *msg) {
        //handle << msg->str() << endl;
    }

    string getDestination() { return destination; }
};


/**
 *
 */
class writerNull : public writerBase {
public:
    writerNull() {/*noop*/}

    writerNull(string *destination) {/*noop*/}

    ~writerNull() {/*noop*/}

    void write(stringstream *msg) {/*noop*/};
};


/**
 *
 */
class writerStderr : public writerBase {
private:
public:
    writerStderr() {/*noop*/}

    ~writerStderr() {}

    void write(stringstream *msg) {
        cerr << msg->str() << endl;
    }
};


/**
 *
 */
class writerStdout : public writerBase {
private:
public:
    writerStdout() {/*noop*/}

    ~writerStdout() {}

    void write(stringstream *msg) {
        cout << msg->str() << endl;
    }
};

/**
 *
 */
class writerSyslog : public writerBase {
public:
    writerSyslog() = delete;

    writerSyslog(string *d) {
        destination = *d;
        //ToDo: connect to the server
    }

    ~writerSyslog() {
        //handle.close();
    }

    void write(stringstream *msg) {
        //handle << msg->str() << endl;
    }

    string getDestination() { return destination; }
};

#endif /* Logging_Writer_h */