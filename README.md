# sembump
Takes a semver and bump type `major, minor, patch` and increments it.

## Install
```
$ go get github.com/threecommaio/sembump
```

## Usage
```
$ sembump
usage: sembump <semver> <major|minor|patch>

$ sembump v1.2.0 major
v2.2.1

$ sembump v1.2.0 minor
v1.3.1

$ sembump v1.2.0 patch
v1.2.1
```