package Atlassian

// Descriptor - a Generic descriptor interface for working with the Atlassian API
type Descriptor interface {
    Load(fileName string) error
    Create() (any, error)
    Read() (any, error)
    Update() (any, error)
    Delete() (any, error)
    List() (any, error)
}
