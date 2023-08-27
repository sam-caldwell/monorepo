/*
 * DataSource::Network::destructor
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */
#include "Network.h"
/*
 *
 */
DataSource::Network::~Network(){
    close(socketFd);
}