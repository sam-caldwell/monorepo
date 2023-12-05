EnvironmentVariableParser
=========================

Given a ConfigStateMap object defining
a set of environment variables, internal
configuration names, their types, error
messages and validator functions, this
class will iterate over the map to parse
current environment variable values, then
validate them before storing them in an
ArbitraryKvList linked list structure.

The consumer can then consume the arbitrary
but strongly typed environment variables
using its begin(), end() and next() methods.

