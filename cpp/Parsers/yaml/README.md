Parser
======

`Parser` is an MVP of a grand-unified file format parser. The current initiative
is to parse YAML files into key-value data which can be consumed by `ConfigFileParser`
or other projects. We chose YAML for our first file format since it allows comments
(which unfortunately JSON does not) and it is not as verbose as XML.

We evaluated `yaml-cpp` for this project and decided it was not appropriate for our
use-case and did not appear to have a sufficient and active community around it.

### Goals

- Avoid external dependencies
- Allow for very minimal implementations when appropriate.
- Allow for eventual support of JSON and possibly XML if I ever feel it is needed.
- ***DO NOT USE RECURSION*** to support processing of large files.
- Allow for stream support in the future.

### ToDos:

- Use C++ preprocessor directives to allow a consuming application to minimize
  what features are enabled and built into an application.
- Implement stream support.

