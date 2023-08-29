Counter
=======

## Description

A set of counters which helps avoid reinvented wheels.

## Usage
```golang
package main

import counters "github.com/sam-caldwell/monorepo/v2/projects/counter"

var count counters.Simple
count.Increment() //output: 0
count.Increment() //output: 1
count.Increment() //output: 2
count.Value(0) //output: 2

var countIf counters.Conditional
countif.Increment(true) //output: 0
countif.Value()          //output: 1
countif.Increment(false) //output: 1
countif.Decrement(true) //output:1

var count = CreateFunc(1, 1)
fmt.Println(count()) // prints 2
fmt.Println(count()) // prints 3
fmt.Println(count()) // prints 4

count = CreateFunc(10, -1)
fmt.Println(count()) // prints 9
fmt.Println(count()) // prints 8
fmt.Println(count()) // prints 7

```

`countOk(ok bool)` - returns incremented count if input is `true` or `-1` if false
