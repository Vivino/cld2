//+build linux,go1.8,amd64

package cld2

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"plugin"

	"github.com/Vivino/cld2/internal/info"
)

//go:generate go build -buildmode=plugin -o lib/cld2go.so github.com/Vivino/cld2/internal/plugin

func init() {
	LoadPlugin = func(path string) error {
		if Enabled {
			return nil
		}
		p, err := plugin.Open(path)
		if err != nil {
			return err
		}
		return usePlugin(p)

	}
	p, err := findPlugin("./lib", "/lib", "lib", ".", filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "Vivino", "cld2", "lib"))
	if err != nil {
		log.Printf("CLD2: Loading plugin: %v", err)
		return
	}
	err = usePlugin(p)
	if err != nil {
		log.Printf("CLD2: plugin init: %v", err)
		return
	}
}

func usePlugin(p *plugin.Plugin) error {
	f, err := p.Lookup("PluginDetect")
	if err != nil {
		return err
	}
	var ok bool
	Detect, ok = f.(func(text string) string)
	if !ok {
		return fmt.Errorf("PluginDetect: wrong signature: %T", f)
	}

	f, err = p.Lookup("PluginDetectLang")
	if err != nil {
		return err
	}
	dl, ok := f.(func(text string) uint16)
	if !ok {
		return fmt.Errorf("PluginDetectLang: wrong signature: %T", f)
	}
	DetectLang = func(text string) Language {
		return Language(dl(text))
	}

	f, err = p.Lookup("PluginDetectThree")
	if err != nil {
		return err
	}
	dt, ok := f.(func(text string) info.Languages)
	if !ok {
		return fmt.Errorf("PluginDetectThree: wrong signature: %T", f)
	}
	DetectThree = func(text string) Languages {
		return infoToLanguages(dt(text))
	}
	Enabled = true
	return nil
}

func findPlugin(paths ...string) (*plugin.Plugin, error) {
	for i, path := range paths {
		var err error
		paths[i], err = filepath.Abs(filepath.Join(path, "cld2go.so"))
		if err != nil {
			fmt.Fprintln(os.Stdout, "could not create absolute path for (%s): %v", path, err)
			continue
		}
		p, err := plugin.Open(paths[i])
		if err == nil {
			return p, nil
		}
		fmt.Fprintln(os.Stdout, "did not find cld2go.so in path:", paths[i])
	}
	return nil, fmt.Errorf("cld2go.so could not be found in any of: %v", paths)
}
