CommandLineParser
=================

This project creates a class for parsing command-line arguments
into a standardized `Configuration` class instance as part of the
`Application` core framework.

## Usage

### Caveat

Rather than using this standalone, please use
`projects/Application` to start your application within the
repository's larger framework.

### Direct Usage

1. Define a `ConfigStateMap`

```c++
ConfigStateMap *map;
map = new ConfigStateMap;
```

2. Define the Cli parameters to be parsed using
   `ConfigStateMap`:

```c++
map->add("--help", "option.help", Bool, true, true);
map->add("-h", "option.help", Bool, true, true);
map->add("--option1", "option.a", String, true, false);
map->add("--option2", "option.b", String, true, false);
map->add("--option3", "option.c", Int, true, false);
```

3. Then define your `Configuration` object:

  ```c++
  Configuration cfg;
  ```

4. Create the `CommandLineParser` object instance and pass in
   the inputs:

```c++
CommandLineParser parser(&cfg, map, argc, argv);
```

where `argc` and `argv` are the commandline argument
parameters passed by the operating system into `main()`.

### Expected Behavior

### General

1. The `CommandLineParser` will iterate over
   the `ConfigStateMap`
2. It then loads any mapped command-line parameters
   from `argv` into
   the `Configuration` object, casting the `argv` string
   to the appropriate datatype.

### Boolean are special

For a boolean variable identified in `ConfigStateMap` the
value is `true` if the flag is defined or `false` by default.