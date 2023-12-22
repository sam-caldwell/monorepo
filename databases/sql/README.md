SQL (Databases)
===============

## Objectives
* Define a means of managing all database schema for all projects.
* Ensure the database schema are not tied to any specific framework/language to allow broad use.

## Database Migrations and tests.
1. The `monorepo build` command will discover the database and apply its migrations.
2. The `monorepo test` command will discover all databases, apply the tests and run golang tests defined alongside
   the SQL code.

## Database Documentation
* [ER Diagram](../../../docs/projects/Tracker/trackerDB/trackerDb.er.diagram)