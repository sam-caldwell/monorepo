/**
 * @name projects/common/Breadcrumbs/Breadcrumbs.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef Breadcrumbs_H
#define Breadcrumbs_H

#include <string>
#include <vector>
#include <type_traits>
#include "../../../common/formatter/formatter.h"

using namespace std;

template<class T>
constexpr T default_value = {};
template<> constexpr char default_value<char> = 0x00;
template<> constexpr double default_value<double> = 0;
template<> constexpr float default_value<float> = 0;
template<> constexpr int default_value<int> = 0;
template<> string default_value<string> = "";
template<> constexpr unsigned default_value<unsigned> = 0;


/**
 * @name Breadcrumbs<T>
 * @brief The "Breadcrumbs" class allows a user to collect a set of elements of type T,
 *        to then pop from that collection, view the last element added or to dump the
 *        entire collection as a dot-delimited string.
 *
 * @tparam T (a type)
 */
template<class T>
class Breadcrumbs {
    static_assert(std::is_same<T, double>::value
    ||
    std::is_same<T, float>::value ||
    std::is_same<T, int>::value ||
    std::is_same<T, unsigned>::value ||
    std::is_same<T, char>::value ||
    std::is_same<T, string>::value,
    "Only double, float, int, unsigned, char and string types are allowed.");
public:
    Breadcrumbs(char _delim = '.') { delim = _delim; };

    Breadcrumbs(T s, char _delim = '.') {
        delim = _delim;
        data.push_back(s);
    }

    ~Breadcrumbs() {}

    inline void delimiter(char c = ' ') {
        delim = c;
    }

    inline bool push(T s) {
        data.push_back(s);
        return true;
    }

    inline bool pop() {
        if(data.size()>0)
            data.pop_back();
        return true;
    }

    inline bool clear() {
        data.clear();
        return true;
    }

    inline int count() {
        return data.size();
    }

    inline T current() {
        if(data.size() > 0)
            return data.back();
        else
            return default_value<T>;
    }

    inline int empty() {
        return data.size() == 0;
    }

    inline T prev() {
        return (data.size() >= 2) ? data.at(data.size() - 2) : default_value<T>;
    }

    inline string toString() {
        stringstream buffer;
        //for (vector<T>::iterator it = data.begin(); it != data.end(); ++it)
        for (int i = 0; i < data.size(); i++)
            buffer << data.at(i) << delim;
        return Strings::trim(Strings::rtrimChar(buffer.str(), delim));
    }

private:
    char delim;
    vector <T> data;
};

#endif  /* Breadcrumbs_H */
