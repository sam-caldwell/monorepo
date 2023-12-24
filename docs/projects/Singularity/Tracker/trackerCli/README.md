## Objectives:
1. This tool is the first point for new feature development.
	1. Command-line provides easier implementation
	2. Command-line can facilitate feature integration testing with less effort and more consistent results.
2. This tool should allow users to interact via command line (terminal) for everything they could otherwise do in the trackerUI (web application).
3. Command syntax should be intuitive and consistent.

## Usage

The basic command syntax is as follows:
```bash
tracker <subcommand> <object> <options>
```

#### Subcommands
* `create`
* `join`
* `get`
* `update`
* `delete`
* `help`

#### Objects
| Object | Description |
| --- | --- |
| `avatar` |  An avatar is an image file which identifies a user. | 
| `icon` |  An icon is an image file which can be used to identify a system object. | 
| `user` |  A user is a human or system which interacts with the Tracker. | 
| `team` |  A team is a collection of users. | 
| `project` |  A project is a collection of related tickets representing a collective effort. | 
| `ticketType` |  A ticket type is a classification of tickets with a common workflow. | 
| `workflow` |  A workflow is a collection of steps needed to complete some task. | 
| `workflowStep` |  A workflow step is the state of a ticket within the progression of a workflow. | 
| `workflowAction`|  A workflow action is some message the Tracker will emit to the Kafka streams for processing by integrated systems, which may, for example indicate a message to notify a user of some state or to cause another systems to take some action in response to a ticket reaching some state in the workflow. | 
| `ticket` | A ticket is the atomic unit of work representing a single task within the Tracker system. | 
| `comment` | A comment is a block of text giving clarity to a ticket. | 
| `attachment` | An attachment is a file or other data object attached to a ticket. | 

### `tracker create avatar`
Given an avatar image file, this command will resize and upload the file into the Tracker system and return its `avatarId` (UUID).
```bash
tracker create avatar [source file]
```

#### Theory of Operation:
1. Determine the file's hash (SHA-256).
2. Determine the file's mime-type.
3. Resize the Avatar to 300x300 pixels.
4. Create an `avatarId` record in the `avatars` database table based on the file's hash and mime type.
5. Upload the file to the Tracker file store (Ceph).
6. Print the `avatarId` to `stdout`.
7. Terminate.

### `tracker create icon`
Given an icon image file, this command will resize the image to a standard icon size (50x50 pixels), upload the resized artifact to the Tracker system and return its `iconId` (UUID).
```bash
tracker create icon [source file]
```
#### Theory of Operation:
1. Determine the file hash (SHA-256).
2. Determine the file mime type
3. Resize the icon file to 50x50 pixels.
4. Create an `iconId` record in the `icons` database table based on the file's hash and mime type.
5. Upload the file to the Tracker file store (Ceph).
6. Print the `iconId` to `stdout`.
7. Terminate.

### `tracker create user`
Given the following input parameters, create a user object in the `users` database and return the new user's `userId` (UUID):
```bash
tracker create user --email <emailAddress> \
					--firstName"<first name>" \
					--lastName "<last name>" \
					--avatar <avatarId> \
					[-phone <phoneNumber>] \
					[--description <string>]
```

### `tracker create team`
Given the following input parameters, create a user object in the `teams` database table and return the new `teamId` (UUID):
```bash
tracker create team --name <team name> \
					--iconId <iconId> \
					--ownerId <userId> \
					--permissions {owner:<o>,team:<t>,everyone:<e>} \
					[--description <string>]
```

### `tracker join team|user`
Given a `userId` and `teamId` (UUIDs, respectively), this command will create a record in `teamMemberships` to join a user to the team.
```bash
tracker join team --userId <userId> --teamId <teamId>
tracker join user --userId <userId> --teamId <teamId>
```
(Note: both forms of the command are functionally identical)

### `tracker create projects`
Given the following input parameters, create a new project object in the `projects` database table and return the new `projectId` (UUID):
```bash
tracker create project --name <project name> \
					   --iconId <iconId> \
					   --ownerId <ownerId> \
					   --teamId <teamId> \
					   --permissions {owner:<o>,team:<t>,everyone:<e>} \
					   [--description <string>]
```

### `tracker create workflow`
Given the following input parameters, create a workflow in the `workflows` database table, returning the new `workflowId` (UUID):
```bash
tracker create workflow --name <workflow name> \
						--iconId <iconId> \
						--ownerId <ownerId> \
						--teamId <teamId> \
					   --permissions {owner:<o>,team:<t>,everyone:<e>} \
					   [--description <string>]
```

### `tracker create workflowAction`
Given the following input parameters, create a record in the `workflowActions` database table and print the `actionId` (UUID) to `stdout`:
```bash
tracker create workflowAction --name <action name> \
							  --topic <string> \
							  --message <format string> \
							  [--description <string>]
```

### `tracker create workflowStep`
Given the following input parameters, create a record in the `workflowSteps` database table and print the `stepId` (UUID) to `stdout`:
```bash
tracker create workflowStep --name <step name> \
							 --workflow <workflowId> \
							 [--prevStepId <stepId>] \
							 [--nextStepId <stepId>]
							 [--description <string>]
```
Notes:
> `--prevStepId` will assume 'start' node if not present.
> `--nextStepId` will assume 'terminate' node if not present.

### `tracker create ticketType`
Given the following input parameters, create a record in the `ticketTypes` database table and print the generated `ticketTypeId` (UUID) to stdout:
```bash
tracker create ticketType --name <type name> \
						  --iconId <iconId> \
						  --workflowId <workflowId> \
						  [--description <string>]
```
This command will map a ticket type (e.g. "task" to a workflow and icon).

### `tracker join projectTicketType`
This command will create a record on `projectTicketTypes` to associate a ticket type with a project that supports said type:
```bash
tracker join projectTicketType --project <projectId> --ticketType <ticketTypeId>
```

### `tracker create ticket`
This command will create a new ticket:
```bash
tracker create ticket --project <id or "project name"> \
					  --type <id or "type name"> \
					  [--assignee <id or email>] \
					  [--permissions {owner:<o>,team:<t>,everyone:<e>}] \
					  --summary <string> \
					  --description <string>
```
Notes:
> if project name is used instead of a uuid, the name will be looked up to obtain the UUID.
> if type name is used instead of a uuid, the name will be looked up to obtain the UUID.
> if `--assignee` is not specified, the project `ownerId` is used.
> If `--permissions` are not specified, the project permissions are used.

### `tracker create attachment`
This command will upload a given attachment and attach the artifact to the specified ticket:
```bash
tracker create attachment --ticket <id> --file <path/url>
```

### `tracker create comment`
This command will add a comment to the associated ticket:
```bash
tracker create comment --ticket <id> [--comment <string>]
```
Notes:
> If no `--comment <string>` is provided, the comment will be read from `stdin`.