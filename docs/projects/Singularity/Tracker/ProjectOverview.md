| Subsystem | Description |
	| --- | --- |
	| [TrackerCli](./trackerCli/README.md) | Command-line interface for Tracker System |
	| [TrackerUI](./trackerUi/README.md) | A web interface to the Tracker System |
	| [TrackerAPI](./trackerApi/README.md) | A Golang API server providing core functionality to the Tracker System. |
	| [TrackerDb](./trackerDb/README.md) | The PostgreSQL database core. |

## [TrackerCli](./trackerCli/README.md)
1. Provides command-line interface to the tracker system.
2. Provides a means of rapid feature development for the client-server and database.
3. Facilitates scripted solutions to work with Tracker.
## [TrackerUI](./trackerUi/README.md)
1. Provides a web UI for the Tracker.
## [TrackerAPI](./trackerApi/README.md)
1. Provides an API server for the Tracker to interact with the database and other integrated systems.
2. Tracker API supports integration with--
	1. PostgreSQL database (primary data store).
	2. Kafka (communications stream to integrated systems).
## [TrackerDb](./trackerDb/README.md)
1. Provides a primary data store for the workflow data.
