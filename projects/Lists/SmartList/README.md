SmartList
=========
----
## Description

An arbitrary type list struct with methods to make it easier to use.

## Usage:
### Important

1. You can Add() any type of data supported by `any`, but...
2. Once you add the first record, every record afterward must be of the same type.
3. Only if you empty the `SmartList` of all records can you change type.

### Declaration

```text
var list SmartList
```

### Multi-type Support
```text
var listBool SmartList
list.Add(true)
list.Add(true)
list.Add(false)
list.Add(true)

var listInt SmartList
list.Add(1)
list.Add(5)
list.Add(8)
list.Add(1)
```

### Add Records: `.Add()`
The `.Add()` method will typecheck your input and if ok, append it to the list
of records in the list:
```text
list.Add("foo")  // Add record to list (and set the list to be strings only.
list.Add("bar")
list.Add("moo")
list.Add("bar")
```
But if you attempt to add a value that is not the same type as `list[0]` an
error will be raised.
```text
list.Add(true) // will return `exit.ErrTypeMismatch` because true is bool
               // and the list contains strings.
```

### TypeCheck Methods: `.TypeCheck()`
The `.Add()` method uses `.TypeCheck()` to verify that a new record is of the
same type as `list[0]`.  If this is not the case, an `exit.ErrTypeMismatch` 
error is returned.
```text
list.TypeCheck(1) // will return `exit.ErrTypeMismatch` because 1 is int 
                  // and the list contains strings.
```

### String Method: `.String()`
The contents of `SmartList` can be returned as `[]string` using the `.String()` 
method.
```text
list.String() //Returns []string{"foo","bar","moo","mar"}
```
