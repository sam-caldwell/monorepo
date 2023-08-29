Command: create-project
=======================

## Description

This program creates a new project within the monorepo as a directory
under `projects/<project_name>` and adds the minimum skeleton
(`MANIFEST.yaml` and `README.md`).

## Usage

```text
create-project [arguments]

    arguments:
        -author <string>
        -project <string>
        -language <string>
        -manifestOnly
        
    languages:
        For -language, only supported languages are allowed
        These are identified in `monorepo.SupportedLanguages`
```