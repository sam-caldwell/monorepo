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
