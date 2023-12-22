Workflow Actions
================

## Objectives
* A `workflowAction` may be `null` if it does nothing.
* But a non-null `workflowAction` describes a topic and message template which Tracker will publish to the message queues used to execute integration logic.  For example, this could cause a message to be sent (emails, chat messages) or other actions (CI/CD logic).
