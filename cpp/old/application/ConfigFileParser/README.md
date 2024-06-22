ConfigFileParser
================

This project extends Configuration to parse a supported
configuration file (e.g. YAML) into the shared configuration
tree.  The `parse()` function will leverage `projects/Parser`
project to parse a given `YAML` configuration file into the
`Configuration` tree using `ConfigStateMap` object.  In the 
end, the `ConfigFileParser` will expose only the configuration
tree.
