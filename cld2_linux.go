//+build linux

package cld2

import (
	"log"
	"plugin"

	"github.com/klauspost/cld2/internal/info"
)

//go:generate go build -buildmode=plugin -o lib/cld2go.so internal/plugin/cld2_linux.go

func init() {
	p, err := plugin.Open("lib/cld2go.so")
	if err != nil {
		log.Printf("CLD2: Load cld2go.so: %v", err)
		return
	}

	f, err := p.Lookup("Detect")
	if err != nil {
		log.Printf("CLD2: Load Detect: %v", err)
		return
	}
	Detect = f.(func(text string) string)

	f, err := p.Lookup("DetectLang")
	if err != nil {
		log.Printf("CLD2: Load DetectLang: %v", err)
		return
	}
	DetectLang = func(text string) Language {
		l := f.(func(text string) uint16)
		return Language(l(text))
	}

	f, err := p.Lookup("DetectThree")
	if err != nil {
		log.Printf("CLD2: Load DetectThree: %v", err)
		return
	}
	DetectThree = func(text string) Languages {
		l := f.(func(text string) info.Languages)
		return infoToLanguages(l(text))
	}
	Enabled = true
}
