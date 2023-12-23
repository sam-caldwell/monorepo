Tickets
=======

## Objectives
* A `Ticket` is a single unit of work.  The ticket represents the work to be performed as it travels through a `workflows`.

## Data Constraints
* `id` (ticket identifier) must be in [`entity`](./020-entity.md).
* `projectId` must be in [`projects`.`id`](./110-projects.md)
* `authorId` must be in `users`.`id`
* `assigneeId` must be in `users`.`id`
* `ticketType` must be a supported type for the given `projectId`
* `workflowStepId` must be a valid step in `workflowSteps`
* `owner`, `team`, `everyone` permissions must be conflict free.  See `verifyPermissions` function.
## Functions

| Function | Description |
| --- | --- |
| `createTicket()` | This function creates a ticket record along, initializing the current `workflowStepId` to the starting step of the associated `ticketType`.
| `deleteTicket()` | This function deletes a given ticket record. |
| `getTicketsByAssignee()` | Given a `userId` representing an assignee, return all tickets to which the user is assigned.|
| `getTicketsByAuthor()`  | Given a `userId` representing an assignee, return all tickets for which the user is an author. |
| `getTicketsByProject()` | Given a `projectId`, return a list of tickets for the given project.| 
| `getTicketsByType()` | Given a ticket type, return a list of associated tickets.|


### `createTicket()`
* Create a new `ticket` record with its--
	* `projectId` (`projects`.`id`)
	* `authorId` (`users`.`id` entity)
	* `assignee` (`users`.`id` entity)
	* `type` (`ticketTypes`.`id` entity)
	* permissions (`owner`,`team`,`everyone`)
	* summary (`varchar`) and description (`text`)
* Initialize the `ticket` initial `workflowSteps`.`id` value:
	* Given the `type` (`ticketTypes`.`id`), we can obtain the `workflowId` from `ticketTypes`.
	* This can then be used to query the `workflowSteps` table for the step `id` and obtain the `start` step.  This is the first step (where `name` is "start" and `prevStepId` is `null`) in the `workflow`.
