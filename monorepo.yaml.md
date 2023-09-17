Monorepo Configuration file
============================

-----

# Description

This file documents how `monorepo.yaml` should work. The `monorepo.yaml`
file is the default configuration file for the monorepo command to interact
with the repository.

The file can be edited manually or using `monorepo config`.

# Configuration File

---

## `artifact_directory`

The local directory where build artifacts will be stored.
if build operation is remote, the artifact will be copied
to this local directory once completed.

The 'artifact_directory' is relative to the repository
root directory.

### Example

```
artifact_directory: build
```

---

## `build_hosts`

Defines the hosts on which distributed builds can be run using the `monorepo exec` command. A "build host" is a
computer on which monorepo operations such as builds will be run.

### Fields:

| Field      | Definition                                               |
|------------|----------------------------------------------------------|
| name       | Must be unique                                           |                                  
| enabled    | Enables/disables the build host.                         |                
| host       | Should be a valid hostname, FQDN or IP address.          | 
| connect    | Values (local \| ssh) indicate connectivity to the host. | 
| username   | If null, we use the current user when tools run.         |
| opsys      | Values must be in the platforms list                     |
| hypervisor | Values must be in the hypervisors list                   |
| cri        | Values must be in the cri list                           |

### Example

```yaml
build_hosts:
  - name: local mac laptop
    description: This is the local mac dev machine.
    host: mac-builder.local
    connect: local
    enabled: true
    username: null
    opsys: macos
    hypervisor: parallels
    cri: docker
  - name: local linux laptop
    description: This is a local linux dev machine.
    host: linux-builder.local
    connect: ssh
    enabled: false
    username: null
    cri: docker
    opsys: linux
    hypervisor: kvm
  - name: local windows laptop
    host: windows-builder.local
    connect: ssh
    enabled: false
    username: null
    cri: null
    opsys: windows
    hypervisor: vmware
```

---

## `container_runtimes`

Defines the supported container runtimes which can be configured for a `build_host`.

- This will constrain the runtime to specific platforms (opsys/cpu)
- Defining a supported CRI simply makes the CRI available for use in defining a build host (pairing platforms
  with a CRI, if appropriate).
- Examples: Docker, Kubernetes

### Fields

| Field       | Definition                                   |
|-------------|----------------------------------------------|
| name        | The name of the container runtime definition |
| description | A descriptive string                         |
| platforms   | A list of operating system strings           |

### Example

```yaml
container_runtimes:

  - name: docker
    enabled: true
    platforms:
      - darwin
      - macos
      - linux
      - windows
  - name: minikube
    enabled: true
    platforms:
      - macos
      - linux
```

---

## `hypervisors`

Define the hypervisors supported by the monorepo

- This will constrain the hypervisor to specific platforms (opsys/cpu)
- Defining a supported hypervisor makes the hypervisor available for use in defining a build host
  (pairing platforms with a hypervisor, if appropriate).
- Examples: parallels, vmware, kvm, etc.

### Fields

| Field       | Definition                                   |
|-------------|----------------------------------------------|
| name        | The name of the container runtime definition |
| description | A descriptive string                         |
| platforms   | A list of operating system strings           |

### Example

```yaml
hypervisors:

  - name: parallels
    enabled: true
    platforms:
      - macos
  - name: vmware
    enabled: false
    platforms:
      - macos
      - linux
      - windows
  - name: hyper-v
    enabled: false
    platforms:
      - windows
  - name: virtualbox
    enabled: false
    platforms:
      - linux
      - windows
  - name: kvm
    enabled: false
    platforms:
      - linux
```

---

## `language`

Define the supported programming languages for the monorepo and how the language will be handled (scripts).

### Fields

| Field       | Definition                                                                                          |
|-------------|-----------------------------------------------------------------------------------------------------|
| name        | The name of the container runtime definition                                                        |
| description | A descriptive string                                                                                |
| enabled     | A boolean indicating that the language is enabled for use.                                          |
| directory   | A directory path relative to the repository root where the language sources will be found.          |
| tasks       | A map of tasks describing executable steps `monorepo exec` will use to perform specific operations. |
| tasks.build | The build task.                                                                                     |
| tasks.lint  | The linter task.                                                                                    |
| tasks.scan  | The security scanner task.                                                                          |
| tasks.sign  | The cryptographic signing task.                                                                     |
| tasks.test  | The task for runningn language-specific tests.                                                      |

### Syntax

```yaml
languages:
  name: Programming language name
  enabled: Enables/disables hypervisor
  directory: Relative path to source files (REPO_ROOT/go/<directory>)
  tasks:
    build: <operation descriptor>
    lint: <operation descriptor>
    scan: <operation descriptor>
    sign: <operation descriptor>
    test: <operation descriptor>
```

| Task Value               | Description                                              |
|--------------------------|----------------------------------------------------------|
| `<operation descriptor>` | An executable string                                     |
| `null`                   | When null, a disabled message will be written to stdout. |

### Example

```yaml
languages:
  - name: Go
    description: Golang Programming language
    directory: go
    enabled: false
    tasks:
      build: |
        GOOS=${OPSYS} \
        GOARCH=${CPU_ARCH} \
        go build -o ${OPSYS}/${CPU_ARCH}/${PROJECT_NAME}${EXTENSION} go/${PROJECT_NAME}/main.go
      lint: |
        GOOS=${OPSYS} \
        GOARCH=${CPU_ARCH} \
        go vet -v -asmdecl -atomic -bool -buildtags -assign -cgocall -composites \
             -copylocks -httpresponse -lostcancel -methods -nilfunc -printf \
             -rangeloops -shift -tests -unreachable -unsafeptr -unusedresult \
             -unusedstringmethods -unusedfuncs ./go/...
      scan: null
      sign: null
      test: |
        GOOS=${OPSYS} \
        GOARCH=${CPU_ARCH} \
        go test -v  ./${PROJECT_NAME}
```

---

## `license`: Repo-level LICENSE file

Define the license file for the repository, where the file is relative to the repository.

### Example

```yaml
license: LICENSE.txt
```

---

## `platforms`

Define the supported platforms. A platform is a combination of supported operating system and CPU combination.

### Fields

| Field                           | Definition                                                   |
|---------------------------------|--------------------------------------------------------------|
| `name`                          | The name of the container runtime definition                 |
| `description`                   | A descriptive string                                         |
| `enabled`                       | A boolean indicating that the platform is available for use. |
| `supported_versions`            | A map of opsys `family`, `version` or `cpu_arch` patterns    |
| `supported_versions`.`family`   | A regular expression used to match opsys `family`            |
| `supported_versions`.`version`  | A regular expression used to match opsys `version`           |
| `supported_versions`.`cpu_arch` | A regular expression used to match the `cpu_arch`            |

### Example

```yaml
platforms:

  - name: macos
    description: MacOS Ventura or later
    enabled: true
    supported_versions:
      - family: macOS Ventura
        version: 13.[0-9]+.[0-9]+ # accept any 13.5.xx version
        cpu_arch: (x86_64|arm64)
  - name: linux
    description: Ubuntu or Debian Linux
    enabled: true
    supported_version:
      - family: Ubuntu
        version: 2[23456789].[0-9]+
        cpu_arch: (x86_64|arm64)
```