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
 - cmd/sqldbtest/badgeMaker
 - cmd/sqldbtest/bump-version
 - cmd/sqldbtest/findBuildProjects
 - cmd/sqldbtest/has-executable
 - cmd/sqldbtest/locks
 - cmd/sqldbtest/set-version
 - cmd/sqldbtest/what
 - cmd/validators/yamlValidator

```
