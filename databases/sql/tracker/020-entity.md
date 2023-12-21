entity table
============

## Objectives
* An entity is a universal object identifier (uuid) for the entire tracker database. This is intended as the basis 
  of our accountability system.  The entity system provides a write-only object registry with timestamps and context.

## Use case
* We never delete an entity Id.  So even if we delete the object it was created for, there remains an artifact which
  can be referenced by logs and given context by this table to allow an audit trail.
