package packageManager

// pkgManagerFunc - A function signature representing our package manager wrapper functions.
type pkgManagerFunc func(DependencyDescriptor) (string, error)
