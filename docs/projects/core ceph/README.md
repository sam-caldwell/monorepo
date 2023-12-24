## Objectives
1. Create a containerized Ceph deployment for local development which makes local deployment possible with less effort.
2. Create a containerized Ceph deployment for the on-premises environments used by `Singularity` and other projects.  The production deployment should be functionally consistent with the local development environment.
3. Ensure that the solutions are managed through the `monorepo` command for automated build, test, install and deploy.
------
## What is Ceph?
Ceph is a highly scalable storage solution for providing s3-compatible object storage, block storage or file storage infrastructure.

---
## References
[[Ceph Research Notes]]
* [Ceph Docs](https://docs.ceph.com/en/latest/start/intro/)

---
## Definition of Done / Requirements
1. ***Developer Requirements***
	1. As a developer, I should be able to execute `monorepo clean` to terminate all instances of my development environment containers, including the local `ceph` cluster.
	2. As a developer, I should be able to execute `monorepo build` to build the container images used by `ceph` 
	3. As a developer, I should be able to execute `monorepo test` to run unit tests against the container images used by `ceph` to ensure they are working as expected before I use them in my project.
2. ***User Requirements (building/operating derivative projects)***
	1. As a user (building derivative projects), I should be able to execute `monorepo build` to build local copies of my `ceph` container images.
	2. As a user (building derivative projects), I should be able to execute `monorepo install` to spawn a local `ceph cluster` and install any associated tooling (programs).
	3. As a user (building derivative projects), I should be able to execute `monorepo deploy` to push the `ceph` container images to the internal container hub for deployment to the on-premises runtime environment (servers).
---

