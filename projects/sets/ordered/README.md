orderedSet
==========

## Description

An ordered set of any type. Given a set of n elements, where
all n elements are of the same type, this structure will ensure
the ordering is unchanged and represents the order in which the
elements were received.

* Elements are stored in order of storage
* Elements are addressable by index.
* Elements are guaranteed unique
* Elements are of the same type

## Methods

### `.Add(item any) error`

> Add a new item to the set and throw an error if it already exists.

### `.Count() int`

> Return a count of records

### `.Delete(pos int) error`

> Delete an existing

### `.List() []any`

> Return a list of elements `[]any` from the given set

### `.Pop() any`

> Pop element 0 from the set and return its value

### `.Push(item any)`

> Push an element to the top of the set.
