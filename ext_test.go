package fpath

import "testing"

func TestSplitExt(t *testing.T) {
	tests := []struct {
		name string
		file string
		ext  string
	}{
		{"/foo/bar/baz.go", "/foo/bar/baz", "go"},
		{"a.c", "a", "c"},
		{"a.c.o", "a.c", "o"},
		{"foo.long-extension", "foo", "long-extension"},
		{"/foo/bar/baz", "/foo/bar/baz", ""},
		{"foo", "foo", ""},
		{"/foo/bar/", "/foo/bar/", ""},
		{"/foo/.bar/", "/foo/.bar/", ""},
		{"/", "/", ""},
		{"", "", ""},
		{".", ".", ""},
		{"..", "..", ""},
		{"", "", ""},
		{"/foo/bar/.a.c", "/foo/bar/.a", "c"},
		{".a", ".a", ""},
		{".a.c", ".a", "c"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			file, ext := SplitExt(test.name)
			if file != test.file {
				t.Errorf("file error: expected '%s', got '%s'", test.file, file)
			} else if ext != test.ext {
				t.Errorf("ext error: expected '%s', got '%s'", test.ext, ext)
			}
		})
	}
}

func TestReplaceExt(t *testing.T) {
	tests := []struct {
		name   string
		ext    string
		result string
	}{
		{"/foo/bar/baz.go", "cpp", "/foo/bar/baz.cpp"},
		{"foo.go", "c", "foo.c"},
		{"a.tmp.c", "o", "a.tmp.o"},
		{"foo.long-extension", "o", "foo.o"},
		{"/foo/bar/baz", "go", "/foo/bar/baz.go"},
		{"foo", "go", "foo.go"},
		{"/foo/bar/baz.go", "", "/foo/bar/baz"},
		{"foo.go", "", "foo"},
		{"foo", "go", "foo.go"},
		{"/foo/bar/", "go", ""},
		{".", "go", ""},
		{"", "go", ""},
		{"", "", ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ReplaceExt(test.name, test.ext)
			if result != test.result {
				t.Errorf("expected '%s', got '%s'", test.result, result)
			}
		})
	}
}

func TestRemoveExt(t *testing.T) {
	tests := []struct {
		name   string
		result string
	}{
		{"/foo/bar/baz.go", "/foo/bar/baz"},
		{"foo.go", "foo"},
		{"foo", "foo"},
		{"foo.cpp.o", "foo.cpp"},
		{".foo.go", ".foo"},
		{".foo", ".foo"},
		{"foo/", "foo/"},
		{"/", "/"},
		{".", "."},
		{"..", ".."},
		{"", ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := RemoveExt(test.name)
			if result != test.result {
				t.Errorf("expected '%s', got '%s'", test.result, result)
			}
		})
	}
}
