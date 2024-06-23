/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 * @brief Implement a go-like message queue for passing data
 * between threads safely.  This implements both a
 * boundless and bounded (buffered) Channel implementation
 * with an iterator for use in looping.
 */

#ifndef Channel_H
#define Channel_H

#include "../../exceptions/exceptions.h"
#include "../../../common/SimpleLock/src/SimpleLock.h"
#include <queue>

#include <queue>
#include <mutex>
#include <condition_variable>
#include <iostream>

template <typename T>
class Channel {
public:
 Channel() : closed(false) {}

 ~Channel() {
  close(); // Ensure channel is closed and resources are cleaned up
 }

 void operator<<(const T& value) {
  std::unique_lock<std::mutex> lock(mtx);
  if (closed) {
   throw std::runtime_error("Channel closed");
  }
  queue.push(value);
  cv.notify_one();
 }

 bool operator>>(T& value) {
  std::unique_lock<std::mutex> lock(mtx);
  cv.wait(lock, [this] { return closed || !queue.empty(); });
  if (!queue.empty()) {
   value = queue.front();
   queue.pop();
   return true;
  }
  return false;
 }

 void close() {
  std::unique_lock<std::mutex> lock(mtx);
  closed = true;
  cv.notify_all();
 }

 bool is_closed() const {
  std::unique_lock<std::mutex> lock(mtx);
  return closed;
 }

private:
 std::queue<T> queue;
 mutable std::mutex mtx;  // mutable to allow locking in const member function
 std::condition_variable cv;
 bool closed;
};

#endif  /** Channel_H */