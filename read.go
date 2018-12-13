package gtp

import (
	//"fmt"
	"io/ioutil"
	"path/filepath"
	//"strings"
)

// readFiles reads files and returns source for the parsing process.
func readFiles(basePath, innerPath string, opts *Options) (*source, error) {
	// Read the base file.
	base, err := readFile(basePath, opts)
	if err != nil {
		return nil, err
	}

	// Read the inner file.
	inner, err := readFile(innerPath, opts)
	if err != nil {
		return nil, err
	}

	var includes []*File

	/*
	// Find include files from the base file.
	if err := findIncludes(base.data, opts, &includes, base); err != nil {
		return nil, err
	}

	// Find include files from the inner file.
	if err := findIncludes(inner.data, opts, &includes, inner); err != nil {
		return nil, err
	}
	*/

	return NewSource(base, inner, includes), nil
}

// readFile reads a file and returns a file struct.
func readFile(path string, opts *Options) (*File, error) {
	var data []byte
	var err error

	if path != "" {
		name := filepath.Join(opts.BaseDir, path + "." + opts.Extension)

		if opts.Asset != nil {
			data, err = opts.Asset(name)
		} else {
			data, err = ioutil.ReadFile(name)
		}

		if err != nil {
			return nil, err
		}
	}

	return NewFile(path, data), nil
}
