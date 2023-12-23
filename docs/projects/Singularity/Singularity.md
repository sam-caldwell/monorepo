## Objective
* Create a scalable platform for building heterogenous services and projects which will reach a minimum economy of scale for the home automation and other work.
* Singularity is the overall BHAG project to unify several smaller projects. 

The following diagram illustrates the three projects making up `singularity`:
1. [***Tracker Project***](./docs/projects/Singularity/Tracker/ProjectOverview.md) - a workflow engine for integrating ticket tracking and task automation into a single system 

	| Subsystem | Description |
	| --- | --- |
	| TrackerCli | A command-line interface to the Tracker System |
	| TrackerUI | A web interface to the Tracker System |
	| TrackerAPI | A Golang API server providing core functionality to the Tracker System. |
	| TrackerDb | The PostgreSQL database core. |

3. [***Data Collectors***](./docs/projects/Singularity/DataCollectors/ProjectOverview.md) - a single microservices platform for collecting information from arbitrary online and on-premises services and for storing that data to a long-term `CEPH` archival system.

4. [***Home Automation***](./docs/projects/Singularity/HomeAutomation/ProjectOverview.md) - Various hardware projects (such as those using ESP32 and Rock64) to collect wireless signals, environmental signals, camera feeds and other data into a structured form, stored on the Data Collectors CEPH cluster.

![[Pasted image 20231223165323.png]]