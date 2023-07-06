Command: monorepo
=================

## Description
This project provides the `monorepo` command for maintaining and interacting
with the monorepo itself. 

---
## Usage:

---
#### Get Help
Display the help / usage message.
> ```text
>   monorepo --help
> ```

---
#### Show Version
Display the program version.
> ```text
>   monorepo --version
> ```

---
#### Build Projects
This command will build either every enabled project or a specific enabled project.
> ```text
>    monorepo build [options]
> ```

> * ***Options***
>   * Build a specific project `--project <name>`
>     * Will only build if the project is enabled
>   * Will only build projects in the `cmd` directory as build implies producing an executable or library.
>   * Noop `--noop` will simply show what project(s) would have been built.
>   * Verbose Logging `--verbose` will enable ANSI colored logging on stdout.

---
#### Initialize Projects
This command will install the monorepo dependencies.
> ```text
>    monorepo init [options]
> ```

> * ***Options***
>   * Noop `--noop` will simply show what project(s) would have been built.
>   * Verbose Logging `--verbose` will enable ANSI colored logging on stdout.
---
#### Lint The Repo
This command will run every linter against the entire repository.
> ```text
>    monorepo lint
> ```

> * ***Options***
>   * Noop `--noop` will simply show what project(s) would have been built.
>   * Verbose Logging `--verbose` will enable ANSI colored logging on stdout.

---
#### List Projects (enabled or disabled)
This command will list all the enabled or disabled projects in the repo.
> ```text
>   monorepo list [options]
> ```

> * ***Options***
>   * Filter options (`--enabled`| `--disabled` | `--all`) determined which projects are listed
>   * Noop `--noop` will simply show what project(s) would have been built.
>   * Verbose Logging `--verbose` will enable ANSI colored logging on stdout.
 
---
#### Scan The Repo
This command will run every scanner against the entire repository.
> ```text
>   monorepo scan [options]
> ```

> * ***Options***
>  * Noop `--noop` will simply show what project(s) would have been built.
>  * Verbose Logging `--verbose` will enable ANSI colored logging on stdout.
