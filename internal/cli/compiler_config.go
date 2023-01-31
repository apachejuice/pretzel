package cli

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"

	"golang.org/x/exp/slices"
)

// An object representing the compiler configuration JSON.
type CompilerConfig struct {
	// Exclude patterns: by default all files from the source tree are included in the compilation
	ExcludePatterns []string `json:"exclude"`
	// The source tree directory (default: "source")
	SourceDir string `json:"sourceDir"`
	// A limit of errors at which to stop.
	ErrorLimit int `json:"errorLimit"`

	// The full path to the configuration file.
	Path string
	// A list of filenames to use as the code.
	Files []string
	// Ignored file count
	Ignored int
}

func (c *CompilerConfig) setDefaults() {
	if c.SourceDir == "" {
		c.SourceDir = "source"
	}
}

func (c *CompilerConfig) FindFiles() error {
	sourceDir := filepath.Join(filepath.Dir(c.Path), c.SourceDir)
	patterns := make([]*regexp.Regexp, 0)
	for _, pat := range c.ExcludePatterns {
		compiled, err := regexp.Compile(pat)
		if err != nil {
			return err
		}

		patterns = append(patterns, compiled)
	}

	return filepath.WalkDir(sourceDir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		if err != nil {
			return nil
		}

		for _, pattern := range patterns {
			if pattern.Match([]byte(path)) {
				c.Ignored++
				return nil
			}
		}

		ext := filepath.Ext(path)
		if ext == ".pz" && !slices.Contains(c.Files, path) { // for some reason it does things twice?
			c.Files = append(c.Files, path)
		}

		return nil
	})
}

// Reads the compiler configuration from the filename.
func ReadCompilerConfig(path string) (*CompilerConfig, error) {
	abspath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(abspath)
	if err != nil {
		return nil, err
	}

	var conf CompilerConfig
	err = json.Unmarshal(content, &conf)
	if err != nil {
		return nil, err
	}

	conf.Path = abspath
	conf.Files = make([]string, 0)
	conf.setDefaults()
	return &conf, nil
}
