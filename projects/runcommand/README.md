RunCommand
==========

## Description

A command-execution facility in golang

## Usage

The RunCommand package provides an abstracted way of executing system commands within
Go code. It consists of a `CommandExecutor` interface and two implementing types:

* `RealCommandExecutor` and
* `MockCommandExecutor`.

The CommandExecutor interface defines a method Execute which accepts the name of
the command to be run, along with its arguments, and returns the output of the
command as a byte slice or an error.

### RealCommandExecutor

The `RealCommandExecutor` type is a concrete implementation of the `CommandExecutor`
interface. This executor runs the provided command with its arguments and captures
its output.

Here is an example of using `RealCommandExecutor`:

```
executor := runcommand.RealCommandExecutor{}
output, err := executor.Execute("echo", "hello")
    if err != nil {
    // Handle error
}
fmt.Println(string(output)) // Prints: hello
```

### MockCommandExecutor

The `MockCommandExecutor` type is another implementation of `CommandExecutor` but with a
difference. It doesn't actually execute any command. Instead, it returns predefined
output and error which you can set at the time of creation. This executor is
particularly useful when writing tests for your functions that depend on external
commands. It allows you to define expected command output and error, and thereby
test your function's behavior in a controlled way.

Here is an example of using `MockCommandExecutor` in a test:

```text
func TestYourFunction(t *testing.T) {
	mock := runcommand.MockCommandExecutor{
		Output: "expected output",
		Error:  nil,
	}
	
	// Pass the mock to your function
	result := YourFunction(mock)
	
	// Check that the function behaves as expected
}

```

By abstracting the execution of commands, the `RunCommand` package allows you to write
code that is easier to test and maintain. It provides clear separation between your
application logic and the external command execution. In your tests, it provides a way
to simulate different command outputs and errors, so you can ensure your code behaves
correctly in different situations.
