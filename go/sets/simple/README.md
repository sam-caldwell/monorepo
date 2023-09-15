simpleSet
=========

## Description

A simple golang set because who wants to reinvent wheels all the time
when you really just want to solve problems?

## Methods

| Method       | Syntax                                     | Description                                                  |
|--------------|--------------------------------------------|--------------------------------------------------------------|
| Add()        | `.Add(item any) error`                     | Add a item to set. Throw error type check fails              |
| Count()      | `.Count() int`                             | Return a count of records                                    |
| Delete()     | `.Delete(item any) error`                  | Delete an existing record                                    |
| Empty()      | `.Empty() bool`                            | Returns boolean true if set empty or set is uninitialized    |
| GetFirst()   | `.GetFirst() any`                          | Returns the first record in the set (not always first added) |
| GetType()    | `.GetType() reflect.Kind`                  | Returns the set type                                         |
| Has()        | `.Has(item any) bool`                      | Return true if item exists or false otherwise.               |
| Init()       | `.Init()`                                  | Initialize the state's internal data structure               |
| List()       | `.List() (result []any)`                   | Return a list of items                                       |
| ListString() | `.ListString(sort bool) (result []string)` | Return the set as a list of strings.                         |
| TypeCheck()  | `.TypeCheck(item any) bool`                | Return bool true if set empty or of same type                |
