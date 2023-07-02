Disable a Build Project
=======================

You can disable builds for a project using the following:

```bash 
make disable/build/program/${YOUR_PROJECT_NAME}
```

* This will create `cmd/<project>/build.disabled` in your local repo.
* You **can** check in a `build.disabled` file and GitHub actions will respect it.
* The `build.disabled` contains the user, host and timestamp of the operation.
  For example:
  ```text
    samcaldwell@mbp.local : Fri Jun 16 21:56:57 CDT 2023
  ```
