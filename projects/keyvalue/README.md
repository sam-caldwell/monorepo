KeyValue Map
============

## Description

A struct and methods for working with maps of key-value data
that doesn't involve redoing the same efforts over and over,
repeating yourself again and again...endlessly. :-)

## Features

[KeyValue](./KeyValue.go)

| Feature                                          | Description                                                         |
|--------------------------------------------------|---------------------------------------------------------------------|
| [`KeyValue`](./KeyValue.go)                      | Generic struct which stores a collection of key-value pairs.        |
| [`.Set()`](./Set.go)                             | Add a new key-value pair to the `KeyValue` collection.              |
| [`.Get()`](./Get.go)                             | Return a value for a given `key` from the `KeyValue` collection.    |
| [`.Exists()`](./Exists.go)                       | Return `true` if the given `key` exists in `KeyValue` collection.   |
| [`.FindKey()`](./FindKey.go)                     | Return the `KeyValue` `Pair` for a given `key`                      |
| [`.FindValue()`](./FindValue.go)                 | Return the `KeyValue` `Pair` with a matching `value`                |
| [`.FromBytes()`](./FromBytes.go)                 | Parse and import a `[]byte` input into the `KeyValue` collection.   |
| [`.FromCommandShell()](./FromCommandShell.go)    | Parse and import `[]byte` output from a command shell execution.    |
| [`.FromFile()](./FromFile.go)                    | Read, parse, import a file into the `KeyValue` collection.          |
| [`.FromString()`](./FromString.go)               | Read, parse, import a string into `KeyValue` collection.            |
| [`.Initialize()`](./Initialize.go)               | Initialize the state of the `KeyValue` collection.                  | 
| [`.KeyWidth()`](./KeyWidth.go)                   | Return the maximum key width for the `KeyValue` collection state.   |
| [`.MergeFromKv()`](./MergeFromKv.go)             | Merge two `KeyValue` collections.                                   |
| [`.RenameKey()`](./RenameKey.go)                 | Rename a given `KeyValue` collection record key.                    |
| [`.ToOrderedPairList()`](./ToOrderedPairList.go) | Convert a `KeyValue` collection into an `OrderedPair` list.         |
| [`.ToString()`](./ToString.go)                   | Stringify the `KeyValue` internal state.                            |
| [`.ToStringArray()`](./ToStringArray.go)         | Return a list of strings, where each string is a key-value pairing. |
| [`.ValueWidth()`](./ValueWidth.go)               | Return the maximum value width for the `KeyValue` collection.       |
| [`.Walk()`](./Walk.go)                           | Walk the `KeyValue` collection with a given walker function.        |
| [`.WriteFile()`](./WriteFile.go)                 | Write the `KeyValue` collection to a file.                          |
