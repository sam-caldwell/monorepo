/**
 * @name types/Mapper.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef Mapper_H
#define Mapper_H

#import <map>

using namespace std;

template<typename KeyType, typename ValueType>
class Mapper {
public:
    Mapper() {}

    ~Mapper() {}

    /**
     * @name push
     * @brief push a key-value pair to the map.
     *
     * @param key <KeyType>
     * @param value <ValueType>
     * @return bool
     */
    bool push(KeyType key, ValueType value) {
        data[key] = value;
        lastKey = key;
        lastValue = value;
        return true;
    }

    /**
     * @name pop
     * @brief pop the given key-value element off the map.
     *
     * @param key <KeyType>
     * @param value <ValueType>
     * @return bool
     */
    bool pop(KeyType key) {
        if (data.find(key) != data.end()) {
            data.erase(key);
            return true;
        } else
            return false;
    }

    /**
     * @name current
     * @brief return the last-pushed value type.
     *
     * @return <ValueType>
     */
    inline ValueType current() {
        return lastValue;
    }

    /**
     * @name currentKey
     * @brief return the last-pushed key
     *
     * @return <KeyType>
     */
    inline KeyType currentKey() {
        return lastKey;
    }

    /**
     * @name count
     * @brief return the count
     *
     * @return unsigned
     */
    unsigned count() { return data.size(); }

    /**
     * @name empty
     * @brief return whether or not the class is empty.
     *
     * @return bool
     */
    inline bool empty() { return data.size() == 0; }

    /**
     * @name clear
     * @brief delete all data in the class.
     *
     * @return bool
     */
    inline bool clear() {
        data.clear();
        return true;
    }

    /**
     * @name get
     * @brief Return the value associated with the given key.
     *
     * @param key <KeyType>
     * @return <ValueType>
     */
    unsigned get(KeyType key) {
        if (data.find(key) != data.end()) {
            return data[key];
        } else {
            throw out_of_range("key " + key + " not found");
        }
    }

private:
    KeyType lastKey;
    ValueType lastValue;
    map <KeyType, ValueType> data;
};

#endif /**Mapper_H*/
