List the Enabled Projects
=========================

It's easy to find out what projects are in the `cmd` directory that are **ENABLED**:

```bash
make list/build/projects
```

This will generate output like this...
```bash
samcaldwell@mbp go % make list/build/projects
start: list/build/projects [OPSYS: darwin]

current binary projects (enabled):
 - output:
 - cmd/ansi/color
 - cmd/calculateSubnets/calculateSubnets
 - cmd/tools/badgeMaker
 - cmd/tools/bump-version
 - cmd/tools/findBuildProjects
 - cmd/tools/has-executable
 - cmd/tools/locks
 - cmd/tools/set-version
 - cmd/tools/what
 - cmd/validators/yamlValidator

```
