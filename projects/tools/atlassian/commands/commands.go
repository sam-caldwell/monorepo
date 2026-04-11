package commands

type Commands uint8

const (
    Create Commands = iota
    Read
    Update
    Delete
    List
)

func (cmd *Commands) String() string {
    switch *cmd {
    case Create:
        return "create"
    case Read:
        return "read"
    case Update:
        return "update"
    case Delete:
        return "delete"
    case List:
        return "list"
    default:
        return "invalid"
    }
}
