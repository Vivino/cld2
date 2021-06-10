//+build !cld2_disable,cgo,!linux,!darwin darwin,linux,!go1.8,cgo darwin,linux,go1.8,!amd64

// Package cld2 implements language detection using the
// Compact Language Detector.
//
// This package includes the relevant sources from the cld2
// project, so it doesn't require any external dependencies.
// For more information about CLD2, see https://github.com/CLD2Owners/cld2
package cld2

import (
	"github.com/Vivino/cld2/internal/cld2plugin"
	"github.com/Vivino/cld2/internal/info"
)

func init() {
	Detect = func(text string) string {
		return cld2plugin.Detect(text)
	}

	DetectLang = func(text string) Language {
		return Language(cld2plugin.DetectLang(text))
	}

	DetectThree = func(text string) Languages {
		return infoToLanguages(cld2plugin.DetectThree(text).(info.Languages))
	}
	Enabled = true
}
