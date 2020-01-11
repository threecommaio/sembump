// sembump: take a semver 2.0 compliant string and increment
// credit goes to: https://github.com/peterj/semver for the bump function
//
// ThreeComma
// hello@threecomma.io
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/blang/semver"
)

// VersionBump type represents which part of
// the version is being bumped
type VersionBump int

const (
	// Major version bump
	Major VersionBump = iota
	// Minor version bump
	Minor
	// Patch version bump
	Patch
)

// Bump bumps the provided version based on the version bump type
func Bump(version string, bump VersionBump) (string, error) {
	hasPrefix := strings.HasPrefix(version, "v")
	if hasPrefix {
		version = strings.TrimPrefix(version, "v")
	}

	parsedVersion, err := semver.Make(version)
	if err != nil {
		return "", fmt.Errorf("invalid version provided")
	}

	switch bump {
	case Major:
		{
			parsedVersion.Major++
			break
		}
	case Minor:
		{
			parsedVersion.Minor++
			break
		}
	case Patch:
		{
			parsedVersion.Patch++
			break
		}
	}

	version = parsedVersion.String()
	if hasPrefix {
		version = "v" + version
	}
	return version, nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: sembump <semver> <major|minor|patch>")
		os.Exit(1)
	}

	var err error
	var nextVersion string

	version := os.Args[1]
	bumpType := os.Args[2]

	switch bumpType {
	case "major":
		nextVersion, err = Bump(version, Major)
	case "minor":
		nextVersion, err = Bump(version, Minor)
	case "patch":
		nextVersion, err = Bump(version, Patch)
	default:
		fmt.Printf("unknown bump type: (%s). [major, minor, or patch is only supported]", bumpType)
	}

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	fmt.Println(nextVersion)
}
