package monorepo

// DependencyList - simplified view of dependencies for FlattenProjectDependencies() and SortManifestsByDependencies()
type DependencyList struct {
	Name         string
	Dependencies []string
}
