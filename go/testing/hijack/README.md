Testing/Hijack
==============
 <h1 class="blinking-text" style="color:yellow;padding:2px 2px 2px 2px;margins:1px 1px 1px 1px">This is a work in progress</h1>
## Description



This project creates a 'monkey patch' solution of sorts. It's ugly, unsafe and generally something
to be avoided at all costs except in rare cases when writing tests.

<style>
@keyframes blink {
  0% { opacity: 1; }
  50% { opacity: 0; }
  100% { opacity: 1; }
}

.blinking-text {
  animation: blink 1s infinite;
}
</style>

<div style="border-color:red;border-width: 2px;border-style: outset;padding:2px 2px 2px 2px;margins:1px 1px 1px 1px">
    <h1 class="blinking-text" style="color:red;padding:2px 2px 2px 2px;margins:1px 1px 1px 1px">WARNING:</h1>
    <p style="color:red">
        <i>DO NOT USE THIS IN PRODUCTION!</i>  
        ...and be sparing when using this in tests.</p>
    <p style="color:red">
        Be very suspicious of this code.  This code does some really low level stuff to make 
        it possible to 'monkey patch' golang.  This means it violates a few of the rules about
        direct memory stuff that could cause crashes and such.  If you ignore this and use this
        in production, do not report any errors or outages.  You were warned.
    </p>
</div>

## What this actually does...

1. This is VERY unsafe. Did I say <span style='color:red'>"Do not use this in production?"</span>
   You are likely to crash your app with this tool.

2. This project will allow you to edit the golang binary ***DURING RUNTIME*** by finding and overwriting functions
   and methods ***IN MEMORY***.

3. There is a "restore" capability, and it is intended to rollback changes. But Golang is doing things in memory
   while you are too. So...eh? It may break.

4. We try to keep the golang garbage collector happy by keeping track of references to things, but...eh? It may break.

## Usage

### A Quick Example...

The following is an example of how to use the `testing/hijack` package:

```golang
package main

import (
   "fmt"
   "github.com/sam-caldwell/monorepo//projects/testing/hijack"
   "github.com/sam-caldwell/monorepo//projects/wrappers/os"
)

func main() {
   /*
    * We create a hijack.TrackerTable to track all our memory objects.
    * This allows us to restore a change. But it also means we keep a
    * reference to the objects in memory which should theoretically keep
    * the garbage collector from freeing an object and creating chaos.
    */
	
   myHijacker := hijack.NewTrackerTable()
   
   /*
    * Now that we have a TrackerTable we can use the .Capture() method
    * to hijack a function.  
    * 
    * Here, we takeover the fmt.Println() function used to write to the 
    * screen and ignore the user's input and simply write "Everything is
    * fine" to the screen.
    */
   
   err := myHijacker.Capture(fmt.Println, func(msg ...any) (n int, err error) {
      fmt.FPrintln(os.Stdout, "Everything is fine")
      return 0, nil
   })
   
   /*
    * If we had any error in hijacking the function, we can handle it below.
    */
   
   if err != nil {
      panic(fmt.Sprintf("We failed to hijack fmt.Println. %v", err))
   }
   
   /*
    * Now, when we call fmt.Println, we are actually calling the
    * hijacked version of fmt.Println and our output will be 'Everything is fine'
    */
   fmt.Println("Call Chuck Norris, we've been hijacked!")
   /*
    * When we are done, we can call the .Restore() method and handle any error.
    * Assuming there are no errors, fmt.Println will be restored to a working
    * version.
    */
   if err := myHijacker.Free(fmt.Println); err != nil {
      panic(fmt.Sprintf("restore failed: %v", err))
   }
   /*
    * Now, when we call fmt.Println, it's all working as expected.
    */
   fmt.Println("It's all back to normal.")

}
```

### Methods and Functions

### hijack.NewTrackerTable()

> > Creates, initializes and returns a `GlobalPatchTable`. Please use this. Don't try doing it yourself.

### hijack.Capture(targetFunc, imposterFunc) error

> > Given a Target Function (the thing you want to take over) and Imposter Function (the code you want instead),
> > This method will locate, copy and overwrite the `targetFunc` with the `imposterFunc`.
>>
>> The change will be tracked, so it can be reverted...theoretically.
> > <span style="color:red">Do not use this with a method, only with functions. For methods, use .CaptureMethod()</span>

### hijack.Free(targetFunc) error

> > Given a target function (the thing you took over with .Capture()), this method will revert the change using the
> > TrackerTable and restore original functionality...theoretically. <span style="color:red">Do not use this with a
> > method, only with functions. For methods, use .FreeMethod()</span>

### hijack.CaptureMethod(targetObject, methodName, imposterFunc) error

> > Given a targetObject (struct) and methodName (string), this method will look up the methodName within the
> > targetObject and patch that method with the given imposterFunc specified.

### hijack.FreeMethod(targetObject, methodName) error

> > Given a targetObject (struct) and methodName (string), this method will look up and restore the functionality of
> > the given method. As all good programmers would say, this all works great! `***`Pats hood`***`  you'll get lots of
> > mileage from this baby!

### hijack.FreeAll() error

> > Where hijack.Free() will revert a specific Captured function, this method will roll back all changes currently known
> > to the TrackerTable.
