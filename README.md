# fpath

Simple file path manipulation focusing on splits and replacements.

## Installation

```
$ go get -u github.com/j-schwar/fpath
```

## Examples

### File Extensions

Splitting a path on its extension:
```go
file, ext := fpath.SplitExt("foo.go")
// file -> "foo"
// ext -> "go"
```

Replacing the extension of a file:
```go
objFile := fpath.ReplaceExt("foo.cpp", "o")
// objFile -> "foo.o"
```

Removing a file extension:
```go
file := fpath.RemoveExt("foo.go")
// file -> "foo"
```
