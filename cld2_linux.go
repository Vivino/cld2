//+build linux,go1.8

package cld2

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"plugin"

	"github.com/klauspost/cld2/internal/info"
)

//go:generate go build -buildmode=plugin -o lib/cld2go.so github.com/klauspost/cld2/internal/plugin

func init() {
	p, err := findPlugin("lib", ".", "/lib", filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "klauspost", "cld2", "lib"))
	if err != nil {
		log.Printf("CLD2: Error loading plugin: %v", err)
		return
	}

	f, err := p.Lookup("PluginDetect")
	if err != nil {
		log.Printf("CLD2: Load Detect: %v", err)
		return
	}
	Detect = f.(func(text string) string)

	f, err = p.Lookup("PluginDetectLang")
	if err != nil {
		log.Printf("CLD2: Load DetectLang: %v", err)
		return
	}
	dl := f.(func(text string) uint16)
	DetectLang = func(text string) Language {
		return Language(dl(text))
	}

	f, err = p.Lookup("PluginDetectThree")
	if err != nil {
		log.Printf("CLD2: Load DetectThree: %v", err)
		return
	}

	dt := f.(func(text string) info.Languages)
	DetectThree = func(text string) Languages {
		return infoToLanguages(dt(text))
	}
	Enabled = true
}

func findPlugin(paths ...string) (*plugin.Plugin, error) {
	for i, path := range paths {
		var err error
		paths[i], err = filepath.Abs(filepath.Join(path, "cld2go.so"))
		if err != nil {
			continue
		}
		p, err := plugin.Open(paths[i])
		if err == nil {
			return p, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("cld2go.so could not be found in any of: %v", paths))
}
