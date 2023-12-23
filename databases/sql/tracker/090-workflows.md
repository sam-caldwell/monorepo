### Table: Workflows
* A `workflow` is a collection of steps representing the state of the ticket in a longer progression.
* The `workflow` steps represents a step in a progression, and it may also map to an `action`.

### Table: WorkflowSteps
* The steps of a given workflow are represented by a set of records in the `WorkflowSteps` table.
* These steps represent the step as a position in the workflow.
* The `workflowSteps` records may have a `null` or UUID action value.
		* Where `action` is `null`, no action will be taken and the step will only exist as a marking of state within the workflow.
		* Where `action` is not `null`, the given UUID will identify an action defined in the [`workflowactions`](140-workflowActions.md) table (below).

### Workflows Functions:
##### [`createWorkflow()`](./096-createWorkflow.sql)
* The `createWorkflow()` function will create the `workflow` object/record and establish its permissions, ownership and identity.
* Once the record is inserted into the `workflows` table, the function calls `initializeWorkflowSteps()` to create the `start` and `terminate` nodes.

##### [`initializeWorkflowSteps()`](./094-initializeWorkflowSteps.sql)
* When a `workflow` object/record is created, a corresponding `start` and `terminate` record must be created in `workflowSteps`.
* The `initializeWorkflowSteps()` creates these two `workflowSteps` records as the initial linked list state in the database, as illustrated, below: <br/>![[workflowSteps-initial-state.png]]

##### [`deleteWorkflowById()`](./096-deleteWorkflowById.sql)
* Given a `workflowId` (UUID) delete the workflow record.

##### [`getWorkflowById()`](./096-getWorkflowById.sql)
* Given the `workflowId` (UUID) return a JSON object corresponding to the workflow record.

##### [`getWorkflowsByOwnerId()`](./096-getWorkflowsByOwnerId.sql)
* Given a user UUID (`ownerId`), return a list of JSON workflow objects where each resulting object is owned by the given owner.

##### [`getWorkflowsByTeamId()`](./096-getWorkflowsByTeamId.sql)
* Given a `teamId` (UUID), return a list of JSON workflow objects where each resulting object is owned by the given team.

### WorkflowSteps Functions:
##### [`getStartNode()`](./093-getStartNode.sql)
* Given a `workflowId`, the `getStartNode()` function will return the `workflowStepId` of the `'start'` node of the identified workflow.

##### [`getTerminalNode()`](./093-getTerminalNode.sql)
* Given a `workflowId`, the `getTerminalNode()` function will return the `workflowStepId` of the `'terminate'` node of the identified workflow.

##### [`workflowNodeValidation()`](./093-workflowNodeValidation.sql)
* Given an insert/update on the `workflowSteps` table, the `workflowNodeValidation()` function will be invoked by trigger.
* The purpose of this function is to ensure the linked list of `workflowSteps` maintains a proper structure.
##### [`createWorkflowStep()`](./096-createWorkflowStep.sql)
* This function creates a properly linked `workflowsteps` node for a given workflow.
##### [`deleteWorkflowStep()`](./096-deleteWorkflowStep.sql)
* This function deletes a given workflow step, relinking the remaining steps around the deleted node.
##### [`linkNodes()`](./094-linkNodes.sql)
* When a workflowStep is deleted, this function relinks the surrounding nodes.
