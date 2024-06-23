/**
 * @name networking/TcpCommon/src/client/equalityOperator.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
namespace networking {
    /**
     * @name operator ==
     * @brief Create an operator to test lhs==rhs
     * @param other Client reference
     * @return bool
     */
    bool Client::operator==(const Client &other) const {
        if ((this->SocketFileDescriptor.get() == other.SocketFileDescriptor.get()) &&
            (this->_ip == other._ip)) {
            return true;
        }
        return false;
    }

    /**
     * @name operator !=
     * @brief Create an operator to test lhs!=rhs
     * @param other Client reference
     * @return bool
     */
    bool Client::operator!=(const Client &other) const {
        if ((this->SocketFileDescriptor.get() != other.SocketFileDescriptor.get()) ||
            (this->_ip != other._ip)) {
            return true;
        }
        return false;
    }
}