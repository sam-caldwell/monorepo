package commands

type Commands uint8

const (
    Create Commands = iota
    Read
    Update
    Delete
    List
)
