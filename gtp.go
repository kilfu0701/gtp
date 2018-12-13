package gtp

import (
	"html/template"
)

func Load(basePath, innerPath string, opts *Options) (*template.Template, error) {
	opts = InitializeOptions(opts)
	_ = basePath + ":" + innerPath

	/*
	if !opts.DynamicReload {
		if tpl, ok := getCache(name); ok {
			return &tpl, nil
		}
	}
	*/

	// Read files.
	_, err := readFiles(basePath, innerPath, opts)
	if err != nil {
		return nil, err
	}

	/*
	// Parse the source.
	rslt, err := ParseSource(src, opts)
	if err != nil {
		return nil, err
	}

	// Compile the parsed result.
	tpl, err := CompileResult(name, rslt, opts)
	if err != nil {
		return nil, err
	}

	if !opts.DynamicReload {
		setCache(name, *tpl)
	}
	*/

	return nil, nil
}
