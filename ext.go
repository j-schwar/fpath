package fpath

// SplitExt splits a file path into file+directories and file extension
// portions.
func SplitExt(path string) (file string, ext string) {
	n := len(path)
	file = path
	ext = ""

	for i := n - 1; i >= 0; i -= 1 {
		if path[i] == '.' {
			if i == n-1 || i == 0 {
				file = path
				ext = ""
			} else {
				file = path[0:i]
				ext = path[i+1:]
			}
			break
		} else if path[i] == '/' {
			break
		}
	}

	return
}

// ReplaceExt replaces the file extension section of a file path with a new one.
//
// When specifying a new extension, it is not necessary to include the '.'
// prefix as it will be added automatically.
func ReplaceExt(path string, ext string) string {
	file, _ := SplitExt(path)
	last := file[len(file)-1]
	if file == "" || last == '/' || last == '.' {
		return ""
	}
	if ext == "" {
		return file
	}
	return file + "." + ext
}

// RemoveExt returns a path without the file extension.
func RemoveExt(path string) string {
	r := ReplaceExt(path, "")
	if r == "" {
		return path
	}
	return r
}
