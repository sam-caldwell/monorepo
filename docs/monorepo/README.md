Monorepo Automation
===================

Monorepo Command
-------------
The `monorepo` command is the unified tool for managing and interacting with this monorepo.
This manages the configuration of the monorepo as well as allows a user to perform common
actions abstracted from the underlying host, such as:

* Building projects
* Linting source code
* Performing security scans
* Running tests
* Distributing builds/tasks across one or more machines.

Command Usage
-------------

## General Commands

| Command                    | Description                        | 
|----------------------------|------------------------------------|
| `monorepo -h \| --help`    | Show the command usage             |
| `monorepo -v \| --version` | Show the command's current version |

## General Options

All commands support the following general options:

| Option    | Description                         |
|-----------|-------------------------------------|
| `--color` | display all output using ANSI color |
| `--debug` | print verbose logging messages      |

## Monorepo Management Commands

### Project Management

The following commands allow the user to configure/manage the monorepo itself. (The changes made by this command
can be reviewed and committed to git separately.)

| Command                                            | Description                                                 | 
|----------------------------------------------------|-------------------------------------------------------------|
| `monorepo project list [--language <lang>]`        | Print a list of all projects in the monorepo.               |
|                                                    | The --language flag can be used to filter the list to       |
|                                                    | a specific language.                                        |
| `monorepo project show <name>`                     | Show the RESOLVED configuration for a given project, where  |
|                                                    | the resolved configuration is the final configuration state |
|                                                    | after starting with <root>/monorepo.yaml, then applying the |
|                                                    | language-specific monorepo.yaml and finally after applying  |
|                                                    | the project-specific configuration.                         |
| `monorepo project delete <name>`                   | Delete a given project from the monorepo                    |
|                                                    |                                                             |
| `monorepo project create <name> [project options]` | Create a new project in the monorepo using the given        |
|                                                    | name and options.  Any options not specified will result    |
|                                                    | in the defaults being used.                                 |
| `monorepo project enable <feature> <name>`         | Enable the given project and feature                        |
| `monorepo project disable <featur> <name>`         | Disable the given project and feature                       |

#### Project Options

