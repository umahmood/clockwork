package clockwork

import "fmt"

// Semantic versioning - http://semver.org/
const (
	Major = 1
	Minor = 2
	Patch = 1
)

// Version returns library version.
func Version() string {
	return fmt.Sprintf("%d.%d.%d", Major, Minor, Patch)
}
