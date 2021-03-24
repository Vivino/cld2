// Package cld2 implements language detection using the
// Compact Language Detector.
//
// This package includes the relevant sources from the cld2
// project, so it doesn't require any external dependencies.
// For more information about CLD2, see https://github.com/CLD2Owners/cld2
package cld2

import (
	"errors"

	"github.com/Vivino/cld2/internal/info"
)

// Detect returns the language code for detected language
// in the given text.
var Detect = func(text string) string { return UNKNOWN_LANGUAGE.Code() }

// DetectLang returns the language code for detected language
// in the given text.
// ENGLISH is returned if the language cannot be detected.
var DetectLang = func(text string) Language { return UNKNOWN_LANGUAGE }

// DetectThree returns up to three language guesses.
// Extended languages are enabled.
// Unknown languages are removed from the resultset.
var DetectThree = func(text string) Languages { return Languages{} }

// ErrNoPlugins is returned if attempting to load plugin on non-supported platform.
var ErrNoPlugins = errors.New("CLD2: Architecture does not support plugins")

// LoadPlugin can be used to load plugin on Linux and Go 1.8.
// If support has been compiled into the binary calling this will have no effect.
// If binary was compiled without CGO and platform does not support plugins ErrNoPlugins
// will be returned.
var LoadPlugin = func(path string) error {
	if Enabled {
		return nil
	}
	return ErrNoPlugins
}

// Enabled will be set to true if compiled with cgo, or plugin loaded
// successfully
var Enabled bool

func infoToLanguages(in info.Languages) Languages {
	res := Languages{
		TextBytes: in.TextBytes,
		Reliable:  in.Reliable,
	}
	for _, est := range in.Estimates {
		if Language(est.Language) == UNKNOWN_LANGUAGE {
			continue
		}
		res.Estimates = append(res.Estimates, Estimate{Language: Language(est.Language), NormScore: est.NormScore, Percent: est.Percent})
	}
	return res
}