| Option                              | Description                                        |
|-------------------------------------|----------------------------------------------------|
| `--description <string>`            | Project description                                |
| `--author <string>`                 | Project author (used in copyright strings)         |
| `--output <executable   \| source>` | Project output (executable or reusable source code | 
| `--enable <feature>`                | Enable a feature flag (linter,build,tests)         |
| `--addPlatform <opsys string>`      | Add a supported platform                           |

### Build Host Configuration

The following commands allow the user to manage the build hosts on which monorepo tasks will execute.
There must be at least one host configured for tasks to run, including localhost.

| Command                                      | Description                                                     | 
|----------------------------------------------|-----------------------------------------------------------------|
| `monorepo host list`                         | List the currently configured build hosts.                      |
| `monorepo host show <name>`                  | Show the RESOLVED configuration for a given build host, where   |
|                                              | the resolved configuration is the configuration state after all |
|                                              | configuration is loaded, validated and resolved in order        |
|                                              | of precedence.                                                  |
| `monorepo host delete <name> `               | Delete the named host configuration.                            |
| `monorepo host create <name> [host options]` | Create a new host configuration for the monorepo using the      | 
|                                              | given name and options.                                         |
| `monorepo host enable <name>`                | Enable a given build host.                                      |
| `monorepo host disable <name>`               | Disable a given build host.                                     |
| `monorepo host check <name>`                 | Execute a test command on a given build host to verify          |
|                                              | connectivity.                                                   |

#### Host Options

| Option              | Description                                                                    | 
|---------------------|--------------------------------------------------------------------------------|
| `--host <string>`   | The IP address or FQDN and port for the target host.                           |
|                     | If localhost is used, no port is required since local execution will be used.  |
| `--enabled`         | Sets the host to enabled (default is disabled).                                |
| `--username`        | Declare the username used for connectivity.  If not specified the current user |
|                     | is used.  Assumes all authentication is key-based.                             |
| `--platform <name>` | Specify the target host's platform.                                            |
| `--hypervisor`      | Define what hypervisor is expecte on remote host.                              |
| `--cri`             | Define what container runtime. is expecte on remote host.                      |

### Container Runtime Configuration

The following commands allow the user to manage the container runtimes:

| Command                                    | Description                                                      | 
|--------------------------------------------|------------------------------------------------------------------|
| `monorepo cri list`                        | List the supported container runtimes (e.g. docker, kubernetes). |
| `monorepo cri create <name> [CRI Options]` | Create a new container runtime interface (CRI) configuration.    |
| `monorepo cri delete <name>`               | Delete a container runtime interface (CRI) configuration.        |
| `monorepo cri show <name>`                 | Show the current container runtime interface configuration.      |

#### CRI Options

| Option              | Description                                        | 
|---------------------|----------------------------------------------------|
| `--enabled`         | Enable the resulting config. (default is disabled) |
| `--platform <name>` | Specify the target host's platform.                |

### Hypervisor Configuration

The following commands allow the user to manage the supported hypervisors:

| Command                                                  | Description                                | 
|----------------------------------------------------------|--------------------------------------------|
| `monorepo hypervisor list`                               | List the supported hypervisors             |
| `monorepo hypervisor create <name> [Hypervisor Options]` | Create a new hypervisor configuration.     |
| `monorepo hypervisor delete <name>`                      | Delete a hypervisor configuration.         |
| `monorepo hypervisor show <name>`                        | Show the current hypervisor configuration. |
| `monorepo hypervisor enable <name>`                      | enable the given hypervisor                | 
| `monorepo hypervisor disable <name>`                     | disable the given hypervisor               | 

#### Hypervisor Options

| Option              | Description                                        | 
|---------------------|----------------------------------------------------|
| `--enabled`         | Enable the resulting config. (default is disabled) |
| `--platform <name>` | Specify the target host's platform.                |

#### Language Configuration

The following commands allow the user to manage the supported languages:

| Command                                              | Description                                             | 
|------------------------------------------------------|---------------------------------------------------------|
| `monorepo language list`                             | List the currently supported languages                  |
| `monorepo language show <name>`                      | Show the language configuration for the named language  |
| `monorepo language create <name> [Language Options]` | Create a language directory structure and configuration |
|                                                      | file.                                                   |
| `monorepo language delete <name>`                    | Delete a given language configuration file (source code |
|                                                      | will remain).                                           |
| `monorepo language enable <name>`                    | enable the language                                     |
| `monorepo language disable <name>`                   | disable the language                                    |

#### Language Options

| Option                        | Description                                                         | 
|-------------------------------|---------------------------------------------------------------------|
| `--directory <relative path>` | The child directory/path where the language source code will exist. |
| `--enabled`                   | Enable the language config (disabled by default)                    |
| `--platform <name>`           | Specify the platform supported by the given language.               |

### Platform Configuration

The following commands allow the user to manage the supported platforms (opsys, arch):

| Command                                              | Description                                         | 
|------------------------------------------------------|-----------------------------------------------------|
| `monorepo platform list`                             | List the currently supported platforms              |
| `monorepo platform create <name> [Platform Options]` | Create a new supported platform configuration file. |
| `monorepo plaform show <name>`                       | Show the named platform configuration               |
| `monorepo platform delete <name>`                    | Delete the named platform                           |
| `monorepo platform enable <name>`                    | Enable the given platform                           |
| `monorepo platform disable <name>`                   | Disable the given platform                          |

#### Platform Options

| Option                    | Description                                              | 
|---------------------------|----------------------------------------------------------|
| `--description <string>`  | A simple description of the platform.                    |
| `--connect local\|ssh`    | Connection method (local execution or key-based SSH)     |
| `--enabled`               | Enable the configuration                                 |
| `--opsys_family <regex>`  | A regex for a single supported operating system families |
| `--opsys_version <regex>` | A regex for a single supported operating system version. |
| `--cpu_arch <regex>`      | A regex for a single supported CPU Architecture          |

## Monorepo Build Automation Tooling

### Builders

The following commands will facilitate building projects for the monorepo:

| Command                         | Description                                                 | 
|---------------------------------|-------------------------------------------------------------|
| `monorepo build all`            | Build all projects for all languages                        |
| `monorepo build project <name>` | Build the named project                                     |
| `monorepo build order`          | Print the list of projects in the order they would be built |

### Linters

The following commands will facilitate linting projects for the monorepo:

| Command                        | Description                         | 
|--------------------------------|-------------------------------------|
| `monorepo lint all`            | Lint all projects for all languages |
| `monorepo lint project <name>` | Lint the named project              |

### Test Runner

The following commands will facilitate linting projects for the monorepo:

| Command                        | Description                         | 
|--------------------------------|-------------------------------------|
| `monorepo test all`            | Test all projects for all languages |
| `monorepo test project <name>` | Test the named project              |




    

