Command: list-projects
======================

## Description

This command with list the projects in the monorepo using optional filters
to display only the matching projects. This project relies heavily on the
`repotool/filters` and `repotool/lister` projects.

<div>
<h1 style="color:red">WARNING:</h1>
<p style="color:red">
    A MANIFEST.txt file must exist in the project for this tool
    to see and process it.  If no MANIFEST.txt file exists, the
    tool will assume it is simply a directory.
</p>
</div>

## Usage

```text
list-projects [display flags] [filter flags]

    display flags:
        -color   // Show the output in color
        -banner  // Pretty print the output with a banner. 

    filter flags:
        -commands      // Show only command projects
        -buildEnabled  // Show only build-enabled projects
        -lintEnabled  // Show only lint-enabled projects
        -scanEnabled  // Show only scan-enabled projects (e.g. security scanners)
        -signEnabled  // Show only sign-enabled projects (e.g. projects ready to code-sign)
        -packEnabled  // Show only packaging-enabled projects (e.g. projects ready for packaging)
        -os <opsys>  // Show only projects which support the given operating system
        -arch <arch>  // Show only projects which support the given CPU Architecture 
```

| filter flag... | ...if present                                    | ...if not present                       |
|----------------|--------------------------------------------------|-----------------------------------------|
| -commands      | list only `cmd/<project>/<program>` projects     | List only `projects/<project>` projects | 
| -buildEnabled  | list projects with buildEnabled in manifest      | no action                               | 
| -lintEnabled   | list projects with lintEnabled in manifest       | no action                               | 
| -scanEnabled   | list projects with scanEnabled in manifest       | no action                               | 
| -signEnabled   | list projects with signEnabled in manifest       | no action                               | 
| -packEnabled   | llist projects with packEnabled in manifest      | no action                               |
|                |                                                  |                                         |
| -buildDisabled | list projects with buildDisabled in manifest     | no action                               |
| -lintDisabled  | list projects with lintDisabled in manifest      | no action                               |
| -scanDisabled  | list projects with scanDisabled in manifest      | no action                               |
| -signDisabled  | list projects with signDisabled in manifest      | no action                               |
| -packDisabled  | list projects with packDisabled in manifest      | no action                               |
|                |                                                  |                                         |
| -os <opsys>    | list projects with the matching operating system | no action                               |
| -arch <arch>   | list projects with the matching CPU Architecture | no action                               |

