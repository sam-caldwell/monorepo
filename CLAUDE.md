# CLAUDE.md

## Overview

Personal polyglot monorepo owned by Sam Caldwell. Contains shared libraries, CLI tools, infrastructure configs, and experiments across Go, Python, and shell. The monorepo exists to avoid per-project setup overhead and share common code/tooling. Licensed MIT (with per-project overrides — check subdirectory LICENSE files).

## Repository Layout

```
projects/     - ~45 packages organized by project (Go): utilities, crypto, networking, CLI tools, testing, AWS/Atlassian SDKs
containers/   - Docker images built from local base (no third-party pulls)
  opsys/      - Base OS images: ubuntu/24.04 (CIS-hardened), distroless (CIS-hardened scratch-based)
  language/   - Multi-stage builds: java/21, node (20/21/latest), python (3.14/latest)
  services/   - PostgreSQL (16/17/18 with ubuntu+distroless targets), Ceph, DNS
CMake/        - Reusable CMake modules (logging, colors, project discovery)
```

## Build & Test

### Go
```sh
go test ./projects/...
```
- Module: `github.com/sam-caldwell/monorepo`, Go 1.26+ (toolchain go1.26.0)
- Uses a `replace` directive pointing to `./` for local development
- Run `staticcheck` for linting (config: `staticcheck.conf`)

### Containers
```sh
make build/docker
```
- Build order: opsys -> language -> services (see `containers/README.md`)
- All images built from local `opsys/ubuntu:24.04` base — no third-party images
- Language images provide `builder` and `runtime` multi-stage targets
- PostgreSQL images provide `ubuntu` and `runtime` (distroless) targets

### Python
```sh
pip install -r requirements.txt
pytest
flake8
```

## Code Style & Conventions

- **EditorConfig is authoritative** — see `.editorconfig`:
  - Go: tabs, 4-width
  - Python: spaces, 4-width
  - JS/TS: spaces, 2-width
  - LF line endings everywhere
- **Go**: standard `gofmt` style; packages are flat under `projects/<name>/`
- **Commit messages**: lowercase, descriptive, no conventional-commits prefix (e.g., "Adding a Print() method to Matrix")

## Key Dependencies (Go)

- AWS SDK v2 (S3, EC2)
- go-git v5
- 1Password connect-sdk-go
- ProtonMail gopenpgp
- cobra (CLI framework)
- lib/pq (PostgreSQL)
- gocv (OpenCV bindings)
- libvirt Go bindings

## Security

- Snyk is configured (`.snyk`) for vulnerability scanning
- Dependabot watches Go module dependencies (weekly)
- Container base images are CIS-hardened (sysctl, SSH, SUID removal, login policies)
- Report vulnerabilities to mail@samcaldwell.net or via PR
- Some projects are filtered from GitHub publication (exploit PoCs, etc. stay on local git server)

## Things to Know

- Not all code is published here — a local git server filters what goes to GitHub
- The `jokes/` directory is exactly what it sounds like
- `cmake-build-debug/` is a local build artifact directory (gitignored)
- Git LFS is configured for archives, ISOs, and certificates (see `.gitattributes`)
